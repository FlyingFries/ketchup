package timerange_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/FlyingFries/ketchup/timerange"
)

func tMust(h string) time.Time {
	t, err := time.Parse("15:04", h)
	if err != nil {
		panic(err)
	}
	return t
}

var jsonTestCases = []struct {
	raw  []byte
	in   time.Time
	want bool
}{
	{[]byte(`["08:00", "12:00", "13:00", "21:00"]`), tMust("00:59"), false},
	{[]byte(`["01:00", "12:00"]`), tMust("01:00"), true},
	{[]byte(`["01:00", "12:00"]`), tMust("12:00"), true},
	{[]byte(`["01:00", "12:00"]`), tMust("12:01"), false},
	{[]byte(`["00:00", "23:59"]`), tMust("00:01"), true},
	{[]byte(`["00:00", "23:59"]`), tMust("12:59"), true},
}

func TestTimeRange(t *testing.T) {

	for _, tc := range jsonTestCases {
		var r timerange.Range
		err := json.Unmarshal(tc.raw, &r)
		if err != nil {
			t.Error(err)
		}
		is := r.IsOpen(tc.in)
		if is != tc.want {
			t.Errorf(" (%s).IsOpen(%s) = %t; want %t", r, tc.in, is, tc.want)
		}
	}
}
