package protofif

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimestampUnmarshalJSON(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  int64
	}{
		{"rfc3339", `"2026-09-22T10:20:30Z"`, 1790072430},
		{"epoch seconds", `1789976036`, 1789976036},
		{"epoch zero", `0`, 0},
		{"epoch negative", `-1`, -1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var ts Timestamp
			if err := json.Unmarshal([]byte(tc.input), &ts); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.input, err)
			}
			if ts.Seconds != tc.want {
				t.Errorf("got seconds %d, want %d", ts.Seconds, tc.want)
			}
		})
	}
}

func TestTimestampUnmarshalJSONNull(t *testing.T) {
	ts := Timestamp{Seconds: 42}
	if err := json.Unmarshal([]byte(`null`), &ts); err != nil {
		t.Fatalf("unmarshal null: %v", err)
	}
	if ts.Seconds != (time.Time{}).Unix() {
		t.Errorf("got seconds %d, want zero time", ts.Seconds)
	}
}

// Generated messages carry timestamps as pointer fields; devices report them as
// epoch seconds, so decoding must not fail on a bare number.
func TestTimestampUnmarshalJSONPointerField(t *testing.T) {
	var battery struct {
		ValidTimestamp  *Timestamp `json:"valid_timestamp"`
		StatusTimestamp *Timestamp `json:"status_timestamp"`
	}
	body := `{"valid_timestamp":0,"status_timestamp":1789976036}`
	if err := json.Unmarshal([]byte(body), &battery); err != nil {
		t.Fatalf("unmarshal %s: %v", body, err)
	}
	if battery.ValidTimestamp.Seconds != 0 {
		t.Errorf("got valid_timestamp %d, want 0", battery.ValidTimestamp.Seconds)
	}
	if battery.StatusTimestamp.Seconds != 1789976036 {
		t.Errorf("got status_timestamp %d, want 1789976036", battery.StatusTimestamp.Seconds)
	}
}

func TestTimestampUnmarshalJSONInvalid(t *testing.T) {
	var ts Timestamp
	err := json.Unmarshal([]byte(`"not-a-date"`), &ts)
	if err == nil {
		t.Fatal("expected an error for a malformed timestamp string")
	}
}
