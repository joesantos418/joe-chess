package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsFirstRank_A1(t *testing.T) {
	is := isFirstRank(board.A1)
	assert.True(t, is)
}

func TestIsFirstRank_B3(t *testing.T) {
	is := isFirstRank(board.B3)
	assert.False(t, is)
}

func TestIsEighthRank_A8(t *testing.T) {
	is := isEighthRank(board.A8)
	assert.True(t, is)
}

func TestIsEighthRank_B3(t *testing.T) {
	is := isEighthRank(board.B3)
	assert.False(t, is)
}

func TestCanPromotePawn_WhitePawnEighthRank(t *testing.T) {
	can := canPromotePawn(board.WHITE_PAWN_A, board.A8)
	assert.True(t, can)
}

func TestCanPromotePawn_WhitePawnThirdRank(t *testing.T) {
	can := canPromotePawn(board.WHITE_PAWN_A, board.A3)
	assert.False(t, can)
}

func TestCanPromotePawn_BlackPawnFirstRank(t *testing.T) {
	can := canPromotePawn(board.BLACK_PAWN_A, board.A1)
	assert.True(t, can)
}

func TestCanPromotePawn_BlackPawnThirdRank(t *testing.T) {
	can := canPromotePawn(board.BLACK_PAWN_A, board.A3)
	assert.False(t, can)
}

func TestCanPromotePawn_NoPiece(t *testing.T) {
	can := canPromotePawn(board.NO_PIECE, board.A3)
	assert.False(t, can)
}

func TestIsValidWhitePromotion_Pawn(t *testing.T) {
	is := isValidWhitePromotion(board.WHITE_PAWN_A)
	assert.False(t, is)
}

func TestIsValidWhitePromotion_Queen(t *testing.T) {
	is := isValidWhitePromotion(board.WHITE_QUEEN)
	assert.True(t, is)
}

func TestIsValidBlackPromotion_Pawn(t *testing.T) {
	is := isValidBlackPromotion(board.BLACK_PAWN_A)
	assert.False(t, is)
}

func TestIsValidBlackPromotion_Queen(t *testing.T) {
	is := isValidBlackPromotion(board.BLACK_QUEEN)
	assert.True(t, is)
}

func TestIsValidPromotion_ThirdRank(t *testing.T) {
	is := isValidPromotion(board.WHITE_PAWN_A, board.A3)
	assert.False(t, is)
}

func TestIsValidPromotion_WhitePawnFirstRank(t *testing.T) {
	is := isValidPromotion(board.WHITE_PAWN_A, board.A1)
	assert.False(t, is)
}

func TestIsValidPromotion_WhitePawnEighthRank(t *testing.T) {
	is := isValidPromotion(board.WHITE_PAWN_A, board.A8)
	assert.False(t, is)
}

func TestPromotePawn(t *testing.T) {
	new := board.WHITE_QUEEN
	to := board.A8
	g := NewGame()

	err := promotePawn(new, to, g)

	assert.Nil(t, err)
}

func TestPromotePawn_Invalid(t *testing.T) {
	new := board.WHITE_PAWN_A
	to := board.A8
	g := NewGame()

	err := promotePawn(new, to, g)

	assert.NotNil(t, err)
}
