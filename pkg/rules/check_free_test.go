package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	g := NewGame()

	is := isEmpty(board.A1, g)
	assert.False(t, is)

	is = isEmpty(board.A2, g)
	assert.False(t, is)

	is = isEmpty(board.A3, g)
	assert.True(t, is)

	is = isEmpty(board.A4, g)
	assert.True(t, is)

	is = isEmpty(board.A5, g)
	assert.True(t, is)

	is = isEmpty(board.A6, g)
	assert.True(t, is)

	is = isEmpty(board.A7, g)
	assert.False(t, is)

	is = isEmpty(board.A8, g)
	assert.False(t, is)

	is = isEmpty(board.B1, g)
	assert.False(t, is)

	is = isEmpty(board.B2, g)
	assert.False(t, is)

	is = isEmpty(board.B3, g)
	assert.True(t, is)

	is = isEmpty(board.B4, g)
	assert.True(t, is)

	is = isEmpty(board.B5, g)
	assert.True(t, is)

	is = isEmpty(board.B6, g)
	assert.True(t, is)

	is = isEmpty(board.B7, g)
	assert.False(t, is)

	is = isEmpty(board.B8, g)
	assert.False(t, is)

	is = isEmpty(board.C1, g)
	assert.False(t, is)

	is = isEmpty(board.C2, g)
	assert.False(t, is)

	is = isEmpty(board.C3, g)
	assert.True(t, is)

	is = isEmpty(board.C4, g)
	assert.True(t, is)

	is = isEmpty(board.C5, g)
	assert.True(t, is)

	is = isEmpty(board.C6, g)
	assert.True(t, is)

	is = isEmpty(board.C7, g)
	assert.False(t, is)

	is = isEmpty(board.C8, g)
	assert.False(t, is)

	is = isEmpty(board.D1, g)
	assert.False(t, is)

	is = isEmpty(board.D2, g)
	assert.False(t, is)

	is = isEmpty(board.D3, g)
	assert.True(t, is)

	is = isEmpty(board.D4, g)
	assert.True(t, is)

	is = isEmpty(board.D5, g)
	assert.True(t, is)

	is = isEmpty(board.D6, g)
	assert.True(t, is)

	is = isEmpty(board.D7, g)
	assert.False(t, is)

	is = isEmpty(board.D8, g)
	assert.False(t, is)

	is = isEmpty(board.E1, g)
	assert.False(t, is)

	is = isEmpty(board.E2, g)
	assert.False(t, is)

	is = isEmpty(board.E3, g)
	assert.True(t, is)

	is = isEmpty(board.E4, g)
	assert.True(t, is)

	is = isEmpty(board.E5, g)
	assert.True(t, is)

	is = isEmpty(board.E6, g)
	assert.True(t, is)

	is = isEmpty(board.E7, g)
	assert.False(t, is)

	is = isEmpty(board.E8, g)
	assert.False(t, is)

	is = isEmpty(board.F1, g)
	assert.False(t, is)

	is = isEmpty(board.F2, g)
	assert.False(t, is)

	is = isEmpty(board.F3, g)
	assert.True(t, is)

	is = isEmpty(board.F4, g)
	assert.True(t, is)

	is = isEmpty(board.F5, g)
	assert.True(t, is)

	is = isEmpty(board.F6, g)
	assert.True(t, is)

	is = isEmpty(board.F7, g)
	assert.False(t, is)

	is = isEmpty(board.F8, g)
	assert.False(t, is)

	is = isEmpty(board.G1, g)
	assert.False(t, is)

	is = isEmpty(board.G2, g)
	assert.False(t, is)

	is = isEmpty(board.G3, g)
	assert.True(t, is)

	is = isEmpty(board.G4, g)
	assert.True(t, is)

	is = isEmpty(board.G5, g)
	assert.True(t, is)

	is = isEmpty(board.G6, g)
	assert.True(t, is)

	is = isEmpty(board.G7, g)
	assert.False(t, is)

	is = isEmpty(board.G8, g)
	assert.False(t, is)

	is = isEmpty(board.H1, g)
	assert.False(t, is)

	is = isEmpty(board.H2, g)
	assert.False(t, is)

	is = isEmpty(board.H3, g)
	assert.True(t, is)

	is = isEmpty(board.H4, g)
	assert.True(t, is)

	is = isEmpty(board.H5, g)
	assert.True(t, is)

	is = isEmpty(board.H6, g)
	assert.True(t, is)

	is = isEmpty(board.H7, g)
	assert.False(t, is)

	is = isEmpty(board.H8, g)
	assert.False(t, is)
}

