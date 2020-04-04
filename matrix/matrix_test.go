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
	actual, err := A.Transpose()
	require.Nil(t, err)
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
	A, err := I.Transpose()
	require.Nil(t, err)
	assert.Equal(t, A, I)
}

/*
	Scenario: Calculating the determinant of a 2x2 matrix
	Given the following 2x2 matrix A:
		| 1  | 5 |
		| -3 | 2 |
	Then determinant(A) = 17
*/
func TestDeterminant2x2(t *testing.T) {
	A := Matrix{
		Row{1, 5},
		Row{-3, 2},
	}
	actual, err := A.Determinant()
	require.Nil(t, err)
	assert.Equal(t, 17.0, actual)
}

/*
Scenario: A submatrix of a 3x3 matrix is a 2x2 matrix
Given the following 3x3 matrix A:
	| 1  | 5 | 0 |
	| -3 | 2 | 7 |
	| 0  | 6 | -3 |
Then submatrix(A, 0, 2) is the following 2x2 matrix:
	| -3 | 2 |
	| 0  | 6 |
*/
func Test3x3Submatrix(t *testing.T) {
	A := Matrix{
		Row{1, 5, 0},
		Row{-3, 2, 7},
		Row{0, 6, -3},
	}
	actual, err := A.Submatrix(0, 2)
	require.Nil(t, err)
	expected := Matrix{
		Row{-3, 2},
		Row{0, 6},
	}
	assert.Equal(t, expected, actual)
}

