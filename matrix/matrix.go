package matrix

import (
	"fmt"

	"github.com/distrill/gotrace/tuples"
)

// Matrix - indexed row/column (NOTE this is the opposite of canvas)
type Matrix []Row

// Row - matrix row
type Row []float64

func NewMatrix(m, n int) Matrix {
	matrix := make([]Row, m)
	for i := range matrix {
		matrix[i] = make(Row, n)
	}
	return matrix
}

func NewSquareMatrix(m int) Matrix {
	return NewMatrix(m, m)
}

func NewIdentityMatrix(m int) Matrix {
	mat := NewSquareMatrix(m)
	for i := 0; i < m; i++ {
		mat[i][i] = 1
	}
	return mat
}

// Mul - multiple a matrix by another matrix
func (m Matrix) MulM(o Matrix) (Matrix, error) {
	if len(m[0]) != len(o) {
		return nil, fmt.Errorf("LH matrix width must equal RH matrix height")
	}

	result := NewMatrix(len(m), len(m[0]))

	for r := 0; r < len(m); r++ {
		for c := 0; c < len(o[0]); c++ {
			e := 0.0
			for i := 0; i < len(m[0]); i++ {
				e += m[r][i] * o[i][c]
			}
			result[r][c] = e
		}
	}

	return result, nil
}

// MulT - multiply a matrix by a tuple
func (m Matrix) MulT(t tuples.Tuple) (tuples.Tuple, error) {
	o := Matrix{
		Row{t.X},
		Row{t.Y},
		Row{t.Z},
		Row{t.W},
	}
	r, err := m.MulM(o)
	if err != nil {
		// FIXME I do not like this, should return nil, but that makes
		// this function signature inconsistent from the rest of the project
		return tuples.Tuple{}, err
	}
	return tuples.Tuple{r[0][0], r[1][0], r[2][0], r[3][0]}, nil
}

func (m Matrix) Transpose() Matrix {
	result := NewMatrix(len(m[0]), len(m))
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			result[r][c] = m[c][r]
		}
	}
	return result
}