func TestCanOcuppy(t *testing.T) {
	g := NewGame()

	is := canOcuppy(board.WHITE_PAWN_A, board.A1, g)
	assert.False(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A2, g)
	assert.False(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A3, g)
	assert.True(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A4, g)
	assert.True(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A5, g)
	assert.True(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A6, g)
	assert.True(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A7, g)
	assert.True(t, is)

	is = canOcuppy(board.WHITE_PAWN_A, board.A8, g)
	assert.True(t, is)
}

func TestIsFreeVertical(t *testing.T) {
	g := NewGame()

	is := isFreeVertical(0, 1, 3, g)
	assert.True(t, is)

	is = isFreeVertical(0, 0, 3, g)
	assert.False(t, is)

	is = isFreeVertical(0, 6, 4, g)
	assert.True(t, is)

	is = isFreeVertical(0, 7, 4, g)
	assert.False(t, is)
}

func TestIsFreeHorizontal(t *testing.T) {
	g := NewGame()

	is := isFreeHorizontal(2, 1, 3, g)
	assert.True(t, is)

	is = isFreeHorizontal(0, 0, 3, g)
	assert.False(t, is)

	is = isFreeHorizontal(5, 6, 4, g)
	assert.True(t, is)

	is = isFreeHorizontal(7, 7, 4, g)
	assert.False(t, is)
}

func TestIsFreeDiagonal(t *testing.T) {
	g := NewGame()

	is := isFreeDiagonal(2, 0, 4, 2, g)
	assert.True(t, is)

	is = isFreeDiagonal(5, 0, 3, 2, g)
	assert.True(t, is)

	is = isFreeDiagonal(2, 7, 4, 5, g)
	assert.True(t, is)

	is = isFreeDiagonal(5, 7, 3, 5, g)
	assert.True(t, is)

	is = isFreeDiagonal(0, 0, 3, 3, g)
	assert.False(t, is)
}

func TestIsFreeDiagonalNE(t *testing.T) {
	g := NewGame()

	is := isFreeDiagonalNE(2, 0, 4, g)
	assert.True(t, is)

	is = isFreeDiagonalNE(0, 0, 4, g)
	assert.False(t, is)
}

func TestIsFreeDiagonalNW(t *testing.T) {
	g := NewGame()

	is := isFreeDiagonalNW(2, 7, 4, g)
	assert.True(t, is)

	is = isFreeDiagonalNW(0, 7, 2, g)
	assert.False(t, is)
}

func TestIsFreeDiagonalSE(t *testing.T) {
	g := NewGame()

	is := isFreeDiagonalSE(5, 0, 3, g)
	assert.True(t, is)

	is = isFreeDiagonalSE(4, 0, 0, g)
	assert.False(t, is)
}

func TestIsFreeDiagonalSW(t *testing.T) {
	g := NewGame()

	is := isFreeDiagonalSW(4, 7, 2, g)
	assert.True(t, is)

	is = isFreeDiagonalSW(2, 7, 0, g)
	assert.False(t, is)
}

func TestIsVertical(t *testing.T) {
	is := isVertical(1, 0, 2, 0)
	assert.True(t, is)

	is = isVertical(2, 0, 2, 0)
	assert.False(t, is)
}

