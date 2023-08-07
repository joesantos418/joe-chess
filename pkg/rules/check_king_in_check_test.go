package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestGetColor(t *testing.T) {
	color := getColor(board.WHITE_KING)
	assert.Equal(t, WHITES, color)

	color = getColor(board.BLACK_KING)
	assert.Equal(t, BLACKS, color)

	color = getColor(board.WHITE_PAWN_A)
	assert.Equal(t, WHITES, color)

	color = getColor(board.BLACK_PAWN_A)
	assert.Equal(t, BLACKS, color)
}

func TestGetMaxSteps(t *testing.T) {
	steps := getMaxSteps(0, 0, NE)
	assert.Equal(t, 7, steps)

	steps = getMaxSteps(6, 3, NE)
	assert.Equal(t, 1, steps)

	steps = getMaxSteps(2, 6, NE)
	assert.Equal(t, 1, steps)

	steps = getMaxSteps(0, 0, NW)
	assert.Equal(t, 0, steps)

	steps = getMaxSteps(6, 3, NW)
	assert.Equal(t, 1, steps)

	steps = getMaxSteps(2, 6, NW)
	assert.Equal(t, 5, steps)

	steps = getMaxSteps(0, 0, SE)
	assert.Equal(t, 0, steps)

	steps = getMaxSteps(6, 3, SE)
	assert.Equal(t, 4, steps)

	steps = getMaxSteps(2, 6, SE)
	assert.Equal(t, 1, steps)

	steps = getMaxSteps(0, 0, SW)
	assert.Equal(t, 0, steps)

	steps = getMaxSteps(6, 3, SW)
	assert.Equal(t, 3, steps)

	steps = getMaxSteps(2, 6, SW)
	assert.Equal(t, 2, steps)
}

func TestIsInRangeForDiagonalAttack(t *testing.T) {
	var is bool

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_A, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_A, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_B, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_B, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_C, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_C, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_D, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_D, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_E, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_E, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_F, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_F, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_G, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_G, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_H, 1, 0, 2, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_PAWN_H, 1, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_A, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_A, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_B, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_B, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_C, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_C, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_D, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_D, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_E, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_E, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_F, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_F, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_G, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_G, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_H, 6, 0, 5, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_PAWN_H, 6, 0, 3, 2)
	assert.False(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_BISHOP_QUEEN, 0, 0, 0, 0)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_BISHOP_KING, 0, 0, 0, 0)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_BISHOP_QUEEN, 0, 0, 0, 0)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_BISHOP_KING, 0, 0, 0, 0)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_QUEEN, 0, 0, 0, 0)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_QUEEN, 0, 0, 0, 0)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_KING, 0, 0, 1, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.BLACK_KING, 0, 0, 1, 1)
	assert.True(t, is)

	is = isInRangeForDiagonalAttack(board.WHITE_KNIGHT_KING, 0, 0, 1, 1)
	assert.False(t, is)
}

func TestIsAttackDiagonal(t *testing.T) {
	is := isAttackDiagonal(board.BLACK_BISHOP_KING)
	assert.True(t, is)

	is = isAttackDiagonal(board.BLACK_KNIGHT_KING)
	assert.False(t, is)
}

func TestIsDiagonalDanger(t *testing.T) {
	is := isDiagonalDanger(board.WHITE_PAWN_A, board.BLACK_PAWN_B, 6, 1, 5, 0)
	assert.True(t, is)

	is = isDiagonalDanger(board.WHITE_PAWN_A, board.BLACK_PAWN_B, 6, 1, 4, 0)
	assert.False(t, is)
}

// Is there a threat in an empty board? No.
func TestIsThreatenedNE_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	D4Lin := 3
	D4Col := 3

	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)
	assert.False(t, is)
}

