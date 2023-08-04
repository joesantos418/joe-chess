package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsCapture(t *testing.T) {
	g := NewGame()

	g.Board.Move(board.WHITE_PAWN_E, board.E2, board.E6)
	is := isCapture(board.WHITE_PAWN_E, board.D7, g)
	assert.True(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.E7, g)
	assert.True(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.F7, g)
	assert.True(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.F6, g)
	assert.False(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.F5, g)
	assert.False(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.F4, g)
	assert.False(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.E4, g)
	assert.False(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.D4, g)
	assert.False(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.D5, g)
	assert.False(t, is)

	is = isCapture(board.WHITE_PAWN_E, board.D6, g)
	assert.False(t, is)
}

func TestIsDiagonalUpOne(t *testing.T) {
	is := isDiagonalUpOne(board.B2, board.A3)
	assert.True(t, is)

	is = isDiagonalUpOne(board.B2, board.C3)
	assert.True(t, is)

	is = isDiagonalUpOne(board.B2, board.B3)
	assert.False(t, is)
}

func TestIsFirstSquareForWhitePawn(t *testing.T) {
	is := isFirstSquareForWhitePawn(board.A2)
	assert.True(t, is)

	is = isFirstSquareForWhitePawn(board.A1)
	assert.False(t, is)
}

func TestIsVerticalUpOne(t *testing.T) {
	is := isVerticalUpOne(board.A2, board.A3)
	assert.True(t, is)

	is = isVerticalUpOne(board.A1, board.A3)
	assert.False(t, is)
}

func TestIsVerticalUpTwo(t *testing.T) {
	is := isVerticalUpTwo(board.A1, board.A3)
	assert.True(t, is)

	is = isVerticalUpTwo(board.A2, board.A3)
	assert.False(t, is)
}

func TestIsDiagonalDownOne(t *testing.T) {
	is := isDiagonalDownOne(board.B7, board.A6)
	assert.True(t, is)

	is = isDiagonalDownOne(board.B7, board.C6)
	assert.True(t, is)

	is = isDiagonalDownOne(board.B7, board.B6)
	assert.False(t, is)
}

func TestIsFirstSquareForBlackPawn(t *testing.T) {
	is := isFirstSquareForBlackPawn(board.A7)
	assert.True(t, is)

	is = isFirstSquareForBlackPawn(board.A8)
	assert.False(t, is)
}

func TestIsVerticalDownOne(t *testing.T) {
	is := isVerticalDownOne(board.A7, board.A6)
	assert.True(t, is)

	is = isVerticalDownOne(board.A8, board.A6)
	assert.False(t, is)
}

func TestIsVerticalDownTwo(t *testing.T) {
	is := isVerticalDownTwo(board.A8, board.A6)
	assert.True(t, is)

	is = isVerticalDownTwo(board.A7, board.A6)
	assert.False(t, is)
}

func TestIsL(t *testing.T) {
	is := isL(board.D5, board.C7)
	assert.True(t, is)

	is = isL(board.D5, board.B7)
	assert.False(t, is)

	is = isL(board.D5, board.C6)
	assert.False(t, is)

	is = isL(board.D5, board.B6)
	assert.True(t, is)

	is = isL(board.D5, board.E7)
	assert.True(t, is)

	is = isL(board.D5, board.F6)
	assert.True(t, is)

	is = isL(board.D5, board.F4)
	assert.True(t, is)

	is = isL(board.D5, board.E3)
	assert.True(t, is)

	is = isL(board.D5, board.C3)
	assert.True(t, is)

	is = isL(board.D5, board.B4)
	assert.True(t, is)
}

func TestIsVerticalOne(t *testing.T) {
	is := isVerticalOne(board.A2, board.A3)
	assert.True(t, is)

	is = isVerticalOne(board.A1, board.A3)
	assert.False(t, is)
}

func TestIsHorizontalOne(t *testing.T) {
	is := isHorizontalOne(board.A2, board.B2)
	assert.True(t, is)

	is = isHorizontalOne(board.A1, board.A3)
	assert.False(t, is)
}

func TestIsDiagonalOne(t *testing.T) {
	is := isDiagonalOne(board.B2, board.A3)
	assert.True(t, is)

	is = isDiagonalOne(board.A1, board.A3)
	assert.False(t, is)
}

func TestIsVerticalN(t *testing.T) {
	is := isVerticalN(board.A2, board.A3)
	assert.True(t, is)

	is = isVerticalN(board.B1, board.A3)
	assert.False(t, is)

	is = isVerticalN(board.A2, board.A8)
	assert.True(t, is)
}

func TestIsHorizontalN(t *testing.T) {
	is := isHorizontalN(board.A2, board.B2)
	assert.True(t, is)

	is = isHorizontalN(board.A1, board.A3)
	assert.False(t, is)

	is = isHorizontalN(board.A2, board.H2)
	assert.True(t, is)
}

func TestIsDiagonalN(t *testing.T) {
	is := isDiagonalN(board.B2, board.A3)
	assert.True(t, is)

	is = isDiagonalN(board.A1, board.A3)
	assert.False(t, is)

	is = isDiagonalN(board.B2, board.G7)
	assert.True(t, is)
}

