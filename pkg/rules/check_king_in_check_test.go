package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestGetWhiteKingCoordinates(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_KING, board.A1)

	lin, col := getWhiteKingCoordinates(b)

	assert.Equal(t, 0, lin)
	assert.Equal(t, 0, col)
}

func TestGetWhiteKingCoordinates_OtherPieces(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_KING, board.A1)
	b.SetPiece(board.WHITE_PAWN_A, board.B1)

	lin, col := getWhiteKingCoordinates(b)

	assert.Equal(t, 0, lin)
	assert.Equal(t, 0, col)
}

func TestGetBlackKingCoordinates(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_KING, board.A1)

	lin, col := getBlackKingCoordinates(b)

	assert.Equal(t, 0, lin)
	assert.Equal(t, 0, col)
}

func TestIsSquareInCheck_Vertical(t *testing.T) {
	lin := 0
	col := 0
	b := &board.Board{}

	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_ROOK_KING, board.A3)

	c := WHITES
	is := isSquareInCheck(lin, col, b, c)

	assert.True(t, is)
}

func TestIsSquareInCheck_Horizontal(t *testing.T) {
	lin := 0
	col := 0
	b := &board.Board{}

	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_ROOK_KING, board.G1)

	c := WHITES
	is := isSquareInCheck(lin, col, b, c)

	assert.True(t, is)
}

func TestIsSquareInCheck_Diagonal(t *testing.T) {
	lin := 0
	col := 0
	b := &board.Board{}

	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_BISHOP_KING, board.D4)

	c := WHITES
	is := isSquareInCheck(lin, col, b, c)

	assert.True(t, is)
}

func TestIsSquareInCheck_L(t *testing.T) {
	lin := 0
	col := 0
	b := &board.Board{}

	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_KNIGHT_KING, board.C2)

	c := WHITES
	is := isSquareInCheck(lin, col, b, c)

	assert.True(t, is)
}

func TestIsSquareInCheck_No(t *testing.T) {
	lin := 0
	col := 0
	b := &board.Board{}

	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_KNIGHT_KING, board.G1)

	c := WHITES
	is := isSquareInCheck(lin, col, b, c)

	assert.False(t, is)
}

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
	is := isDiagonalDanger(board.BLACK_PAWN_B, 6, 1, 5, 0, WHITES)
	assert.True(t, is)

	is = isDiagonalDanger(board.BLACK_PAWN_B, 6, 1, 4, 0, WHITES)
	assert.False(t, is)
}

// Is there a threat in an empty board? No.
func TestIsThreatenedNE_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	D4Lin := 3
	D4Col := 3

	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)
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

	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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
	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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
	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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
	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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
	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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
	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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
	is := isThreatenedNE(D4Lin, D4Col, b, WHITES)

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

	is := isThreatenedNW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.C6)
	is = isThreatenedNW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.B5)
	is = isThreatenedNW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.B6)
	is = isThreatenedNW(3, 3, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedSE(t *testing.T) {
	b := &board.Board{}

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)

	is := isThreatenedSE(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.E2)
	is = isThreatenedSE(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.F3)
	is = isThreatenedSE(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.G1)
	is = isThreatenedSE(3, 3, b, WHITES)
	assert.True(t, is)

	b.SetPiece(board.NO_PIECE, board.G1)
	is = isThreatenedSE(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.E3)
	is = isThreatenedSE(3, 3, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedSW(t *testing.T) {
	b := &board.Board{}

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)

	is := isThreatenedSW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.C2)
	is = isThreatenedSW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.B3)
	is = isThreatenedSW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.A1)
	is = isThreatenedSW(3, 3, b, WHITES)
	assert.True(t, is)

	b.SetPiece(board.NO_PIECE, board.A1)
	is = isThreatenedSW(3, 3, b, WHITES)
	assert.False(t, is)

	b.SetPiece(board.BLACK_BISHOP_KING, board.C3)
	is = isThreatenedSW(3, 3, b, WHITES)
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
	is := isHorizontalDanger(board.WHITE_QUEEN, 0, 2, WHITES)
	assert.False(t, is)
}

