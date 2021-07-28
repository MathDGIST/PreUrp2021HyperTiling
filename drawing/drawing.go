package drawing

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
)

type Canvas struct {
	P *plot.Plot
}

func NewCanvas() *Canvas {
	p := plot.New()
	p.HideAxes()
	return &Canvas{
		P: p,
	}
}

func (c *Canvas) Draw(fn string) {
	c.P.Save(800, 800, fn)
}

func (c *Canvas) AddPoint(xy []float64) {
	pts := make(plotter.XYs, 1)
	pts[0].X = xy[0]
	pts[0].Y = xy[1]
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	scatter.GlyphStyle.Radius = vg.Points(0)
	c.P.Add(scatter)
}

func (c *Canvas) AddLine(xy [][]float64) {
	if len(xy[0]) != len(xy[1]) {
		panic("Invalid xy array!")
	}
	l := len(xy[0])
	pts := make(plotter.XYs, l)
	for i := 0 ; i < l ; i ++ {
		pts[i].X = xy[0][i]
		pts[i].Y = xy[1][i]
	}
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Panic(err)
	}
	c.P.Add(line)
}