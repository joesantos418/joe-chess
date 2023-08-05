package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func isValidMove(p board.Piece, from, to board.SquareName, g *Game) (bool, InvalidMoveReason) {
	if p == board.NO_PIECE {
		return false, EMPTY_SQUARE
	}

	is := isColorsTurn(p, g)
	if !is {
		return false, NOT_COLORS_TURN
	}

	is = isMoveCorrect(p, from, to, g)
	if !is {
		return false, MOVE_INCORRECT
	}

	is = isMoveFree(p, from, to, g)
	if !is {
		return false, MOVE_OBSTRUCTED
	}

	is = isKingInCheck(p, from, to, g)
	if !is {
		return false, KING_IN_CHECK
	}

	return true, VALID_MOVE
}
