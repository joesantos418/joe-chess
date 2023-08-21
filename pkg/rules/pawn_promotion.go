package rules

import (
	"github.com/arquivei/foundationkit/errors"
	"github.com/joesantos418/joe-chess/pkg/board"
)

func canPromotePawn(p board.Piece, to board.SquareName) bool {
	if isWhitePawn(p) {
		return isEighthRank(to)
	} else if isBlackPawn(p) {
		return isFirstRank(to)
	}

	return false
}

func isEighthRank(to board.SquareName) bool {
	return to == board.A8 ||
		to == board.B8 ||
		to == board.C8 ||
		to == board.D8 ||
		to == board.E8 ||
		to == board.F8 ||
		to == board.G8 ||
		to == board.H8
}

func isFirstRank(to board.SquareName) bool {
	return to == board.A1 ||
		to == board.B1 ||
		to == board.C1 ||
		to == board.D1 ||
		to == board.E1 ||
		to == board.F1 ||
		to == board.G1 ||
		to == board.H1
}

func promotePawn(new board.Piece, to board.SquareName, g *Game) error {
	if !isValidPromotion(new, to) {
		return errors.E(
			"cannot promote pawn",
			errors.KV("promote_to", new),
			errors.KV("move_to", to),
		)
	}

	g.Board.SetPiece(new, to)
	return nil
}

func isValidPromotion(new board.Piece, to board.SquareName) bool {
	if isFirstRank(to) {
		return isValidBlackPromotion(new)
	} else if isEighthRank(to) {
		return isValidWhitePromotion(new)
	}

	return false
}

func isValidBlackPromotion(new board.Piece) bool {
	return new == board.BLACK_QUEEN ||
		new == board.BLACK_ROOK_QUEEN ||
		new == board.BLACK_ROOK_KING ||
		new == board.BLACK_BISHOP_QUEEN ||
		new == board.BLACK_BISHOP_KING ||
		new == board.BLACK_KNIGHT_QUEEN ||
		new == board.BLACK_KNIGHT_KING
}

func isValidWhitePromotion(new board.Piece) bool {
	return new == board.WHITE_QUEEN ||
		new == board.WHITE_ROOK_QUEEN ||
		new == board.WHITE_ROOK_KING ||
		new == board.WHITE_BISHOP_QUEEN ||
		new == board.WHITE_BISHOP_KING ||
		new == board.WHITE_KNIGHT_QUEEN ||
		new == board.WHITE_KNIGHT_KING
}
