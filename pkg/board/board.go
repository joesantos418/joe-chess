package board

import (
	"github.com/arquivei/foundationkit/errors"
)

const MAX_COL = 8
const MAX_LIN = 8

type Square struct {
	Piece Piece
}

type Board struct {
	Squares [8][8]Square
}

func (b *Board) Move(p Piece, from, to SquareName) error {
	lin, col := GetCoordinatesBySquareName(from)
	if b.Squares[lin][col].Piece != p {
		return errors.E(
			"cannot find piece to move",
			errors.KV("piece", p),
			errors.KV("from", from),
			errors.KV("to", to),
		)
	}

	b.Squares[lin][col].Piece = NO_PIECE
	lin, col = GetCoordinatesBySquareName(to)
	b.Squares[lin][col].Piece = p

	return nil
}

func (b *Board) SetPiece(p Piece, n SquareName) {
	lin, col := GetCoordinatesBySquareName(n)
	b.Squares[lin][col].Piece = p
}

func (b *Board) GetPiece(n SquareName) Piece {
	lin, col := GetCoordinatesBySquareName(n)
	return b.Squares[lin][col].Piece
}

func (b *Board) Clone() *Board {
	c := &Board{}

	for lin := 0; lin < MAX_LIN; lin++ {
		for col := 0; col < MAX_COL; col++ {
			c.Squares[lin][col].Piece = b.Squares[lin][col].Piece
		}
	}

	return c
}
