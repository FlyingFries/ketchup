// Package timerange provides simple way to store time ranges - not working yet
// taken from https://gist.github.com/smagch/d2a55c60bbd76930c79f
package timerange

import (
	"fmt"
	"time"
)

var timeLayout = "15:04"

type jsonTime struct {
	time.Time
}

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(timeLayout) + `"`), nil
}

func (t *jsonTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	ret, err := time.Parse(timeLayout, s[1:6])
	if err != nil {
		return err
	}
	t.Time = ret
	return nil
}

/*Day represents open hours in one day
Opening hours must comply with the following form:
	["open", "close", ...]
*/
type Day []jsonTime

type Week [7][]jsonTime

type Opens []jsonTime

//Range contains two time, open and close
type Range [2]jsonTime

//IsOpen checks if time t is in range
func (r *Range) IsOpen(t time.Time) bool {
	h, m, _ := t.Clock()

	h1, m1, _ := r[0].Clock()
	h2, m2, _ := r[1].Clock()

	if h1 == h {
		return m >= m1
	}
	if h2 == h {
		return m <= m2
	}

	return h > h1 && h <= h2
}

func (r Range) String() string {
	return fmt.Sprintf("[%s, %s]", r[0].Format(timeLayout), r[1].Format(timeLayout))
}
