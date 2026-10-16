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

const DateLayout = time.DateOnly

var ErrorInvalidDate = errors.New("invalid date")

func NewDateValue(t time.Time) Date {
	y, m, d := t.Date()
	return Date{Year: int32(y), Month: int32(m), Day: int32(d)}
}

func NewDate(t *time.Time) *Date {
	if t == nil {
		return nil
	}

	date := NewDateValue(*t)
	return &date
}

func NewDateFromValue(t time.Time) *Date {
	date := NewDateValue(t)
	return &date
}

// ParseDate accepts YYYY-MM-DD or an RFC 3339 midnight, keeping the day as written in its offset.
func ParseDate(s string) (*Date, error) {
	t, err := parseDate(s)
	if err != nil {
		return nil, err
	}
	return NewDateFromValue(t), nil
}

func parseDate(s string) (time.Time, error) {
	if t, err := time.Parse(DateLayout, s); err == nil {
		return t, nil
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %q is neither %s nor RFC 3339", ErrorInvalidDate, s, DateLayout)
	}
	if t.Hour() != 0 || t.Minute() != 0 || t.Second() != 0 || t.Nanosecond() != 0 {
		return time.Time{}, fmt.Errorf("%w: %q is not a midnight", ErrorInvalidDate, s)
	}
	return t, nil
}

func (d *Date) AsTime() *time.Time {
	if d == nil {
		return nil
	}
	t := d.AsTimeValue()
	return &t
}

func (d *Date) AsTimeValue() time.Time {
	if d == nil || (d.Year == 0 && d.Month == 0 && d.Day == 0) {
		return time.Time{}
	}
	return time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
}

func (d *Date) Format() string {
	if d == nil {
		return ""
	}
	return d.AsTimeValue().Format(DateLayout)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("%w: %s", ErrorInvalidDate, b)
	}
	return d.scanString(s)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte("null"), nil
	}
	return json.Marshal(d.Format())
}

func (d *Date) AppendValue(b []byte, flags int) ([]byte, error) {
	if d == nil {
		return types.AppendNull(b, flags), nil
	}
	return types.AppendString(b, d.Format(), flags), nil
}

func (d *Date) ScanValue(rd types.Reader, n int) error {
	t, err := types.ScanTime(rd, n)
	if err != nil {
		return err
	}
	*d = NewDateValue(t)
	return nil
}

func (d *Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if d == nil {
		return bson.TypeNull, nil, nil
	}
	return bson.MarshalValue(d.AsTimeValue())
}

func (d *Date) UnmarshalBSONValue(btype bsontype.Type, data []byte) error {
	switch btype {
	case bson.TypeNull, bson.TypeUndefined:
		return nil
	case bson.TypeString:
		var s string
		if err := bson.UnmarshalValue(btype, data, &s); err != nil {
			return err
		}
		return d.scanString(s)
	default:
		t := time.Time{}
		if err := bson.UnmarshalValue(btype, data, &t); err != nil {
			return err
		}
		*d = NewDateValue(t.UTC())
		return nil
	}
}

func (d *Date) GormDataType() string {
	return "date"
}

var _ sql.Scanner = (*Date)(nil)

func (d *Date) Scan(src interface{}) error {
	switch casted := src.(type) {
	case time.Time:
		*d = NewDateValue(casted)
	case *time.Time:
		if casted != nil {
			*d = NewDateValue(*casted)
		}
	case string:
		return d.scanString(casted)
	case []byte:
		return d.scanString(string(casted))
	default:
		return fmt.Errorf("column of type %T cannot be scanned into a Date", src)
	}
	return nil
}

func (d *Date) scanString(s string) error {
	t, err := parseDate(s)
	if err != nil {
		return err
	}
	*d = NewDateValue(t)
	return nil
}

var _ driver.Valuer = (*Date)(nil)

func (d *Date) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}
	return d.AsTimeValue(), nil
}
