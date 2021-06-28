package sdt

import (
	"time"

	"gonum.org/v1/plot"
)

type Ticks struct {
	Ticker plot.Ticker
	Format string
	Time   func(t float64) time.Time
}

func (t Ticks) Ticks(min, max float64) []plot.Tick {
	if t.Format == "" {
		t.Format = time.RFC3339
	}
	ticks := []plot.Tick{}
	tm := time.Unix(int64(min), 0)
	c := 0
	for {
		tick := plot.Tick{Value: float64(tm.Unix())}
		switch delta := max - min; {
		case delta < 864000: // 10 days
			tick.Label = tm.Format(t.Format)
		case delta < 7776000: // 90 days
			if c%5 == 0 {
				tick.Label = tm.Format(t.Format)
			}
		case delta < 15552000: // 180 days
			if tm.Day() == 1 || tm.Day() == 15 {
				tick.Label = tm.Format(t.Format)
			}
		default:
			if tm.Day() == 1 {
				tick.Label = tm.Format(t.Format)
			}
		}
		c = c + 1
		ticks = append(ticks, tick)
		tm = tm.AddDate(0, 0, 1)
		if tm.After(time.Unix(int64(max), 0)) {
			break
		}
	}
	return ticks
}
