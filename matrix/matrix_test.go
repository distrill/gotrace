package matrix

import (
	"testing"

	"github.com/distrill/gotrace/tuples"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	Scenario: Constructing and inspecting a 4x4 matrix
	Given the following 4x4 matrix M:
		| 1    | 2    | 3    | 4    |
		| 5.5  | 6.5  | 7.5  | 8.5  |
		| 9    | 10   | 11   | 12   |
		| 13.5 | 14.5 | 15.5 | 16.5 |
	Then M[0,0] = 1
	And  M[0,3] = 4
	And  M[1,0] = 5.5
	And  M[1,2] = 7.5
	And  M[2,2] = 11
	And  M[3,0] = 13.5
	And  M[3,2] = 15.5
*/

func TestNew4x4Matrix(t *testing.T) {
	m := Matrix{
		Row{1, 2, 3, 4},
		Row{5.5, 6.5, 7.5, 8.5},
		Row{9, 10, 11, 12},
		Row{13.5, 14.5, 15.5, 16.5},
	}
	assert.Equal(t, m[0][0], 1.0)
	assert.Equal(t, m[0][3], 4.0)
	assert.Equal(t, m[1][0], 5.5)
	assert.Equal(t, m[1][2], 7.5)
	assert.Equal(t, m[2][2], 11.0)
	assert.Equal(t, m[3][0], 13.5)
	assert.Equal(t, m[3][2], 15.5)
}

/*
	Scenario: A 2x2 matrix ought to be representable
	Given the following 2x2 matrix M:
		| -3 | 5  |
		| 1  | -2 |
	Then M[0,0] = -3
	And  M[0,1] = 5
	And  M[1,0] = 1
	And  M[1,1] = -2
*/
func TestNew2x2Matrix(t *testing.T) {
	m := Matrix{
		Row{-3, 5},
		Row{1, -2},
	}
	assert.Equal(t, m[0][0], -3.0)
	assert.Equal(t, m[0][1], 5.0)
	assert.Equal(t, m[1][0], 1.0)
	assert.Equal(t, m[1][1], -2.0)
}

/*
	Scenario: A 3x3 matrix ought to be representable
	Given the following 3x3 matrix M:
		| -3 | 5  | 0  |
		| 1  | -2 | -7 |
		| 0  | 1  | 1  |
	Then M[0,0] = -3
	And M[1,1] = -2
	And M[2,2] = 1
*/
func TestNew3x3Matrix(t *testing.T) {
	m := Matrix{
		Row{-3, 5, 0},
		Row{1, -2, -7},
		Row{0, 1, 1},
	}
	assert.Equal(t, m[0][0], -3.0)
	assert.Equal(t, m[1][1], -2.0)
	assert.Equal(t, m[2][2], 1.0)
}

/*
	Scenario: Matrix equality with identical matrices
	Given the following matrix A:
		| 1 | 2 | 3 | 4 |
		| 5 | 6 | 7 | 8 |
		| 9 | 8 | 7 | 6 |
		| 5 | 4 | 3 | 2 |
	And the following matrix B:
		| 1 | 2 | 3 | 4 |
		| 5 | 6 | 7 | 8 |
		| 9 | 8 | 7 | 6 |
		| 5 | 4 | 3 | 2 |
	Then A = B
*/
func TestIdenticalMatrixEquality(t *testing.T) {
	m1 := Matrix{
		Row{1, 2, 3, 4},
		Row{5, 6, 7, 8},
		Row{9, 8, 7, 6},
		Row{5, 4, 3, 2},
	}
	m2 := Matrix{
		Row{1, 2, 3, 4},
		Row{5, 6, 7, 8},
		Row{9, 8, 7, 6},
		Row{5, 4, 3, 2},
	}
	assert.Equal(t, m1, m2)
}

/*
	Scenario: Matrix equality with different matrices
	Given the following matrix A:
		| 1 | 2 | 3 | 4 |
		| 5 | 6 | 7 | 8 |
		| 9 | 8 | 7 | 6 |
		| 5 | 4 | 3 | 2 |
	And the following matrix B:
		| 2 | 3 | 4 | 5 |
		| 6 | 7 | 8 | 9 |
		| 8 | 7 | 6 | 5 |
		| 4 | 3 | 2 | 1 |
	Then A != B
*/
func TestNonidenticalMatrixInequality(t *testing.T) {
	m1 := Matrix{
		Row{1, 2, 3, 4},
		Row{5, 6, 7, 8},
		Row{9, 8, 7, 6},
		Row{5, 4, 3, 2},
	}
	m2 := Matrix{
		Row{2, 3, 4, 5},
		Row{6, 7, 8, 9},
		Row{8, 7, 6, 5},
		Row{4, 3, 2, 1},
	}
	assert.NotEqual(t, m1, m2)
}

/*
	Scenario: Multiplying two matrices
	Given the following matrix A:
		| 1 | 2 | 3 | 4 |
		| 5 | 6 | 7 | 8 |
		| 9 | 8 | 7 | 6 |
		| 5 | 4 | 3 | 2 |
	And the following matrix B:
		| -2 | 1 | 2 | 3  |
		| 3  | 2 | 1 | -1 |
		| 4  | 3 | 6 | 5  |
		| 1  | 2 | 7 | 8  |
	Then A * B is the following 4x4 matrix:
		| 20| 22 | 50  | 48  |
		| 44| 54 | 114 | 108 |
		| 40| 58 | 110 | 102 |
		| 16| 26 | 46  | 42  |
*/
func TestMulTwoMatrices(t *testing.T) {
	m1 := Matrix{
		Row{1, 2, 3, 4},
		Row{5, 6, 7, 8},
		Row{9, 8, 7, 6},
		Row{5, 4, 3, 2},
	}
	m2 := Matrix{
		Row{-2, 1, 2, 3},
		Row{3, 2, 1, -1},
		Row{4, 3, 6, 5},
		Row{1, 2, 7, 8},
	}

	actual, err := m1.MulM(m2)
	expected := Matrix{
		Row{20, 22, 50, 48},
		Row{44, 54, 114, 108},
		Row{40, 58, 110, 102},
		Row{16, 26, 46, 42},
	}
	require.Nil(t, err)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: A matrix multiplied by a tuple
	Given the following matrix A:
		| 1 | 2 | 3 | 4 |
		| 2 | 4 | 4 | 2 |
		| 8 | 6 | 4 | 1 |
		| 0 | 0 | 0 | 1 |
	And b ← tuple(1, 2, 3, 1)
	Then A * b = tuple(18, 24, 33, 1)
*/
func TestMulMatrixTuple(t *testing.T) {
	A := Matrix{
		Row{1, 2, 3, 4},
		Row{2, 4, 4, 2},
		Row{8, 6, 4, 1},
		Row{0, 0, 0, 1},
	}
	b := tuples.Tuple{1, 2, 3, 1}
	actual, err := A.MulT(b)
	require.Nil(t, err)
	expected := tuples.Tuple{18, 24, 33, 1}
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Multiplying a matrix by the identity matrix
	Given the following matrix A:
		| 0 | 1 | 2  | 4  |
		| 1 | 2 | 4  | 8  |
		| 2 | 4 | 8  | 16 |
		| 4 | 8 | 16 | 32 |
	Then A * identity_matrix = A
*/
func TestMulMIdentity(t *testing.T) {
	A := Matrix{
		Row{0, 1, 2, 4},
		Row{1, 2, 4, 8},
		Row{2, 4, 8, 16},
		Row{4, 8, 16, 32},
	}
	I := NewIdentityMatrix(4)
	actual, err := A.MulM(I)
	require.Nil(t, err)
	assert.Equal(t, A, actual)
}

/*
	Scenario: Multiplying the identity matrix by a tuple
	Given a ← tuple(1, 2, 3, 4)
	Then identity_matrix * a = a
*/
func TestMulTIdentity(t *testing.T) {
	I := NewIdentityMatrix(4)
	a := tuples.Tuple{1, 2, 3, 4}
	actual, err := I.MulT(a)
	require.Nil(t, err)
	assert.Equal(t, a, actual)
}

/*
	Scenario: Transposing a matrix
	Given the following matrix A:
		| 0 | 9 | 3 | 0 |
		| 9 | 8 | 0 | 8 |
		| 1 | 8 | 5 | 3 |
		| 0 | 0 | 5 | 8 |
	Then transpose(A) is the following matrix:
		| 0 | 9 | 1 | 0 |
		| 9 | 8 | 8 | 0 |
		| 3 | 0 | 5 | 5 |
		| 0 | 8 | 3 | 8 |
*/
func TestTransposeMatrix(t *testing.T) {
	A := Matrix{
		Row{0, 9, 3, 0},
		Row{9, 8, 0, 8},
		Row{1, 8, 5, 3},
		Row{0, 0, 5, 8},
	}
	actual := A.Transpose()
	expected := Matrix{
		Row{0, 9, 1, 0},
		Row{9, 8, 8, 0},
		Row{3, 0, 5, 5},
		Row{0, 8, 3, 8},
	}
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Transposing the identity matrix
	Given A ← transpose(identity_matrix)
	Then A = identity_matrix
*/
func TestTransposeIdentity(t *testing.T) {
	I := NewIdentityMatrix(4)
	A := I.Transpose()
	assert.Equal(t, A, I)
}