func TestIsHorizontal(t *testing.T) {
	is := isHorizontal(2, 1, 2, 0)
	assert.True(t, is)

	is = isHorizontal(3, 0, 2, 0)
	assert.False(t, is)
}

func TestIsDiagonal(t *testing.T) {
	is := isDiagonal(0, 0, 2, 2)
	assert.True(t, is)

	is = isDiagonal(1, 0, 2, 0)
	assert.False(t, is)

	is = isDiagonal(2, 1, 2, 0)
	assert.False(t, is)

	is = isDiagonal(2, 1, 5, 0)
	assert.False(t, is)
}

func TestIsFreeForKing(t *testing.T) {
	g := NewGame()

	is := isFreeForKing(board.WHITE_KING, board.E1, board.E3, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_KING, board.E1, board.E3)
	is = isFreeForKing(board.WHITE_KING, board.E3, board.E4, g)
	assert.True(t, is)

	is = isFreeForKing(board.WHITE_KING, board.E3, board.G4, g)
	assert.False(t, is)

	is = isFreeForKing(board.WHITE_KING, board.E3, board.E2, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_KING, board.E3, board.E6)
	is = isFreeForKing(board.WHITE_KING, board.E6, board.E7, g)
	assert.True(t, is)
}

func TestIsFreeForQueen(t *testing.T) {
	g := NewGame()

	is := isFreeForQueen(board.WHITE_QUEEN, board.D1, board.D3, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_PAWN_D, board.D2, board.E3)
	is = isFreeForQueen(board.WHITE_QUEEN, board.D1, board.D5, g)
	assert.True(t, is)

	is = isFreeForQueen(board.WHITE_QUEEN, board.D1, board.D7, g)
	assert.True(t, is)

	is = isFreeForQueen(board.WHITE_QUEEN, board.D1, board.D8, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_QUEEN, board.D1, board.D3)
	is = isFreeForQueen(board.WHITE_QUEEN, board.D3, board.F7, g)
	assert.False(t, is)
}

func TestIsFreeForBishop(t *testing.T) {
	g := NewGame()

	is := isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.C1, board.E3, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_PAWN_D, board.D2, board.D3)
	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.C1, board.E3, g)
	assert.True(t, is)

	g.Board.Move(board.WHITE_BISHOP_QUEEN, board.C1, board.E3)
	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.E3, board.F4, g)
	assert.True(t, is)

	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.E3, board.E4, g)
	assert.False(t, is)

	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.E3, board.A7, g)
	assert.True(t, is)

	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.E3, board.G1, g)
	assert.False(t, is)

	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.E3, board.F2, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_BISHOP_QUEEN, board.E3, board.F4)
	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.F4, board.C7, g)
	assert.True(t, is)

	is = isFreeForBishop(board.WHITE_BISHOP_QUEEN, board.F4, board.B8, g)
	assert.False(t, is)
}

func TestIsFreeForKnight(t *testing.T) {
	g := NewGame()

	is := isFreeForKnight(board.WHITE_KNIGHT_QUEEN, board.B1, board.C3, g)
	assert.True(t, is)

	is = isFreeForKnight(board.WHITE_KNIGHT_QUEEN, board.B1, board.D2, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_KNIGHT_QUEEN, board.B1, board.B6)
	is = isFreeForKnight(board.WHITE_KNIGHT_QUEEN, board.B6, board.C8, g)
	assert.True(t, is)

	is = isFreeForKnight(board.WHITE_KNIGHT_QUEEN, board.B6, board.D7, g)
	assert.True(t, is)
}

