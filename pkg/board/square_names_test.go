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

func TestStringSquareNames(t *testing.T) {
	n := A1
	assert.Equal(t, "A1", n.String())
	n = A2
	assert.Equal(t, "A2", n.String())
	n = A3
	assert.Equal(t, "A3", n.String())
	n = A4
	assert.Equal(t, "A4", n.String())
	n = A5
	assert.Equal(t, "A5", n.String())
	n = A6
	assert.Equal(t, "A6", n.String())
	n = A7
	assert.Equal(t, "A7", n.String())
	n = A8
	assert.Equal(t, "A8", n.String())
	n = B1
	assert.Equal(t, "B1", n.String())
	n = B2
	assert.Equal(t, "B2", n.String())
	n = B3
	assert.Equal(t, "B3", n.String())
	n = B4
	assert.Equal(t, "B4", n.String())
	n = B5
	assert.Equal(t, "B5", n.String())
	n = B6
	assert.Equal(t, "B6", n.String())
	n = B7
	assert.Equal(t, "B7", n.String())
	n = B8
	assert.Equal(t, "B8", n.String())
	n = C1
	assert.Equal(t, "C1", n.String())
	n = C2
	assert.Equal(t, "C2", n.String())
	n = C3
	assert.Equal(t, "C3", n.String())
	n = C4
	assert.Equal(t, "C4", n.String())
	n = C5
	assert.Equal(t, "C5", n.String())
	n = C6
	assert.Equal(t, "C6", n.String())
	n = C7
	assert.Equal(t, "C7", n.String())
	n = C8
	assert.Equal(t, "C8", n.String())
	n = D1
	assert.Equal(t, "D1", n.String())
	n = D2
	assert.Equal(t, "D2", n.String())
	n = D3
	assert.Equal(t, "D3", n.String())
	n = D4
	assert.Equal(t, "D4", n.String())
	n = D5
	assert.Equal(t, "D5", n.String())
	n = D6
	assert.Equal(t, "D6", n.String())
	n = D7
	assert.Equal(t, "D7", n.String())
	n = D8
	assert.Equal(t, "D8", n.String())
	n = E1
	assert.Equal(t, "E1", n.String())
	n = E2
	assert.Equal(t, "E2", n.String())
	n = E3
	assert.Equal(t, "E3", n.String())
	n = E4
	assert.Equal(t, "E4", n.String())
	n = E5
	assert.Equal(t, "E5", n.String())
	n = E6
	assert.Equal(t, "E6", n.String())
	n = E7
	assert.Equal(t, "E7", n.String())
	n = E8
	assert.Equal(t, "E8", n.String())
	n = F1
	assert.Equal(t, "F1", n.String())
	n = F2
	assert.Equal(t, "F2", n.String())
	n = F3
	assert.Equal(t, "F3", n.String())
	n = F4
	assert.Equal(t, "F4", n.String())
	n = F5
	assert.Equal(t, "F5", n.String())
	n = F6
	assert.Equal(t, "F6", n.String())
	n = F7
	assert.Equal(t, "F7", n.String())
	n = F8
	assert.Equal(t, "F8", n.String())
	n = G1
	assert.Equal(t, "G1", n.String())
	n = G2
	assert.Equal(t, "G2", n.String())
	n = G3
	assert.Equal(t, "G3", n.String())
	n = G4
	assert.Equal(t, "G4", n.String())
	n = G5
	assert.Equal(t, "G5", n.String())
	n = G6
	assert.Equal(t, "G6", n.String())
	n = G7
	assert.Equal(t, "G7", n.String())
	n = G8
	assert.Equal(t, "G8", n.String())
	n = H1
	assert.Equal(t, "H1", n.String())
	n = H2
	assert.Equal(t, "H2", n.String())
	n = H3
	assert.Equal(t, "H3", n.String())
	n = H4
	assert.Equal(t, "H4", n.String())
	n = H5
	assert.Equal(t, "H5", n.String())
	n = H6
	assert.Equal(t, "H6", n.String())
	n = H7
	assert.Equal(t, "H7", n.String())
	n = H8
	assert.Equal(t, "H8", n.String())
}
