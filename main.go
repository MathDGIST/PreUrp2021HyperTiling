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
	//r := float64(1)
	//z := complex(1,1)
	//t1, t2 := math.Pi, 3*math.Pi/float64(2)
	//g := poincare.NewGeodesic(z, r, t1, t2)
	//xy = g.XY()
	//c.AddLine(xy)
	//
	//m := [][]complex128{{1, -1}, {0, 1}}
    //m = poincare.MatMul(poincare.H2D(), m)
    //m = poincare.MatMul(m, poincare.D2H())
    //for i := 0 ; i < len(xy[0]) ; i ++ {
    //	g[i] = poincare.Mobius(m, g[i])
	//}
	//xy = g.XY()
	//c.AddLine(xy)

	z := complex(math.Sqrt(6)/2, math.Sqrt(2)/2)
	r := float64(1)
	t1 := math.Pi + math.Pi / float64(12)
	t2 := math.Pi + math.Pi / float64(4)
	g := poincare.NewGeodesic(z, r, t1, t2)
	xy = g.XY()
	c.AddLine(xy)

	for j := 0 ; j < 5 ; j ++ {
		m := poincare.Rotation(math.Pi/6)
		m = poincare.MatMul(poincare.H2D(), m)
		m = poincare.MatMul(m, poincare.D2H())
		for i := 0 ; i < len(g) ; i ++ {
			g[i] = poincare.Mobius(m, g[i])
		}
		xy = g.XY()
		c.AddLine(xy)
	}

	g = poincare.NewGeodesic(z, r, t1, t2)
	z = g[0]
	zz := poincare.Mobius(poincare.D2H(), z)
	xx, yy := real(zz), imag(zz)
	s := make([][]complex128, 2)
	for i := 0 ; i < 2 ; i ++ {
		s[i] = make([]complex128, 2)
	}
	s[0][0] = complex(float64(1)/math.Sqrt(yy),0)
	s[0][1] = complex(-xx/math.Sqrt(yy),0)
	s[1][0] = complex(0,0)
	s[1][1] = complex(math.Sqrt(yy),0)
	m := poincare.MatMul(poincare.H2D(), poincare.Inverse(s))
	m = poincare.MatMul(m, poincare.Rotation(math.Pi/4))
	m = poincare.MatMul(m, s)
	m = poincare.MatMul(m, poincare.D2H())

	for i := 0 ; i < len(g) ; i ++ {
		g[i] = poincare.Mobius(m, g[i])
	}
	xy = g.XY()
	c.AddLine(xy)

	c.Draw("./test.png")
}
