package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const (
	timeFormat     = "2006-01-02 15:04:05"
	jsonTimeFormat = `"2006-01-02 15:04:05"`
)

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(data []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(data, &v); err == nil {
		switch value := v.(type) {
		case float64:
			d.Duration = time.Duration(value)
		case string:
			d.Duration, err = time.ParseDuration(value)
		default:
			err = fmt.Errorf("Invalid duration [%s] ", value)
		}
	}
	return
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func Now() *Time {
	return &Time{Time: time.Now()}
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.ParseInLocation(jsonTimeFormat, string(data), time.Local)
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 21)
	b = t.AppendFormat(b, jsonTimeFormat)
	return b, nil
}

func (t *Time) String() string {
	return t.Format(timeFormat)
}

func (t *Time) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	return nil
}

func (t *Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t==nil||t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