func TestIsHorizontalDanger_NotHorizontalAttacker(t *testing.T) {
	is := isHorizontalDanger(board.BLACK_BISHOP_KING, 0, 2, WHITES)
	assert.False(t, is)
}

func TestIsHorizontalDanger_OutOfRange(t *testing.T) {
	is := isHorizontalDanger(board.BLACK_KING, 0, 2, WHITES)
	assert.False(t, is)
}

func TestIsHorizontalDanger_InDanger(t *testing.T) {
	is := isHorizontalDanger(board.BLACK_KING, 0, 1, BLACKS)
	assert.False(t, is)
}

func TestIsThreatenedLeft_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)

	is := isThreatenedLeft(0, 4, b, BLACKS)

	assert.False(t, is)
}

func TestIsThreatenedLeft_Close(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)
	b.SetPiece(board.WHITE_ROOK_KING, board.D1)

	is := isThreatenedLeft(0, 4, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedLeft_Far(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)
	b.SetPiece(board.WHITE_ROOK_KING, board.B1)

	is := isThreatenedLeft(0, 4, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedLeft_Farther(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.H1)
	b.SetPiece(board.WHITE_ROOK_KING, board.A1)

	is := isThreatenedLeft(0, 4, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedLeft_NoLeft(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A1)
	b.SetPiece(board.WHITE_ROOK_KING, board.H1)

	is := isThreatenedLeft(0, 0, b, BLACKS)

	assert.False(t, is)
}

func TestIsThreatenedRight_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)

	is := isThreatenedRight(0, 4, b, BLACKS)

	assert.False(t, is)
}

func TestIsThreatenedRight_Close(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)
	b.SetPiece(board.WHITE_ROOK_KING, board.F1)

	is := isThreatenedRight(0, 4, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedRight_Far(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)
	b.SetPiece(board.WHITE_ROOK_KING, board.G1)

	is := isThreatenedRight(0, 4, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedRight_Farther(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.E1)
	b.SetPiece(board.WHITE_ROOK_KING, board.H1)

	is := isThreatenedRight(0, 4, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedRight_NoRight(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.H1)
	b.SetPiece(board.WHITE_ROOK_KING, board.A1)

	is := isThreatenedRight(0, 7, b, BLACKS)

	assert.False(t, is)
}

func TestIsInRangeForVerticalAttack_Pawn(t *testing.T) {
	p := board.WHITE_PAWN_A
	from := 0
	to := 2

	is := isInRangeForVerticalAttack(p, from, to)
	assert.False(t, is)
}

func TestIsInRangeForVerticalAttack_Rook(t *testing.T) {
	p := board.WHITE_ROOK_KING
	from := 0
	to := 5

	is := isInRangeForVerticalAttack(p, from, to)
	assert.True(t, is)
}

func TestIsInRangeForVerticalAttack_KingClose(t *testing.T) {
	p := board.WHITE_KING
	from := 0
	to := 1

	is := isInRangeForVerticalAttack(p, from, to)
	assert.True(t, is)
}

func TestIsInRangeForVerticalAttack_KingFar(t *testing.T) {
	p := board.WHITE_KING
	from := 0
	to := 3

	is := isInRangeForVerticalAttack(p, from, to)
	assert.False(t, is)
}

func TestIsAttackVertical_King(t *testing.T) {
	is := isAttackVertical(board.BLACK_KING)
	assert.True(t, is)
}

func TestIsAttackVertical_Bishop(t *testing.T) {
	is := isAttackVertical(board.BLACK_BISHOP_KING)
	assert.False(t, is)
}

func TestIsVerticalDanger_SameColor(t *testing.T) {
	is := isVerticalDanger(board.WHITE_QUEEN, 0, 2, WHITES)
	assert.False(t, is)
}

func TestIsVerticalDanger_NotVerticalAttacker(t *testing.T) {
	is := isVerticalDanger(board.BLACK_BISHOP_KING, 0, 2, BLACKS)
	assert.False(t, is)
}

func TestIsVerticalDanger_OutOfRange(t *testing.T) {
	is := isVerticalDanger(board.BLACK_KING, 0, 2, BLACKS)
	assert.False(t, is)
}

func TestIsVerticalDanger_InDanger(t *testing.T) {
	is := isVerticalDanger(board.BLACK_KING, 0, 1, BLACKS)
	assert.False(t, is)
}

func TestIsThreatenedDown_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A7)

	is := isThreatenedDown(6, 0, b, BLACKS)

	assert.False(t, is)
}

