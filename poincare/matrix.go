package poincare

import (
	"math"
)

type Matrix [][]complex128

func Inverse(m Matrix) Matrix {
	n := make([][]complex128, 2)
	for i := 0 ; i < 2 ; i ++ {
		n[i] = make([]complex128, 2)
	}
	det := m[0][0]*m[1][1] - m[1][0]*m[0][1]
	n[0][0] = m[1][1] / det
	n[0][1] = -m[0][1] / det
	n[1][0] = -m[1][0] / det
	n[1][1] = m[0][0] / det
	return n
}

func D2H() Matrix {
	m := make([][]complex128, 2)
	for i := 0 ; i < 2 ; i ++ {
		m[i] = make([]complex128, 2)
	}
	m[0][0] = complex(0,1)
	m[0][1] = complex(-1,0)
	m[1][0] = complex(-1,0)
	m[1][1] = complex(0,1)
	return m
}

func H2D() Matrix {
	m := make([][]complex128, 2)
	for i := 0 ; i < 2 ; i ++ {
		m[i] = make([]complex128, 2)
	}
	m[0][0] = complex(0,1)
	m[0][1] = complex(1,0)
	m[1][0] = complex(1,0)
	m[1][1] = complex(0,1)
	return m
}

func MatMul(m1, m2 Matrix) Matrix {
	r := len(m1)
	c := len(m2)
	m := make([][]complex128, r)
	for i := 0 ; i < c ; i ++ {
		m[i] = make([]complex128, c)
	}
	for i := 0 ; i < r ; i ++ {
		for j := 0 ; j < c ; j ++ {
			var sum complex128 = 0
			for k := 0 ; k < len(m[0]) ; k ++ {
				sum += m1[i][k]*m2[k][j]
			}
			m[i][j] = sum
		}
	}
	return m
}

func Rotation(t float64) Matrix {
	m := make([][]complex128, 2)
	for i := 0 ; i < 2 ; i ++ {
		m[i] = make([]complex128, 2)
	}
	m[0][0] = complex(math.Cos(t),0)
	m[0][1] = complex(math.Sin(t),0)
	m[1][0] = complex(-math.Sin(t),0)
	m[1][1] = complex(math.Cos(t), 0)
	return m
}

func Mobius(m Matrix, z complex128) complex128 {
	return (m[0][0]*z + m[0][1]) / (m[1][0]*z + m[1][1])
}