func TestIsFreeForRook(t *testing.T) {
	g := NewGame()

	is := isFreeForRook(board.WHITE_ROOK_QUEEN, board.A1, board.A3, g)
	assert.False(t, is)

	is = isFreeForRook(board.WHITE_ROOK_QUEEN, board.A1, board.A2, g)
	assert.False(t, is)

	g.Board.Move(board.WHITE_ROOK_QUEEN, board.A1, board.A3)
	is = isFreeForRook(board.WHITE_ROOK_QUEEN, board.A3, board.A5, g)
	assert.True(t, is)

	is = isFreeForRook(board.WHITE_ROOK_QUEEN, board.A3, board.G3, g)
	assert.True(t, is)

	is = isFreeForRook(board.WHITE_ROOK_QUEEN, board.A3, board.C5, g)
	assert.False(t, is)

	is = isFreeForRook(board.WHITE_ROOK_QUEEN, board.A3, board.A7, g)
	assert.True(t, is)

	is = isFreeForRook(board.WHITE_ROOK_QUEEN, board.A3, board.A8, g)
	assert.False(t, is)
}

func TestIsFreeForPawn(t *testing.T) {
	g := NewGame()

	is := isFreeForPawn(board.WHITE_PAWN_A, board.A2, board.A3, g)
	assert.True(t, is)

	g.Board.Move(board.WHITE_PAWN_A, board.A2, board.B1)
	is = isFreeForPawn(board.WHITE_PAWN_A, board.B1, board.B3, g)
	assert.False(t, is)

	is = isFreeForPawn(board.WHITE_PAWN_A, board.B1, board.B2, g)
	assert.False(t, is)
}

func TestIsMoveFree(t *testing.T) {
	var is bool
	g := NewGame()

	is = isMoveFree(board.WHITE_PAWN_A, board.A2, board.A3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_B, board.B2, board.B3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_C, board.C2, board.C3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_D, board.D2, board.D3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_E, board.E2, board.E3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_F, board.F2, board.F3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_G, board.G2, board.G3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_PAWN_H, board.H2, board.H3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_ROOK_QUEEN, board.A1, board.A4, g)
	assert.False(t, is)

	is = isMoveFree(board.WHITE_ROOK_KING, board.H1, board.H4, g)
	assert.False(t, is)

	is = isMoveFree(board.WHITE_KNIGHT_QUEEN, board.B1, board.C3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_KNIGHT_KING, board.G1, board.H3, g)
	assert.True(t, is)

	is = isMoveFree(board.WHITE_BISHOP_QUEEN, board.C1, board.D4, g)
	assert.False(t, is)

	is = isMoveFree(board.WHITE_BISHOP_KING, board.F1, board.H3, g)
	assert.False(t, is)

	is = isMoveFree(board.WHITE_QUEEN, board.D1, board.D4, g)
	assert.False(t, is)

	is = isMoveFree(board.WHITE_KING, board.E1, board.E4, g)
	assert.False(t, is)

	is = isMoveFree(board.WHITE_KING, board.E1, board.E2, g)
	assert.False(t, is)

	is = isMoveFree(board.BLACK_PAWN_A, board.A7, board.A6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_B, board.B7, board.B6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_C, board.C7, board.C6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_D, board.D7, board.D6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_E, board.E7, board.E6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_F, board.F7, board.F6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_G, board.G7, board.G6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_PAWN_H, board.H7, board.H6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_ROOK_QUEEN, board.A8, board.A5, g)
	assert.False(t, is)

	is = isMoveFree(board.BLACK_ROOK_KING, board.H8, board.H5, g)
	assert.False(t, is)

	is = isMoveFree(board.BLACK_KNIGHT_QUEEN, board.B8, board.C6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_KNIGHT_KING, board.G8, board.F6, g)
	assert.True(t, is)

	is = isMoveFree(board.BLACK_BISHOP_QUEEN, board.C8, board.E6, g)
	assert.False(t, is)

	is = isMoveFree(board.BLACK_BISHOP_KING, board.F8, board.H6, g)
	assert.False(t, is)

	is = isMoveFree(board.BLACK_QUEEN, board.D8, board.D4, g)
	assert.False(t, is)

	is = isMoveFree(board.BLACK_KING, board.E8, board.E6, g)
	assert.False(t, is)

}