func TestIsThreatenedDown_Close(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A7)
	b.SetPiece(board.WHITE_ROOK_KING, board.A6)

	is := isThreatenedDown(6, 0, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedDown_Far(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A7)
	b.SetPiece(board.WHITE_ROOK_KING, board.A3)

	is := isThreatenedDown(6, 0, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedDown_Farther(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A7)
	b.SetPiece(board.WHITE_ROOK_KING, board.A1)

	is := isThreatenedDown(6, 0, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedDown_NoDown(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A1)
	b.SetPiece(board.WHITE_ROOK_KING, board.A3)

	is := isThreatenedDown(0, 0, b, BLACKS)

	assert.False(t, is)
}

func TestIsThreatenedUp_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A7)

	is := isThreatenedUp(6, 0, b, WHITES)

	assert.False(t, is)
}

func TestIsThreatenedUp_Close(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A2)
	b.SetPiece(board.WHITE_ROOK_KING, board.A3)

	is := isThreatenedUp(1, 0, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedUp_Far(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A2)
	b.SetPiece(board.WHITE_ROOK_KING, board.A6)

	is := isThreatenedUp(1, 0, b, BLACKS)

	assert.True(t, is)
}

func TestIsThreatenedUp_Farther(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A1)
	b.SetPiece(board.WHITE_ROOK_KING, board.A7)

	is := isThreatenedUp(0, 0, b, BLACKS)

	assert.True(t, is)
}

// coverage helped show that this test was giving the right result but not for
// the right reason
func TestIsThreatenedUp_NoUp(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.BLACK_PAWN_A, board.A8)
	b.SetPiece(board.WHITE_ROOK_KING, board.A3)

	is := isThreatenedUp(7, 0, b, BLACKS)

	assert.False(t, is)
}

func TestIsLAttacker_Knight(t *testing.T) {
	is := isLAttacker(board.WHITE_KNIGHT_QUEEN)
	assert.True(t, is)
}

func TestIsLAttacker_Bishop(t *testing.T) {
	is := isLAttacker(board.WHITE_BISHOP_QUEEN)
	assert.False(t, is)
}

func TestIsDangerL_OutOfRange(t *testing.T) {
	b := &board.Board{}
	is := isDangerL(-1, 0, b, WHITES)
	assert.False(t, is)
}

func TestIsDangerL_EmptyBoard(t *testing.T) {
	b := &board.Board{}
	is := isDangerL(0, 0, b, WHITES)
	assert.False(t, is)
}

func TestIsDangerL_SameColor(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A2)
	b.SetPiece(board.WHITE_KNIGHT_KING, board.C3)

	is := isDangerL(2, 2, b, WHITES)
	assert.False(t, is)
}

func TestIsDangerL_Bishop(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A2)
	b.SetPiece(board.WHITE_BISHOP_KING, board.C3)

	is := isDangerL(2, 2, b, WHITES)
	assert.False(t, is)
}

