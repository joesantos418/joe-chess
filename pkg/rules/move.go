package rules

import "github.com/joesantos418/joe-chess/pkg/board"

type Move struct {
	Piece board.Piece
	From  board.SquareName
	To    board.SquareName
}
