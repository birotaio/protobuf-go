package protofif

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var (
	ErrorInvalidTimestamp = errors.New("invalid timestamp")
	ErrorInvalidDuration  = errors.New("invalid duration")
)

func Now() *Timestamp {
	t := time.Now()
	return &Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func NewTimestampValue(t time.Time) Timestamp {
	return Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
		Loc:     t.Location().String(),
	}
}

func NewTimestamp(t *time.Time) *Timestamp {
	if t == nil {
		return nil
	}

	return &Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
		Loc:     t.Location().String(),
	}
}

func NewTimestampFromValue(t time.Time) *Timestamp {
	return &Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
		Loc:     t.Location().String(),
	}
}

func (t *Timestamp) AsTime() *time.Time {
	if t == nil {
		return nil
	}
	loc, err := time.LoadLocation(t.Loc)
	if err != nil {
		loc = time.Local
	}
	ts := time.Unix(t.Seconds, int64(t.Nanos)).In(loc)
	return &ts
}

func (t *Timestamp) Add(d *Duration) *Timestamp {
	timestamp := NewTimestampValue(t.AsTime().Add(d.AsDurationValue()))
	return &timestamp
}

func (t *Timestamp) AsTimeValue() time.Time {
	if t == nil {
		return time.Time{}
	}
	loc, err := time.LoadLocation(t.Loc)
	if err != nil {
		loc = time.Local
	}
	return time.Unix(t.Seconds, int64(t.Nanos)).In(loc)
}

func (ts *Timestamp) UnmarshalJSON(b []byte) error {
	timeUnmarshaled := time.Time{}
	err := json.Unmarshal(b, &timeUnmarshaled)
	if err != nil {
		return err
	}
	*ts = NewTimestampValue(timeUnmarshaled)
	return nil
}

func (ts *Timestamp) MarshalJSON() ([]byte, error) {
	if ts == nil {
		return nil, nil
	}
	return json.Marshal(ts.AsTime())
}

func (ts *Timestamp) AppendValue(b []byte, flags int) ([]byte, error) {
	if ts == nil {
		return types.AppendNull(b, flags), nil
	}
	return types.AppendTime(b, ts.AsTimeValue(), flags), nil
}

func (ts *Timestamp) ScanValue(rd types.Reader, n int) error {
	tm, err := types.ScanTime(rd, n)
	if err != nil {
		return err
	}
	*ts = NewTimestampValue(tm)
	return nil
}

func (ts *Timestamp) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if ts == nil {
		return bson.TypeNull, nil, nil
	}
	return bson.MarshalValue(ts.AsTime())
}

func (ts *Timestamp) UnmarshalBSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	t := time.Time{}
	err := bson.UnmarshalValue(bson.TypeDateTime, data, &t)
	if err != nil {
		return err
	}
	*ts = NewTimestampValue(t)
	return nil
}

func (ts *Timestamp) GormDataType() string {
	return "datetime"
}

var _ sql.Scanner = (*Timestamp)(nil)

func (ts *Timestamp) Scan(src interface{}) error {
	switch casted := src.(type) {
	case time.Time:
		*ts = NewTimestampValue(casted)
	case *time.Time:
		if casted != nil {
			*ts = NewTimestampValue(*casted)
		}
	case string:
		t, err := time.Parse(time.RFC3339, casted)
		if err != nil {
			return err
		}
		*ts = NewTimestampValue(t)
	default:
		return fmt.Errorf("column of type %T cannot be scanned into a Timestamp", src)
	}
	return nil
}

var _ driver.Valuer = (*Timestamp)(nil)

func (ts *Timestamp) Value() (driver.Value, error) {
	if ts == nil {
		return nil, nil
	}
	return ts.AsTimeValue(), nil
}

// Same thing for time.Duration
func NewDurationValue(d time.Duration) Duration {
	return Duration{
		Nanoseconds: d.Nanoseconds(),
	}
}

func NewDuration(d time.Duration) *Duration {
	return &Duration{
		Nanoseconds: d.Nanoseconds(),
	}
}

func (d *Duration) AsDuration() *time.Duration {
	if d == nil {
		return nil
	}
	td := time.Duration(d.Nanoseconds)
	return &td
}

func (d *Duration) AsDurationValue() time.Duration {
	if d == nil {
		return time.Duration(0)
	}
	return time.Duration(d.Nanoseconds)
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var int64Val int64
	err := json.Unmarshal(b, &int64Val)
	if err != nil {
		return err
	}
	*d = Duration{Nanoseconds: int64Val}
	return nil
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte(`""`), nil
	}
	return json.Marshal(d.Nanoseconds)
}

func (d *Duration) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if d == nil {
		return bson.TypeNull, nil, nil
	}
	return bson.MarshalValue(d.Nanoseconds)
}

func (d *Duration) UnmarshalBSONValue(btype bsontype.Type, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var int64Val int64
	err := bson.UnmarshalValue(btype, data, &int64Val)
	if err != nil {
		return err
	}

	*d = Duration{Nanoseconds: int64Val}
	return nil
}
