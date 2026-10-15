package protofif

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestDateUnmarshalJSON(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"date only", `"2026-05-11"`, "2026-05-11"},
		{"utc midnight", `"2026-05-11T00:00:00Z"`, "2026-05-11"},
		{"offset midnight keeps the written day", `"2026-05-11T00:00:00+02:00"`, "2026-05-11"},
		{"negative offset midnight keeps the written day", `"2026-05-11T00:00:00-10:00"`, "2026-05-11"},
		{"leap day", `"2024-02-29"`, "2024-02-29"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var d Date
			if err := json.Unmarshal([]byte(tc.input), &d); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.input, err)
			}
			if got := d.Format(); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestDateUnmarshalJSONInvalid(t *testing.T) {
	for _, input := range []string{
		`"2026-05-10T22:00:00Z"`,
		`"2026-05-11T00:00:00.5Z"`,
		`"2026-02-30"`,
		`"11/05/2026"`,
		`"not-a-date"`,
		`1789976036`,
	} {
		t.Run(input, func(t *testing.T) {
			var d Date
			err := json.Unmarshal([]byte(input), &d)
			if !errors.Is(err, ErrorInvalidDate) {
				t.Errorf("got error %v, want ErrorInvalidDate", err)
			}
		})
	}
}

func TestDateUnmarshalJSONPointerField(t *testing.T) {
	var user struct {
		Birthdate *Date `json:"birthdate"`
		Other     *Date `json:"other"`
	}
	if err := json.Unmarshal([]byte(`{"birthdate":"1990-07-14","other":null}`), &user); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if user.Birthdate.Format() != "1990-07-14" {
		t.Errorf("got birthdate %s, want 1990-07-14", user.Birthdate.Format())
	}
	if user.Other != nil {
		t.Errorf("got other %v, want nil", user.Other)
	}
}

func TestDateMarshalJSON(t *testing.T) {
	b, err := json.Marshal(struct {
		Birthdate *Date `json:"birthdate"`
	}{&Date{Year: 1990, Month: 7, Day: 14}})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"birthdate":"1990-07-14"}`; string(b) != want {
		t.Errorf("got %s, want %s", b, want)
	}
}

func TestDateBSONRoundTrip(t *testing.T) {
	type user struct {
		Birthdate *Date `bson:"birthdate"`
	}
	b, err := bson.Marshal(user{&Date{Year: 1990, Month: 7, Day: 14}})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw struct {
		Birthdate time.Time `bson:"birthdate"`
	}
	if err := bson.Unmarshal(b, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}
	if want := time.Date(1990, 7, 14, 0, 0, 0, 0, time.UTC); !raw.Birthdate.Equal(want) {
		t.Errorf("stored %s, want %s", raw.Birthdate, want)
	}

	var got user
	if err := bson.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Birthdate.Format() != "1990-07-14" {
		t.Errorf("got %s, want 1990-07-14", got.Birthdate.Format())
	}
}

func TestDateUnmarshalBSON(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input interface{}
		want  string
	}{
		{"utc datetime", time.Date(1990, 7, 14, 0, 0, 0, 0, time.UTC), "1990-07-14"},
		{"string", "1990-07-14", "1990-07-14"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b, err := bson.Marshal(bson.M{"birthdate": tc.input})
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var got struct {
				Birthdate *Date `bson:"birthdate"`
			}
			if err := bson.Unmarshal(b, &got); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got.Birthdate.Format() != tc.want {
				t.Errorf("got %s, want %s", got.Birthdate.Format(), tc.want)
			}
		})
	}
}

func TestDateScan(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input interface{}
	}{
		{"time", time.Date(1990, 7, 14, 0, 0, 0, 0, time.UTC)},
		{"string", "1990-07-14"},
		{"bytes", []byte("1990-07-14")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var d Date
			if err := d.Scan(tc.input); err != nil {
				t.Fatalf("scan: %v", err)
			}
			if d.Format() != "1990-07-14" {
				t.Errorf("got %s, want 1990-07-14", d.Format())
			}
		})
	}
}

func TestDateZero(t *testing.T) {
	var d *Date
	if !d.AsTimeValue().IsZero() || d.AsTime() != nil || d.Format() != "" {
		t.Error("nil date should convert to zero values")
	}
	if !(&Date{}).AsTimeValue().IsZero() {
		t.Error("empty date should convert to the zero time")
	}
}
