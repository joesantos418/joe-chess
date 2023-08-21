package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsTargetVulnerableToBlackEnPassant(t *testing.T) {
	g := NewGame()
	is := isTargetVulnerableToBlackEnPassant(board.A3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.B3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.C3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.D3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.E3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.F3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.G3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.H3, g)
	assert.False(t, is)

	is = isTargetVulnerableToBlackEnPassant(board.D4, g)
	assert.False(t, is)
}

func TestIsTargetVulnerableToWhiteEnPassant(t *testing.T) {
	g := NewGame()
	is := isTargetVulnerableToWhiteEnPassant(board.A6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.B6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.C6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.D6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.E6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.F6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.G6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.H6, g)
	assert.False(t, is)

	is = isTargetVulnerableToWhiteEnPassant(board.D4, g)
	assert.False(t, is)
}

func TestIsBlackEnPassantMove(t *testing.T) {
	is := isBlackEnPassantMove(board.A4, board.B3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.B4, board.C3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.C4, board.D3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.D4, board.E3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.E4, board.F3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.F4, board.G3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.G4, board.H3)
	assert.True(t, is)

	is = isBlackEnPassantMove(board.H4, board.G3)
	assert.True(t, is)
}

func TestIsWhiteEnPassantMove(t *testing.T) {
	is := isWhiteEnPassantMove(board.A5, board.B6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.B5, board.C6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.C5, board.D6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.D5, board.E6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.E5, board.F6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.F5, board.G6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.G5, board.H6)
	assert.True(t, is)

	is = isWhiteEnPassantMove(board.H5, board.G6)
	assert.True(t, is)
}

func TestIsWhitePawn(t *testing.T) {
	is := isWhitePawn(board.BLACK_BISHOP_KING)
	assert.False(t, is)

	is = isWhitePawn(board.WHITE_PAWN_A)
	assert.True(t, is)
}

func TestIsBlackPawn(t *testing.T) {
	is := isBlackPawn(board.BLACK_BISHOP_KING)
	assert.False(t, is)

	is = isBlackPawn(board.BLACK_PAWN_A)
	assert.True(t, is)
}

func TestIsWhiteEnPassant_NotWhitePawn(t *testing.T) {
	p := board.WHITE_BISHOP_KING
	from := board.A1
	to := board.B2
	g := NewGame()
	is := isWhiteEnPassant(p, from, to, g)

	assert.False(t, is)
}

func TestIsWhiteEnPassant_NotEnPassantMove(t *testing.T) {
	p := board.WHITE_PAWN_A
	from := board.A2
	to := board.B4
	g := NewGame()
	is := isWhiteEnPassant(p, from, to, g)

	assert.False(t, is)
}

func TestIsWhiteEnPassant_NotVulnerableTarget(t *testing.T) {
	p := board.WHITE_PAWN_E
	from := board.E5
	to := board.D6
	g := NewGame()

	g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E6)
	g.ProcessMove(board.WHITE_PAWN_E, board.E4, board.E5)
	g.ProcessMove(board.BLACK_PAWN_D, board.D7, board.D5)

	g.HasBlackPawnDJustMovedTwoSquares = false

	is := isWhiteEnPassant(p, from, to, g)

	assert.False(t, is)
}

func TestIsWhiteEnPassant_Yes(t *testing.T) {
	p := board.WHITE_PAWN_E
	from := board.E5
	to := board.D6
	g := NewGame()

	g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E6)
	g.ProcessMove(board.WHITE_PAWN_E, board.E4, board.E5)
	g.ProcessMove(board.BLACK_PAWN_D, board.D7, board.D5)

	is := isWhiteEnPassant(p, from, to, g)

	assert.True(t, is)
}

func TestIsBlackEnPassant_NotBlackPawn(t *testing.T) {
	p := board.BLACK_BISHOP_KING
	from := board.A8
	to := board.B7
	g := NewGame()
	is := isBlackEnPassant(p, from, to, g)

	assert.False(t, is)
}

func TestIsBlackEnPassant_NotEnPassantMove(t *testing.T) {
	p := board.BLACK_PAWN_A
	from := board.A7
	to := board.B6
	g := NewGame()
	is := isBlackEnPassant(p, from, to, g)

	assert.False(t, is)
}

func TestIsBlackEnPassant_NotVulnerableTarget(t *testing.T) {
	p := board.BLACK_PAWN_E
	from := board.E4
	to := board.D3
	g := NewGame()

	g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E3)
	g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
	g.ProcessMove(board.WHITE_PAWN_A, board.A2, board.A4)
	g.ProcessMove(board.BLACK_PAWN_E, board.E5, board.E4)
	g.ProcessMove(board.WHITE_PAWN_D, board.D2, board.D4)

	g.HasWhitePawnDJustMovedTwoSquares = false

	is := isBlackEnPassant(p, from, to, g)

	assert.False(t, is)
}

func TestIsBlackEnPassant_Yes(t *testing.T) {
	p := board.BLACK_PAWN_E
	from := board.E4
	to := board.D3
	g := NewGame()

	g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E3)
	g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
	g.ProcessMove(board.WHITE_PAWN_A, board.A2, board.A4)
	g.ProcessMove(board.BLACK_PAWN_E, board.E5, board.E4)
	g.ProcessMove(board.WHITE_PAWN_D, board.D2, board.D4)

	is := isBlackEnPassant(p, from, to, g)

	assert.True(t, is)
}