func TestIsDangerL_Danger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A2)
	b.SetPiece(board.BLACK_KNIGHT_KING, board.C3)

	is := isDangerL(2, 2, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedL_EmptyBoard(t *testing.T) {
	b := &board.Board{}

	is := isThreatenedL(2, 3, b, WHITES)
	assert.False(t, is)
}

func TestIsThreatenedL_TwoColumnsFar(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A2)
	b.SetPiece(board.BLACK_KNIGHT_KING, board.C3)

	is := isThreatenedL(1, 0, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedL_TwoLinesFar(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A2)
	b.SetPiece(board.BLACK_KNIGHT_KING, board.B4)

	is := isThreatenedL(1, 0, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedDiagonal_Danger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_BISHOP_KING, board.H8)

	is := isThreatenedDiagonal(0, 0, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedDiagonal_NoDanger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F8)

	is := isThreatenedDiagonal(0, 0, b, WHITES)
	assert.False(t, is)
}

func TestIsThreatenedHorizontal_Danger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_QUEEN, board.H1)

	is := isThreatenedHorizontal(0, 0, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedHorizontal_NoDanger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_QUEEN, board.H8)

	is := isThreatenedHorizontal(0, 0, b, WHITES)
	assert.False(t, is)
}

func TestIsThreatenedVertical_Danger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_QUEEN, board.A8)

	is := isThreatenedVertical(0, 0, b, WHITES)
	assert.True(t, is)
}

func TestIsThreatenedVertical_NoDanger(t *testing.T) {
	b := &board.Board{}
	b.SetPiece(board.WHITE_PAWN_A, board.A1)
	b.SetPiece(board.BLACK_QUEEN, board.H8)

	is := isThreatenedVertical(0, 0, b, WHITES)
	assert.False(t, is)
}

func TestCheckKingInCheck_Check(t *testing.T) {
	g := NewGame()
	b := &board.Board{}

	b.SetPiece(board.WHITE_KING, board.E4)
	b.SetPiece(board.BLACK_QUEEN, board.H1)

	is := checkKingInCheck(g, b)
	assert.True(t, is)
}

func TestCheckKingInCheck_NoCheck(t *testing.T) {
	g := NewGame()
	b := &board.Board{}

	b.SetPiece(board.WHITE_KING, board.E4)
	b.SetPiece(board.BLACK_QUEEN, board.H2)

	is := checkKingInCheck(g, b)
	assert.False(t, is)
}

func TestSimulateMove(t *testing.T) {
	before := NewGame().Board

	after, err := simulateMove(board.WHITE_PAWN_A, board.A2, board.A4, before)

	assert.Nil(t, err)
	assert.Equal(t, board.WHITE_PAWN_A, after.GetPiece(board.A4))
}

func TestSimulateMove_Err(t *testing.T) {
	before := NewGame().Board

	_, err := simulateMove(board.WHITE_PAWN_A, board.A3, board.A5, before)

	assert.NotNil(t, err)
	assert.Equal(t, "cannot find piece to move [piece=WHITE_PAWN_A,from=A3,to=A5]", err.Error())
}

func TestGetClonedBoard(t *testing.T) {
	g := NewGame()
	b := getClonedBoard(g)

	assert.Equal(t, g.Board, b)
}

func TestIsKingInCheck_NoCheck(t *testing.T) {
	g := NewGame()

	is, err := isKingInCheck(board.WHITE_PAWN_A, board.A2, board.A4, g)

	assert.Nil(t, err)
	assert.False(t, is)
}

func TestIsKingInCheck_Error(t *testing.T) {
	g := NewGame()

	is, err := isKingInCheck(board.WHITE_PAWN_A, board.A3, board.A4, g)

	assert.NotNil(t, err)
	assert.False(t, is)
}

func TestIsKingInCheck_Check(t *testing.T) {
	g := NewGame()

	g.Board.SetPiece(board.BLACK_QUEEN, board.E2)

	is, err := isKingInCheck(board.WHITE_PAWN_A, board.A2, board.A4, g)

	assert.Nil(t, err)
	assert.True(t, is)
}
