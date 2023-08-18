package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func isValidMove(p board.Piece, from, to board.SquareName, g *Game) (bool, InvalidMoveReason, error) {
	if p == board.NO_PIECE {
		return false, EMPTY_SQUARE, nil
	}

	is := isColorsTurn(p, g)
	if !is {
		return false, NOT_COLORS_TURN, nil
	}

	is = isMoveCorrect(p, from, to, g)
	if !is {
		return false, MOVE_INCORRECT, nil
	}

	is = isMoveFree(p, from, to, g)
	if !is {
		return false, MOVE_OBSTRUCTED, nil
	}

	is, err := isKingInCheck(p, from, to, g)

	if err != nil {
		return false, ERROR_CHECKING, err
	}

	if is {
		return false, KING_IN_CHECK, nil
	}

	return true, VALID_MOVE, nil
}
