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

func (A Matrix) String() string {
	str := "\n---------------------------\n"
	for _, row := range A {
		str += fmt.Sprintf("%v\n", row)
	}
	str += "---------------------------\n"
	return str
}

// Mul - multiple a matrix by another matrix
func (A Matrix) MulM(B Matrix) (Matrix, error) {
	if len(A[0]) != len(B) {
		return nil, fmt.Errorf("LH matrix width must equal RH matrix height")
	}

	result := NewMatrix(len(A), len(A[0]))

	for r := 0; r < len(A); r++ {
		for c := 0; c < len(B[0]); c++ {
			e := 0.0
			for i := 0; i < len(A[0]); i++ {
				e += A[r][i] * B[i][c]
			}
			result[r][c] = e
		}
	}

	return result, nil
}

func (A Matrix) MustMulM(B Matrix) Matrix {
	AB, err := A.MulM(B)
	if err != nil {
		panic(err)
	}
	return AB
}

// MulT - multiply a matrix by a tuple
func (A Matrix) MulT(t tuples.Tuple) (tuples.Tuple, error) {
	B := Matrix{
		Row{t.X},
		Row{t.Y},
		Row{t.Z},
		Row{t.W},
	}
	r, err := A.MulM(B)
	if err != nil {
		// FIXME I do not like this, should return nil, but that makes
		// this function signature inconsistent from the rest of the project
		return tuples.Tuple{}, err
	}
	return tuples.Tuple{r[0][0], r[1][0], r[2][0], r[3][0]}, nil
}

func (A Matrix) MustMulT(t tuples.Tuple) tuples.Tuple {
	At, err := A.MulT(t)
	if err != nil {
		panic(err)
	}
	return At
}

// MulS - multiply a matrix by a scalar
func (A Matrix) MulS(s float64) (Matrix, error) {
	if len(A) < 1 {
		return nil, fmt.Errorf("matrix must have length at least 1 for multiplication")
	}

	B := NewMatrix(len(A), len(A[0]))
	for r := 0; r < len(A); r++ {
		for c := 0; c < len(A[0]); c++ {
			B[r][c] = A[r][c] * s
		}
	}
	return B, nil
}

// DivS - divide a matrix by a scalr
func (A Matrix) DivS(s float64) (Matrix, error) {
	return A.MulS(1 / s)
}

func (A Matrix) Transpose() (Matrix, error) {
	if len(A) < 1 {
		return nil, fmt.Errorf("matrix must have length at least 1 for multiplication")
	}

	B := NewMatrix(len(A[0]), len(A))
	for r := 0; r < len(A); r++ {
		for c := 0; c < len(A[0]); c++ {
			B[r][c] = A[c][r]
		}
	}
	return B, nil
}

func (A Matrix) Determinant() (float64, error) {
	if len(A) < 2 || len(A[0]) < 2 {
		return 0, fmt.Errorf("invalid matrix dimensions for determinant %v x %v", len(A), len(A[0]))
	}

	if len(A) != len(A[0]) {
		return 0, fmt.Errorf("matrix must be square to calculate determinant, got %v x %v", len(A), len(A[0]))
	}

	// 2x2 case is different from the general case
	if len(A) == 2 {
		return (A[0][0] * A[1][1]) - (A[0][1] * A[1][0]), nil
	}

	// general case for larger than 2x2
	det := 0.0
	for r := 0; r < len(A[0]); r++ {
		cof, err := A.Cofactor(0, r)
		if err != nil {
			return 0, err
		}
		det += A[0][r] * cof
	}
	return det, nil
}

func (A Matrix) Submatrix(m, n int) (Matrix, error) {
	if m < 0 || m >= len(A) {
		return nil, fmt.Errorf("invalid index %v for len %v", m, len(A))
	}
	if n < 0 || n >= len(A[0]) {
		return nil, fmt.Errorf("invalid index %v for len %v", n, len(A[0]))
	}
	sub := Matrix{}
	for r := range A {
		if r != m {
			row := Row{}
			for c := range A[r] {
				if c != n {
					row = append(row, A[r][c])
				}
			}
			sub = append(sub, row)
		}
	}
	return sub, nil
}

// Minor - determinant of the submatrix at m, n
func (A Matrix) Minor(m, n int) (float64, error) {
	sub, err := A.Submatrix(m, n)
	if err != nil {
		return 0, err
	}
	// d, _ := sub.Determinant()
	// fmt.Printf("%v %v (%v) %v", m, n, d, sub)
	return sub.Determinant()
}

// Cofactor - minor, negated if m + n = odd number
func (A Matrix) Cofactor(m, n int) (float64, error) {
	min, err := A.Minor(m, n)
	if err != nil {
		return 0, err
	}
	if (m+n)%2 == 0 {
		return min, nil
	}
	return -min, nil
}

func (A Matrix) Invertible() (bool, error) {
	det, err := A.Determinant()
	if err != nil {
		return false, err
	}
	return det != 0, nil
}

func (A Matrix) Inverse() (Matrix, error) {
	inv, err := A.Invertible()
	if err != nil {
		return nil, err
	}
	if !inv {
		return nil, fmt.Errorf("cannot invert non-invertible matrix")
	}

	det, err := A.Determinant()
	if err != nil {
		return nil, err
	}

	B := NewMatrix(len(A), len(A[0]))
	for r := range A {
		for c := range A[0] {
			cof, err := A.Cofactor(r, c)
			if err != nil {
				return nil, err
			}
			// swap row and column here to get transpose
			B[c][r] = cof / det
		}
	}

	return B, nil
}
