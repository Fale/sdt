package sdt

import (
	"testing"

	"gonum.org/v1/plot"
)

type ExpectedResults struct {
	Min   float64
	Max   float64
	Ticks []plot.Tick
}

func TestAbs(t *testing.T) {
	ers := []ExpectedResults{
		{Min: 1622505600, Max: 1622592000, Ticks: []plot.Tick{{Value: 1622505600, Label: "01/06"}, {Value: 162259200, Label: "02/06"}}}, // 2021-06-01 -> 2021-06-02
	}
	for _, er := range ers {
		ti := Ticks{Format: "02/01"}
		ticks := ti.Ticks(er.Min, er.Max)
		if len(ticks) != len(er.Ticks) {
			t.Errorf("tickers(%f,%f) length = %d; want %d", er.Min, er.Max, len(ticks), len(er.Ticks))
		}
	}
}
