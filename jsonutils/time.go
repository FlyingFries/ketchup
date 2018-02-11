package jsonutils

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func (ts Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(ts.UTC().UnixNano()/int64(time.Millisecond), 10) + `"`), nil
}

func (ts Time) UnmarshalJSON(b []byte) error {
	return fmt.Errorf("Time UnmarshalJSON not supported")
}
func (ts Time) Value() (driver.Value, error) {
	return ts.Time, nil
}
func (ts *Time) Scan(src interface{}) error {
	ts.Time = src.(time.Time)
	return nil
}

func NewTime(t time.Time) *Time {
	return &Time{t}
}
