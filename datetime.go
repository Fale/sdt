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
		case delta < 864000:
			// delta is less than 10 days
			// - mayor: every day (min: 0, max: 10)
			// - minor: every day (min: 0, max: 10)
			tick.Label = tm.Format(t.Format)
			ticks = append(ticks, tick)
		case delta < 7776000:
			// delta is between 10 and 90 days
			// - mayor: every 5 days (min: 2, max: 18)
			// - minor: every day (min: 10, max: 90)
			if c%5 == 0 {
				tick.Label = tm.Format(t.Format)
			}
			ticks = append(ticks, tick)
		case delta < 15552000:
			// delta is between 90 and 180 days
			// mayor: on day 1 and 15 of every month (min: 5, max: 12)
			// minor: on day 1, 5, 10, 15, 20, 25, 30 of every month (min: 17, max: 36)
			if tm.Day() == 1 || tm.Day() == 15 {
				tick.Label = tm.Format(t.Format)
			}
			if tm.Day() == 1 || tm.Day()%5 == 0 {
				ticks = append(ticks, tick)
			}
		case delta < 47347200:
			// delta is between 6 months and 18 months
			// mayor: on day 1 of every month (min: 5, max: 18)
			// minor: on day 1 and 15 of every month (min: 11, max: 36)
			if tm.Day() == 1 {
				tick.Label = tm.Format(t.Format)
			}
			if tm.Day() == 1 || tm.Day() == 15 {
				ticks = append(ticks, tick)
			}
		default:
			// delta is higher than 18 months
			// mayor: on the 1st of january (min: 1, max: inf.)
			// minor: on day 1 of every month (min: 17, max inf.)
			if tm.Day() == 1 && tm.Month() == time.January {
				tick.Label = tm.Format(t.Format)
			}
			if tm.Day() == 1 {
				ticks = append(ticks, tick)
			}
		}
		c = c + 1
		tm = tm.AddDate(0, 0, 1)
		if tm.After(time.Unix(int64(max), 0)) {
			break
		}
	}
	return ticks
}
