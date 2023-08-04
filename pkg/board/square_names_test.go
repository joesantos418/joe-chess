package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGetCoordinateByName_Valid(t *testing.T) {
// 	lin, col = GetCoordinatesBySquareName(A1)

// 	assert.Equal(t, 0, lin)
// 	assert.Equal(t, 0, col)

// 	lin, col = GetCoordinatesBySquareName(E4)

// 	assert.Equal(t, 3, lin)
// 	assert.Equal(t, 4, col)
// }

func TestGetCoordinateByName(t *testing.T) {
	var lin, col int

	lin, col = GetCoordinatesBySquareName(A1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(A8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 0, col)

	lin, col = GetCoordinatesBySquareName(B1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(B8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 1, col)

	lin, col = GetCoordinatesBySquareName(C1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(C8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 2, col)

	lin, col = GetCoordinatesBySquareName(D1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(D8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 3, col)

	lin, col = GetCoordinatesBySquareName(E1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(E8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 4, col)

	lin, col = GetCoordinatesBySquareName(F1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(F8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 5, col)

	lin, col = GetCoordinatesBySquareName(G1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(G8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 6, col)

	lin, col = GetCoordinatesBySquareName(H1)
	assert.Equal(t, 0, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H2)
	assert.Equal(t, 1, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H3)
	assert.Equal(t, 2, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H4)
	assert.Equal(t, 3, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H5)
	assert.Equal(t, 4, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H6)
	assert.Equal(t, 5, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H7)
	assert.Equal(t, 6, lin)
	assert.Equal(t, 7, col)

	lin, col = GetCoordinatesBySquareName(H8)
	assert.Equal(t, 7, lin)
	assert.Equal(t, 7, col)
}
