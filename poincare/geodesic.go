package poincare

import "math"

type Geodesic []complex128

func NewGeodesic(c complex128, r, t1, t2 float64) Geodesic {
	g := make([]complex128, 100)

	for i := 0 ; i < 100 ; i ++ {
		t := t1 + (t2-t1)/99 * float64(i)
		g[i] = c + complex(r*math.Cos(t), r*math.Sin(t))
	}
	return g
}

func (g Geodesic) XY() [][]float64 {
	xy := make([][]float64, 2)
	xy[0] = make([]float64, len(g))
	xy[1] = make([]float64, len(g))
	for i := 0 ; i < len(g) ; i ++ {
		xy[0][i], xy[1][i] = real(g[i]), imag(g[i])
	}
	return xy
}