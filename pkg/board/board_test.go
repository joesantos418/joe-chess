package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove_Success(t *testing.T) {
	b := &Board{}

	b.SetPiece(WHITE_PAWN_A, A1)
	err := b.Move(WHITE_PAWN_A, A1, B1)

	assert.Nil(t, err)
	assert.Equal(t, NO_PIECE, b.GetPiece(A1))
	assert.Equal(t, WHITE_PAWN_A, b.GetPiece(B1))
}

func TestMove_Error(t *testing.T) {
	b := &Board{}

	b.SetPiece(WHITE_PAWN_A, A1)
	err := b.Move(WHITE_PAWN_A, B1, C1)

	assert.NotNil(t, err)
}

func TestEmptyBoard(t *testing.T) {
	b := &Board{}

	for _, line := range b.Squares {
		for _, sq := range line {
			assert.Equal(t, NO_PIECE, sq.Piece)
		}
	}
}

func TestSetPiece(t *testing.T) {
	b := &Board{}

	b.SetPiece(WHITE_PAWN_A, A1)

	for l, line := range b.Squares {
		for c, sq := range line {
			if l == 0 && c == 0 {
				assert.Equal(t, WHITE_PAWN_A, sq.Piece)
			} else {
				assert.Equal(t, NO_PIECE, sq.Piece)
			}
		}
	}
}

func TestGetPiece(t *testing.T) {
	b := &Board{}

	b.Squares[0][0] = Square{Piece: WHITE_PAWN_A}

	assert.Equal(t, WHITE_PAWN_A, b.GetPiece(A1))
	assert.Equal(t, NO_PIECE, b.GetPiece(C1))
	assert.Equal(t, NO_PIECE, b.GetPiece(H4))
}
