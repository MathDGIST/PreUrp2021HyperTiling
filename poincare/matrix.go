package poincare

type Matrix [][]complex128

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

func Mobius(m Matrix, z complex128) complex128 {

	return 0
}