/*
Scenario: A submatrix of a 4x4 matrix is a 3x3 matrix
Given the following 4x4 matrix A:
	| -6 | 1 | 1  | 6 |
	| -8 | 5 | 8  | 6 |
	| -1 | 0 | 8  | 2 |
	| -7 | 1 | -1 | 1 |
Then submatrix(A, 2, 1) is the following 3x3 matrix:
	| -6 | 1 | 6  |
	| -8 | 8 | 6  |
	| -7 | -1 | 1 |
*/
func Test4x4Submatrix(t *testing.T) {
	A := Matrix{
		Row{1, 5, 0},
		Row{-3, 2, 7},
		Row{0, 6, -3},
	}
	actual, err := A.Submatrix(0, 2)
	require.Nil(t, err)
	expected := Matrix{
		Row{-3, 2},
		Row{0, 6},
	}
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Calculating a minor of a 3x3 matrix
	Given the following 3x3 matrix A:
		| 3 | 5  | 0  |
		| 2 | -1 | -7 |
		| 6 | -1 | 5  |
	And B ← submatrix(A, 1, 0)
	Then determinant(B) = 25
	And minor(A, 1, 0) = 25
*/
func TestMinor3x3(t *testing.T) {
	A := Matrix{
		Row{3, 5, 0},
		Row{2, -1, -7},
		Row{6, -1, 5},
	}
	B, err := A.Submatrix(1, 0)
	require.Nil(t, err)

	det, err := B.Determinant()
	require.Nil(t, err)
	assert.Equal(t, 25.0, det)

	min, err := A.Minor(1, 0)
	require.Nil(t, err)
	assert.Equal(t, 25.0, min)
}

/*
	Scenario: Calculating a cofactor of a 3x3 matrix
	Given the following 3x3 matrix A:
		| 3 | 5  | 0  |
		| 2 | -1 | -7 |
		| 6 | -1 | 5  |
	Then minor(A, 0, 0) = -12
	And cofactor(A, 0, 0) = -12
	And minor(A, 1, 0) = 25
	And cofactor(A, 1, 0) = -25
*/
func TestCofactor3x3(t *testing.T) {
	A := Matrix{
		Row{3, 5, 0},
		Row{2, -1, -7},
		Row{6, -1, 5},
	}

	min, err := A.Minor(0, 0)
	require.Nil(t, err)
	assert.Equal(t, -12.0, min)

	cof, err := A.Cofactor(0, 0)
	require.Nil(t, err)
	assert.Equal(t, -12.0, cof)

	min, err = A.Minor(1, 0)
	require.Nil(t, err)
	assert.Equal(t, 25.0, min)

	cof, err = A.Cofactor(1, 0)
	require.Nil(t, err)
	assert.Equal(t, -25.0, cof)
}

/*
	Scenario: Calculating the determinant of a 3x3 matrix
	Given the following 3x3 matrix A:
		| 1  | 2 | 6  |
		| -5 | 8 | -4 |
		| 2  | 6 | 4  |
	Then cofactor(A, 0, 0) = 56
	And cofactor(A, 0, 1) = 12
	And cofactor(A, 0, 2) = -46
	And determinant(A) = -196
*/
func TestDeterminant3x3(t *testing.T) {
	A := Matrix{
		Row{1, 2, 6},
		Row{-5, 8, -4},
		Row{2, 6, 4},
	}

	cof, err := A.Cofactor(0, 0)
	require.Nil(t, err)
	assert.Equal(t, 56.0, cof)

	cof, err = A.Cofactor(0, 1)
	require.Nil(t, err)
	assert.Equal(t, 12.0, cof)

	cof, err = A.Cofactor(0, 2)
	require.Nil(t, err)
	assert.Equal(t, -46.0, cof)

	det, err := A.Determinant()
	require.Nil(t, err)
	assert.Equal(t, -196.0, det)
}

/*
	Scenario: Calculating the determinant of a 4x4 matrix
	Given the following 4x4 matrix A:
		| -2 | -8 | 3  | 5  |
		| -3 | 1  | 7  | 3  |
		| 1  | 2  | -9 | 6  |
		| -6 | 7  | 7  | -9 |
	Then cofactor(A, 0, 0) = 690
	And cofactor(A, 0, 1) = 447
	And cofactor(A, 0, 2) = 210
	And cofactor(A, 0, 3) = 51
	And determinant(A) = -4071
*/
func TestDeterminant4x4(t *testing.T) {
	A := Matrix{
		Row{-2, -8, 3, 5},
		Row{-3, 1, 7, 3},
		Row{1, 2, -9, 6},
		Row{-6, 7, 7, -9},
	}

	cof, err := A.Cofactor(0, 0)
	require.Nil(t, err)
	assert.Equal(t, 690.0, cof)

	cof, err = A.Cofactor(0, 1)
	require.Nil(t, err)
	assert.Equal(t, 447.0, cof)

	cof, err = A.Cofactor(0, 2)
	require.Nil(t, err)
	assert.Equal(t, 210.0, cof)

	cof, err = A.Cofactor(0, 3)
	require.Nil(t, err)
	assert.Equal(t, 51.0, cof)

	det, err := A.Determinant()
	require.Nil(t, err)
	assert.Equal(t, -4071.0, det)
}

/*
	Scenario: Testing an invertible matrix for invertibility
	Given the following 4x4 matrix A:
		| 6 | 4  | 4 | 4  |
		| 5 | 5  | 7 | 6  |
		| 4 | -9 | 3 | -7 |
		| 9 | 1  | 7 | -6 |
	Then determinant(A) = -2120
	And A is invertible
*/
func TestInvertibleTrue(t *testing.T) {
	A := Matrix{
		Row{6, 4, 4, 4},
		Row{5, 5, 7, 6},
		Row{4, -9, 3, -7},
		Row{9, 1, 7, -6},
	}

	det, err := A.Determinant()
	require.Nil(t, err)
	assert.Equal(t, -2120.0, det)

	inv, err := A.Invertible()
	require.Nil(t, err)
	assert.True(t, inv)
}

/*
	Scenario: Testing a noninvertible matrix for invertibility
	Given the following 4x4 matrix A:
		| -4 | 2  | -2 | -3 |
		| 9  | 6  | 2  | 6  |
		| 0  | -5 | 1  | -5 |
		| 0  | 0  | 0  | 0  |
	Then determinant(A) = 0
	And A is not invertible
*/
func TestInvertibleFalse(t *testing.T) {
	A := Matrix{
		Row{-4, 2, -2, -3},
		Row{9, 6, 2, 6},
		Row{0, -5, 1, -5},
		Row{0, 0, 0, 0},
	}

	det, err := A.Determinant()
	require.Nil(t, err)
	assert.Equal(t, 0.0, det)

	inv, err := A.Invertible()
	require.Nil(t, err)
	assert.False(t, inv)
}

/*
	Scenario: Calculating the inverse of a matrix
	Given the following 4x4 matrix A:
		| -5 | 2  | 6  | -8 |
		| 1  | -5 | 1  | 8  |
		| 7  | 7  | -6 | -7 |
		| 1  | -3 | 7  | 4  |
	And B ← inverse(A)
	Then determinant(A) = 532
	And cofactor(A, 2, 3) = -160
	And B[3,2] = -160/532
	And cofactor(A, 3, 2) = 105
	And B[2,3] = 105/532
	And B is the following 4x4 matrix:
		| 0.21805  | 0.45113  | 0.24060  | -0.04511 |
		| -0.80827 | -1.45677 | -0.44361 | 0.52068  |
		| -0.07895 | -0.22368 | -0.05263 | 0.19737  |
		| -0.52256 | -0.81391 | -0.30075 | 0.30639  |
*/
func TestInverse(t *testing.T) {
	A := Matrix{
		Row{-5, 2, 6, -8},
		Row{1, -5, 1, 8},
		Row{7, 7, -6, -7},
		Row{1, -3, 7, 4},
	}
	B, err := A.Inverse()
	require.Nil(t, err)

	// fmt.Println(A)
	// cof, err := A.Cofactor(1, 1)
	// require.Nil(t, err)
	// fmt.Println(cof)
	// fmt.Println(A)

	det, err := A.Determinant()
	require.Nil(t, err)
	assert.Equal(t, 532.0, det)

	cof, err := A.Cofactor(2, 3)
	require.Nil(t, err)
	assert.Equal(t, -160.0, cof)
	assert.Equal(t, -160/532.0, B[3][2])

	cof, err = A.Cofactor(3, 2)
	require.Nil(t, err)
	assert.Equal(t, 105.0, cof)
	assert.Equal(t, 105/532.0, B[2][3])

	expected := Matrix{
		Row{0.21805, 0.45113, 0.24060, -0.04511},
		Row{-0.80827, -1.45677, -0.44361, 0.52068},
		Row{-0.07895, -0.22368, -0.05263, 0.19737},
		Row{-0.52256, -0.81391, -0.30075, 0.30639},
	}
	for i := range B {
		assert.InDeltaSlice(t, expected[i], B[i], 0.0001)
	}
}

/*
	Scenario: Calculating the inverse of another matrix
	Given the following 4x4 matrix A:
		| 8  | -5 | 9  | 2  |
		| 7  | 5  | 6  | 1  |
		| -6 | 0  | 9  | 6  |
		| -3 | 0  | -9 | -4 |
	Then inverse(A) is the following 4x4 matrix:
		| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
		| -0.07692 | 0.12308  | 0.02564  | 0.03077  |
		| 0.35897  | 0.35897  | 0.43590  | 0.92308  |
		| -0.69231 | -0.69231 | -0.76923 | -1.92308 |
*/
func TestInverse2(t *testing.T) {
	A := Matrix{
		Row{8, -5, 9, 2},
		Row{7, 5, 6, 1},
		Row{-6, 0, 9, 6},
		Row{-3, 0, -9, -4},
	}
	B, err := A.Inverse()
	require.Nil(t, err)
	expected := Matrix{
		Row{-0.15385, -0.15385, -0.28205, -0.53846},
		Row{-0.07692, 0.12308, 0.02564, 0.03077},
		Row{0.35897, 0.35897, 0.43590, 0.92308},
		Row{-0.69231, -0.69231, -0.76923, -1.92308},
	}
	for i := range B {
		assert.InDeltaSlice(t, expected[i], B[i], 0.0001)
	}
}

/*
	Scenario: Calculating the inverse of a third matrix
	Given the following 4x4 matrix A:
		| 9  | 3  | 0  | 9  |
		| -5 | -2 | -6 | -3 |
		| -4 | 9  | 6  | 4  |
		| -7 | 6  | 6  | 2  |
	Then inverse(A) is the following 4x4 matrix:
		| -0.04074 | -0.07778 | 0.14444  | -0.22222 |
		| -0.07778 | 0.03333  | 0.36667  | -0.33333 |
		| -0.02901 | -0.14630 | -0.10926 | 0.12963  |
		| 0.17778  | 0.06667  | -0.26667 | 0.33333  |
*/
func TestInverse3(t *testing.T) {
	A := Matrix{
		Row{9, 3, 0, 9},
		Row{-5, -2, -6, -3},
		Row{-4, 9, 6, 4},
		Row{-7, 6, 6, 2},
	}
	B, err := A.Inverse()
	require.Nil(t, err)
	expected := Matrix{
		Row{-0.04074, -0.07778, 0.14444, -0.22222},
		Row{-0.07778, 0.03333, 0.36667, -0.33333},
		Row{-0.02901, -0.14630, -0.10926, 0.12963},
		Row{0.17778, 0.06667, -0.26667, 0.33333},
	}
	for i := range B {
		assert.InDeltaSlice(t, expected[i], B[i], 0.0001)
	}
}

/*
	Scenario: Multiplying a product by its inverse
	Given the following 4x4 matrix A:
		| 3  | -9 | 7  | 3  |
		| 3  | -8 | 2  | -9 |
		| -4 | 4  | 4  | 1  |
		| -6 | 5  | -1 | 1  |
	And the following 4x4 matrix B:
		| 8 | 2  | 2 | 2 |
		| 3 | -1 | 7 | 0 |
		| 7 | 0  | 5 | 4 |
		| 6 | -2 | 0 | 5 |
	And C ← A * B
	Then C * inverse(B) = A
*/
func TestProductMultiplyItsInverse(t *testing.T) {
	A := Matrix{
		Row{3, -9, 7, 3},
		Row{3, -8, 2, -9},
		Row{-4, 4, 4, 1},
		Row{-6, 5, -1, 1},
	}
	B := Matrix{
		Row{8, 2, 2, 2},
		Row{3, -1, 7, 0},
		Row{7, 0, 5, 4},
		Row{6, -2, 0, 5},
	}
	C, err := A.MulM(B)
	require.Nil(t, err)

	Bi, err := B.Inverse()
	require.Nil(t, err)

	CBi, err := C.MulM(Bi)
	require.Nil(t, err)

	for i := range A {
		assert.InDeltaSlice(t, A[i], CBi[i], 0.0001)
	}
}
