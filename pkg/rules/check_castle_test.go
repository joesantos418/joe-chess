package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsCastle_NoKing(t *testing.T) {
	is := isCastle(board.BLACK_BISHOP_KING, board.G4)
	assert.False(t, is)
}

func TestIsCastle_NoCastle(t *testing.T) {
	is := isCastle(board.WHITE_KING, board.E2)
	assert.False(t, is)
}

func TestIsCastle_Castle(t *testing.T) {
	is := isCastle(board.WHITE_KING, board.C1)
	assert.True(t, is)
}

func TestIsKing_Pawn(t *testing.T) {
	is := isKing(board.WHITE_PAWN_A)
	assert.False(t, is)
}

func TestIsKing_King(t *testing.T) {
	is := isKing(board.WHITE_KING)
	assert.True(t, is)
}

func TestHasKingMoved_NewGame(t *testing.T) {
	g := NewGame()
	has := hasKingMoved(board.WHITE_KING, g)

	assert.False(t, has)
}

func TestHasKingMoved_No(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	assert.Nil(t, err)

	has := hasKingMoved(board.WHITE_KING, g)
	assert.False(t, has)
}

func TestHasKingMoved_Yes(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	assert.Nil(t, err)

	err = g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
	assert.Nil(t, err)

	err = g.ProcessMove(board.WHITE_KING, board.E1, board.E2)
	assert.Nil(t, err)

	has := hasKingMoved(board.WHITE_KING, g)
	assert.True(t, has)
}

func TestHasRookMovedNewGame(t *testing.T) {
	g := NewGame()
	has := hasRookMoved(board.E4, g)

	assert.False(t, has)
}

func TestHasRookMoved_NoCastle(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_A, board.A2, board.A4)
	assert.Nil(t, err)

	err = g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
	assert.Nil(t, err)

	has := hasRookMoved(board.A2, g)
	assert.False(t, has)
}

func TestHasRookMoved_CastleMoved(t *testing.T) {
	g := NewGame()

	g.Board.SetPiece(board.NO_PIECE, board.B1)
	g.Board.SetPiece(board.NO_PIECE, board.C1)
	g.Board.SetPiece(board.NO_PIECE, board.D1)
	g.Board.SetPiece(board.NO_PIECE, board.A2)

	err := g.ProcessMove(board.WHITE_ROOK_QUEEN, board.A1, board.A4)
	assert.Nil(t, err)

	has := hasRookMoved(board.C1, g)
	assert.True(t, has)
}

func TestHasRookMoved_Castle(t *testing.T) {
	g := NewGame()

	g.Board.SetPiece(board.NO_PIECE, board.B1)
	g.Board.SetPiece(board.NO_PIECE, board.C1)
	g.Board.SetPiece(board.NO_PIECE, board.D1)
	g.Board.SetPiece(board.NO_PIECE, board.A2)

	has := hasRookMoved(board.C1, g)
	assert.False(t, has)
}

func TestIsPathFree(t *testing.T) {

}
