package structpb_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"

	spb "google.golang.org/protobuf/types/known/structpb"
)

func TestUnmarshalBSONValue(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want interface{}
	}{
		{
			name: "list",
			in:   bson.A{2.27, 48.82},
			want: []interface{}{2.27, 48.82},
		},
		{
			name: "nested document",
			in:   bson.M{"a": bson.M{"b": int32(1)}},
			want: map[string]interface{}{"a": map[string]interface{}{"b": 1.0}},
		},
		{
			name: "list in document",
			in:   bson.M{"coordinates": bson.A{2.27, 48.82}},
			want: map[string]interface{}{"coordinates": []interface{}{2.27, 48.82}},
		},
		{
			name: "list in nested document",
			in:   bson.M{"a": bson.M{"b": bson.A{int32(1), bson.M{"c": bson.A{"x"}}}}},
			want: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{1.0, map[string]interface{}{"c": []interface{}{"x"}}}}},
		},
		{
			name: "datetime",
			in:   primitive.NewDateTimeFromTime(time.Date(2026, 10, 14, 0, 0, 53, 952000000, time.UTC)),
			want: "2026-10-14T00:00:53.952Z",
		},
		{
			name: "datetime in document",
			in:   bson.M{"last_updated_at": primitive.NewDateTimeFromTime(time.Date(2026, 10, 14, 2, 0, 0, 0, time.FixedZone("CEST", 2*3600)))},
			want: map[string]interface{}{"last_updated_at": "2026-10-14T00:00:00Z"},
		},
		{
			name: "datetime in list in document",
			in:   bson.M{"a": bson.A{primitive.NewDateTimeFromTime(time.Date(2026, 10, 14, 0, 0, 0, 0, time.UTC))}},
			want: map[string]interface{}{"a": []interface{}{"2026-10-14T00:00:00Z"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, raw, err := bson.MarshalValue(tt.in)
			if err != nil {
				t.Fatalf("bson.MarshalValue: %v", err)
			}
			got := new(spb.Value)
			if err := got.UnmarshalBSONValue(typ, raw); err != nil {
				t.Fatalf("UnmarshalBSONValue: %v", err)
			}
			want, err := spb.NewValue(tt.want)
			if err != nil {
				t.Fatalf("NewValue: %v", err)
			}
			if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
