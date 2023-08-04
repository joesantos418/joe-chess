package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func isColorsTurn(p board.Piece, g *Game) bool {
	return (isPieceWhite(p) && g.NowPlays == WHITES) || (isPieceBlack(p) && g.NowPlays == BLACKS)
}

func isPieceWhite(p board.Piece) bool {
	return p == board.WHITE_PAWN_A ||
		p == board.WHITE_PAWN_B ||
		p == board.WHITE_PAWN_C ||
		p == board.WHITE_PAWN_D ||
		p == board.WHITE_PAWN_E ||
		p == board.WHITE_PAWN_F ||
		p == board.WHITE_PAWN_G ||
		p == board.WHITE_PAWN_H ||
		p == board.WHITE_ROOK_QUEEN ||
		p == board.WHITE_ROOK_KING ||
		p == board.WHITE_KNIGHT_QUEEN ||
		p == board.WHITE_KNIGHT_KING ||
		p == board.WHITE_BISHOP_QUEEN ||
		p == board.WHITE_BISHOP_KING ||
		p == board.WHITE_QUEEN ||
		p == board.WHITE_KING
}

func isPieceBlack(p board.Piece) bool {
	return p == board.BLACK_PAWN_A ||
		p == board.BLACK_PAWN_B ||
		p == board.BLACK_PAWN_C ||
		p == board.BLACK_PAWN_D ||
		p == board.BLACK_PAWN_E ||
		p == board.BLACK_PAWN_F ||
		p == board.BLACK_PAWN_G ||
		p == board.BLACK_PAWN_H ||
		p == board.BLACK_ROOK_QUEEN ||
		p == board.BLACK_ROOK_KING ||
		p == board.BLACK_KNIGHT_QUEEN ||
		p == board.BLACK_KNIGHT_KING ||
		p == board.BLACK_BISHOP_QUEEN ||
		p == board.BLACK_BISHOP_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}
