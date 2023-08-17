package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	p := NO_PIECE
	assert.Equal(t, "NO_PIECE", p.String())
	p = WHITE_PAWN_A
	assert.Equal(t, "WHITE_PAWN_A", p.String())
	p = WHITE_PAWN_B
	assert.Equal(t, "WHITE_PAWN_B", p.String())
	p = WHITE_PAWN_C
	assert.Equal(t, "WHITE_PAWN_C", p.String())
	p = WHITE_PAWN_D
	assert.Equal(t, "WHITE_PAWN_D", p.String())
	p = WHITE_PAWN_E
	assert.Equal(t, "WHITE_PAWN_E", p.String())
	p = WHITE_PAWN_F
	assert.Equal(t, "WHITE_PAWN_F", p.String())
	p = WHITE_PAWN_G
	assert.Equal(t, "WHITE_PAWN_G", p.String())
	p = WHITE_PAWN_H
	assert.Equal(t, "WHITE_PAWN_H", p.String())
	p = WHITE_ROOK_QUEEN
	assert.Equal(t, "WHITE_ROOK_QUEEN", p.String())
	p = WHITE_ROOK_KING
	assert.Equal(t, "WHITE_ROOK_KING", p.String())
	p = WHITE_KNIGHT_QUEEN
	assert.Equal(t, "WHITE_KNIGHT_QUEEN", p.String())
	p = WHITE_KNIGHT_KING
	assert.Equal(t, "WHITE_KNIGHT_KING", p.String())
	p = WHITE_BISHOP_QUEEN
	assert.Equal(t, "WHITE_BISHOP_QUEEN", p.String())
	p = WHITE_BISHOP_KING
	assert.Equal(t, "WHITE_BISHOP_KING", p.String())
	p = WHITE_QUEEN
	assert.Equal(t, "WHITE_QUEEN", p.String())
	p = WHITE_KING
	assert.Equal(t, "WHITE_KING", p.String())
	p = BLACK_PAWN_A
	assert.Equal(t, "BLACK_PAWN_A", p.String())
	p = BLACK_PAWN_B
	assert.Equal(t, "BLACK_PAWN_B", p.String())
	p = BLACK_PAWN_C
	assert.Equal(t, "BLACK_PAWN_C", p.String())
	p = BLACK_PAWN_D
	assert.Equal(t, "BLACK_PAWN_D", p.String())
	p = BLACK_PAWN_E
	assert.Equal(t, "BLACK_PAWN_E", p.String())
	p = BLACK_PAWN_F
	assert.Equal(t, "BLACK_PAWN_F", p.String())
	p = BLACK_PAWN_G
	assert.Equal(t, "BLACK_PAWN_G", p.String())
	p = BLACK_PAWN_H
	assert.Equal(t, "BLACK_PAWN_H", p.String())
	p = BLACK_ROOK_QUEEN
	assert.Equal(t, "BLACK_ROOK_QUEEN", p.String())
	p = BLACK_ROOK_KING
	assert.Equal(t, "BLACK_ROOK_KING", p.String())
	p = BLACK_KNIGHT_QUEEN
	assert.Equal(t, "BLACK_KNIGHT_QUEEN", p.String())
	p = BLACK_KNIGHT_KING
	assert.Equal(t, "BLACK_KNIGHT_KING", p.String())
	p = BLACK_BISHOP_QUEEN
	assert.Equal(t, "BLACK_BISHOP_QUEEN", p.String())
	p = BLACK_BISHOP_KING
	assert.Equal(t, "BLACK_BISHOP_KING", p.String())
	p = BLACK_QUEEN
	assert.Equal(t, "BLACK_QUEEN", p.String())
	p = BLACK_KING
	assert.Equal(t, "BLACK_KING", p.String())
}
