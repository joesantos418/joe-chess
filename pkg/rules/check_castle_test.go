package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsKing_Pawn(t *testing.T) {
	is := isKing(board.WHITE_PAWN_A)
	assert.False(t, is)
}

func TestIsKing_King(t *testing.T) {
	is := isKing(board.WHITE_KING)
	assert.True(t, is)
}

// errors here for tomorrow

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
