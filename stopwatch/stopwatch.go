/*
Package stopwatch provides simple way to measure time between certain calls

	sw := stopwatch.New()

	funcToMeasure()
	sw.Lap("my func")

	nextFuncToMeasure()
	sw.Lap("next func")

	go sendToInflux(sw.Laps())

*/
package stopwatch

import (
	"encoding/json"
	"time"
)

// Stopwatch performance metering
type Stopwatch struct {
	t1   time.Time
	laps map[string]time.Duration
}

//New creates stopwatch instance
func New() *Stopwatch {
	return &Stopwatch{time.Now(), make(map[string]time.Duration)}
}

//Lap records new mesure since last
func (sw *Stopwatch) Lap(name string) {
	if _, ok := sw.laps[name]; ok {
		panic("lap already recorded")
	}
	sw.laps[name] = time.Since(sw.t1)
	sw.t1 = time.Now()
}

//Laps returns current measured times
func (sw *Stopwatch) Laps() map[string]time.Duration {
	return sw.laps
}

//String returns json marshaled laps
func (sw *Stopwatch) String() string {
	b, _ := json.Marshal(sw.laps)
	return "Stopwatch: " + string(b)
}
