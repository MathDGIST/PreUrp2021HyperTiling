package main

import (
	"./drawing"
	"./poincare"
	"math"
)

func main() {
	// Drawing Poincare Disc
	c := drawing.NewCanvas()
	t := make([]float64, 101)
	for i := 0 ; i < 101 ; i ++ {
		t[i] = 2*math.Pi * float64(i) / float64(100)
	}
	xy := make([][]float64, 2)
	xy[0] = make([]float64, 101)
	xy[1] = make([]float64, 101)
	for i := 0 ; i < 101 ; i ++ {
		xy[0][i] = math.Cos(t[i])
		xy[1][i] = math.Sin(t[i])
	}
	c.AddLine(xy)
	c.AddPoint([]float64{1.25, 1.25})
	c.AddPoint([]float64{-1.25, -1.25})

	// Test drawing a gedesic
	r := float64(1)
	z := complex(1,1)
	t1, t2 := math.Pi, 3*math.Pi/float64(2)
	g := poincare.NewGeodesic(z, r, t1, t2)
	xy = g.XY()
	c.AddLine(xy)

	m := [][]complex128{{1, -2}, {0, 1}}
    m = poincare.MatMul(poincare.H2D(), m)
    m = poincare.MatMul(m, poincare.D2H())
    for i := 0 ; i < len(xy[0]) ; i ++ {
    	g[i] = poincare.Mobius(m, g[i])
	}
	c.Draw("./test.png")
}
