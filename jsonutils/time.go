package jsonutils

import (
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

func New(t time.Time) *Time {
	return &Time{t}
}
