package main

import (
	"time"

	"github.com/fale/sdt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()

	line := plotter.XYs{
		plotter.XY{X: float64(time.Date(2021, time.June, 10, 0, 0, 0, 0, time.UTC).Unix()), Y: float64(4)},
		plotter.XY{X: float64(time.Date(2021, time.June, 11, 0, 0, 0, 0, time.UTC).Unix()), Y: float64(1)},
		plotter.XY{X: float64(time.Date(2021, time.June, 12, 0, 0, 0, 0, time.UTC).Unix()), Y: float64(8)},
		plotter.XY{X: float64(time.Date(2021, time.June, 13, 0, 0, 0, 0, time.UTC).Unix()), Y: float64(3)},
		plotter.XY{X: float64(time.Date(2021, time.June, 14, 0, 0, 0, 0, time.UTC).Unix()), Y: float64(5)},
	}

	p.Add(plotter.NewGrid())
	if err := plotutil.AddLinePoints(p, "First", line); err != nil {
		return
	}

	p.X.Tick.Marker = sdt.Ticks{Format: "02/01"}

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "points.png"); err != nil {
		return
	}
}