// Is there a threat when there are other pieces in the board, but none
// in the NE diagonal? No.
func TestIsThreatenedNE_OtherPiecesNoneInDiagonal(t *testing.T) {
	b := &board.Board{}

	// filling up the NE quadrant except for the D4-H8 diagonal
	// nothing in D4
	b.SetPiece(board.BLACK_BISHOP_KING, board.E4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.G4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.H4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.D5)
	// nothing in E5
	b.SetPiece(board.BLACK_BISHOP_KING, board.F5)
	b.SetPiece(board.BLACK_BISHOP_KING, board.G5)
	b.SetPiece(board.BLACK_BISHOP_KING, board.H5)
	b.SetPiece(board.BLACK_BISHOP_KING, board.D6)
	b.SetPiece(board.BLACK_BISHOP_KING, board.E6)
	// nothing in F6
	b.SetPiece(board.BLACK_BISHOP_KING, board.G6)
	b.SetPiece(board.BLACK_BISHOP_KING, board.H6)
	b.SetPiece(board.BLACK_BISHOP_KING, board.D7)
	b.SetPiece(board.BLACK_BISHOP_KING, board.E7)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F7)
	// nothing in g7
	b.SetPiece(board.BLACK_BISHOP_KING, board.H7)
	b.SetPiece(board.BLACK_BISHOP_KING, board.D8)
	b.SetPiece(board.BLACK_BISHOP_KING, board.E8)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F8)
	b.SetPiece(board.BLACK_BISHOP_KING, board.G8)
	// nothing in H8

	// adding pieces to other diagonals
	b.SetPiece(board.BLACK_BISHOP_KING, board.C5)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F2)
	b.SetPiece(board.BLACK_BISHOP_KING, board.A1)

	// adding other random pieces for good measure
	b.SetPiece(board.BLACK_BISHOP_KING, board.A6)
	b.SetPiece(board.BLACK_BISHOP_KING, board.B3)
	b.SetPiece(board.BLACK_BISHOP_KING, board.H1)
	b.SetPiece(board.BLACK_BISHOP_KING, board.B8)
	b.SetPiece(board.BLACK_BISHOP_KING, board.C3)

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	// exercise: why not use the 3 directly?
	// hint: mind beacons
	D4Lin := 3
	D4Col := 3

	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.False(t, is)
}

// Is there a threat when there's other piece in the NE diagonal close the
// defending piece? Yes.
func TestIsThreatenedNE_OtherPieceInDiagonalClose(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.E5)

	D4Lin := 3
	D4Col := 3
	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.True(t, is)
}

// Is there a threat when there's other piece in the NE diagonal far from the
// defending piece? Yes.
func TestIsThreatenedNE_OtherPieceInDiagonalFar(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F6)

	D4Lin := 3
	D4Col := 3
	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.True(t, is)
}

// Is there a threat when there's other piece in the NE diagonal as far possible
// from the defending piece? Yes.
func TestIsThreatenedNE_OtherPieceInDiagonalFarthest(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.H8)

	D4Lin := 3
	D4Col := 3
	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.True(t, is)
}

// Is there a threat when there's other white piece in the NE diagonal close the
// defending piece? No.
func TestIsThreatenedNE_OtherWhitePieceInDiagonalClose(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.WHITE_BISHOP_KING, board.E5)

	D4Lin := 3
	D4Col := 3
	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.False(t, is)
}

// Is there a threat when there's other white piece in the NE diagonal far from
// the defending piece? No.
func TestIsThreatenedNE_OtherWhitePieceInDiagonalFar(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.WHITE_BISHOP_KING, board.F6)

	D4Lin := 3
	D4Col := 3
	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.False(t, is)
}

// Is there a threat when there's other white piece in the NE diagonal as far
// possible from the defending piece? No.
func TestIsThreatenedNE_OtherWhitePieceInDiagonalFarthest(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.WHITE_BISHOP_KING, board.H8)

	D4Lin := 3
	D4Col := 3
	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, D4Lin, D4Col, b)

	assert.False(t, is)
}

// exercise: this is the legacy test
// why is it not the best?
// hint: each test case must be one test
// func TestIsThreatenedNE(t *testing.T) {
// 	b := &board.Board{}

// 	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)

// 	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.False(t, is)

// 	b.SetPiece(board.BLACK_BISHOP_KING, board.G6)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.False(t, is)

