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

func TestIsThreatenedNE(t *testing.T) {
	b := &board.Board{}

	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.D4)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F6)

	is := isThreatenedNE(board.WHITE_BISHOP_QUEEN, 3, 3, b)
	assert.True(t, is)
}