func TestIsCorrectForWhitePawn(t *testing.T) {
	g := NewGame()
	is := isCorrectForWhitePawn(board.WHITE_PAWN_A, board.A2, board.A3, g)
	assert.True(t, is)

	is = isCorrectForWhitePawn(board.WHITE_PAWN_A, board.A2, board.A4, g)
	assert.True(t, is)

	is = isCorrectForWhitePawn(board.WHITE_PAWN_A, board.A2, board.A5, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_PAWN_A, board.A2, board.A6)
	is = isCorrectForWhitePawn(board.WHITE_PAWN_A, board.A6, board.B7, g)
	assert.True(t, is)

	g.Board.Move(board.WHITE_PAWN_A, board.A6, board.A3)
	is = isCorrectForWhitePawn(board.WHITE_PAWN_A, board.A3, board.A4, g)
	assert.True(t, is)

	is = isCorrectForWhitePawn(board.WHITE_PAWN_A, board.A3, board.A5, g)
	assert.False(t, is)
}

func TestIsCorrectForBlackPawn(t *testing.T) {
	g := NewGame()
	is := isCorrectForBlackPawn(board.BLACK_PAWN_A, board.A7, board.A6, g)
	assert.True(t, is)

	is = isCorrectForBlackPawn(board.BLACK_PAWN_A, board.A7, board.A5, g)
	assert.True(t, is)

	is = isCorrectForBlackPawn(board.BLACK_PAWN_A, board.A7, board.A4, g)
	assert.False(t, is)

	g.Board.Move(board.BLACK_PAWN_A, board.A7, board.A3)
	is = isCorrectForBlackPawn(board.BLACK_PAWN_A, board.A3, board.B2, g)
	assert.True(t, is)

	g.Board.Move(board.BLACK_PAWN_A, board.A3, board.A6)
	is = isCorrectForBlackPawn(board.BLACK_PAWN_A, board.A6, board.A5, g)
	assert.True(t, is)

	is = isCorrectForBlackPawn(board.BLACK_PAWN_A, board.A6, board.A4, g)
	assert.False(t, is)
}

func TestIsCorrectForRook(t *testing.T) {
	g := NewGame()
	is := isCorrectForRook(board.A1, board.A6, g)
	assert.True(t, is)

	is = isCorrectForRook(board.A1, board.B6, g)
	assert.False(t, is)
}

func TestIsCorrectForKnight(t *testing.T) {
	g := NewGame()
	is := isCorrectForKnight(board.B1, board.A3, g)
	assert.True(t, is)

	is = isCorrectForKnight(board.A1, board.B6, g)
	assert.False(t, is)
}

func TestIsCorrectForBishop(t *testing.T) {
	g := NewGame()
	is := isCorrectForBishop(board.C1, board.A3, g)
	assert.True(t, is)

	is = isCorrectForBishop(board.A1, board.B6, g)
	assert.False(t, is)
}

func TestIsCorrectForQueen(t *testing.T) {
	g := NewGame()
	is := isCorrectForQueen(board.C1, board.A3, g)
	assert.True(t, is)

	is = isCorrectForQueen(board.A1, board.B6, g)
	assert.False(t, is)
}

func TestIsCorrectForKing(t *testing.T) {
	g := NewGame()
	is := isCorrectForKing(board.E1, board.E2, g)
	assert.True(t, is)

	is = isCorrectForKing(board.A1, board.B6, g)
	assert.False(t, is)
}

func TestIsMoveCorrect(t *testing.T) {
	var is bool
	g := NewGame()

	is = isMoveCorrect(board.NO_PIECE, board.A1, board.A3, g)
	assert.False(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_A, board.A2, board.A3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_B, board.B2, board.B3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_C, board.C2, board.C3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_D, board.D2, board.D3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_E, board.E2, board.E3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_F, board.F2, board.F3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_G, board.G2, board.G3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_H, board.H2, board.H3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_ROOK_QUEEN, board.A1, board.A4, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_ROOK_KING, board.H1, board.H4, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_KNIGHT_QUEEN, board.B1, board.A3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_KNIGHT_KING, board.G1, board.H3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_BISHOP_QUEEN, board.C1, board.E3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_BISHOP_KING, board.F1, board.D3, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_QUEEN, board.D1, board.D5, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_KING, board.E1, board.E2, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_A, board.A7, board.A6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_B, board.B7, board.B6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_C, board.C7, board.C6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_D, board.D7, board.D6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_E, board.E7, board.E6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_F, board.F7, board.F6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_G, board.G7, board.G6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_PAWN_H, board.H7, board.H6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_ROOK_QUEEN, board.A8, board.A2, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_ROOK_KING, board.H8, board.H2, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_KNIGHT_QUEEN, board.B8, board.A6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_KNIGHT_KING, board.G8, board.H6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_BISHOP_QUEEN, board.C8, board.A6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_BISHOP_KING, board.F8, board.H6, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_QUEEN, board.D8, board.D1, g)
	assert.True(t, is)

	is = isMoveCorrect(board.BLACK_KING, board.E8, board.E7, g)
	assert.True(t, is)

	is = isMoveCorrect(board.WHITE_PAWN_A, board.A2, board.A2, g)
	assert.True(t, is)
}