// 	b.SetPiece(board.BLACK_BISHOP_KING, board.F7)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.False(t, is)

// 	b.SetPiece(board.BLACK_BISHOP_KING, board.F6)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.True(t, is)

// 	b.SetPiece(board.NO_PIECE, board.F6)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.False(t, is)

// 	b.SetPiece(board.BLACK_BISHOP_KING, board.H8)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.True(t, is)

// 	b.SetPiece(board.NO_PIECE, board.F6)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.False(t, is)

// 	b.SetPiece(board.BLACK_BISHOP_KING, board.E5)
// 	is = isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
// 	assert.True(t, is)
// }

func TestIsThreatenedNW(t *testing.T) {
	b := &board.Board{}

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)

	is := isThreatenedNW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.C6)
	is = isThreatenedNW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.B5)
	is = isThreatenedNW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.B6)
	is = isThreatenedNW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.True(t, is)
}

func TestIsThreatenedSE(t *testing.T) {
	b := &board.Board{}

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)

	is := isThreatenedSE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.E2)
	is = isThreatenedSE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.F3)
	is = isThreatenedSE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.G1)
	is = isThreatenedSE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.True(t, is)

	b.SetPiece(board.NO_PIECE, board.G1)
	is = isThreatenedSE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.E3)
	is = isThreatenedSE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.True(t, is)
}

func TestIsThreatenedSW(t *testing.T) {
	b := &board.Board{}

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)

	is := isThreatenedSW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.C2)
	is = isThreatenedSW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.B3)
	is = isThreatenedSW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.A1)
	is = isThreatenedSW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.True(t, is)

	b.SetPiece(board.NO_PIECE, board.A1)
	is = isThreatenedSW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.C3)
	is = isThreatenedSW(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.True(t, is)
}

func TestIsInRangeForHorizontalAttack_Pawn(t *testing.T) {
	p := board.WHITE_PAWN_A
	from := 0
	to := 2

	is := isInRangeForHorizontalAttack(p, from, to)
	assert.False(t, is)
}

func TestIsInRangeForHorizontalAttack_Rook(t *testing.T) {
	p := board.WHITE_ROOK_KING
	from := 0
	to := 5

	is := isInRangeForHorizontalAttack(p, from, to)
	assert.True(t, is)
}

func TestIsInRangeForHorizontalAttack_KingClose(t *testing.T) {
	p := board.WHITE_KING
	from := 0
	to := 1

	is := isInRangeForHorizontalAttack(p, from, to)
	assert.True(t, is)
}

func TestIsInRangeForHorizontalAttack_KingFar(t *testing.T) {
	p := board.WHITE_KING
	from := 0
	to := 3

	is := isInRangeForHorizontalAttack(p, from, to)
	assert.False(t, is)
}

func TestIsAttackHorizontal_King(t *testing.T) {
	is := isAttackHorizontal(board.BLACK_KING)
	assert.True(t, is)
}

func TestIsAttackHorizontal_Bishop(t *testing.T) {
	is := isAttackHorizontal(board.BLACK_BISHOP_KING)
	assert.False(t, is)
}

func TestIsHorizontalDanger_SameColor(t *testing.T) {
	is := isHorizontalDanger(board.WHITE_PAWN_A, board.WHITE_QUEEN, 0, 2)
	assert.False(t, is)
}

func TestIsHorizontalDanger_NotHorizontalAttacker(t *testing.T) {
	is := isHorizontalDanger(board.WHITE_PAWN_A, board.BLACK_BISHOP_KING, 0, 2)
	assert.False(t, is)
}

func TestIsHorizontalDanger_OutOfRange(t *testing.T) {
	is := isHorizontalDanger(board.WHITE_PAWN_A, board.BLACK_KING, 0, 2)
	assert.False(t, is)
}

func TestIsHorizontalDanger_InDanger(t *testing.T) {
	is := isHorizontalDanger(board.WHITE_PAWN_A, board.BLACK_KING, 0, 1)
	assert.False(t, is)
}
