package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func setupBoard() *board.Board {
	b := &board.Board{}

	setWhites(b)
	setBlacks(b)

	return b
}

func setWhites(b *board.Board) {
	b.SetPiece(board.WHITE_PAWN_A, board.A2)
	b.SetPiece(board.WHITE_PAWN_B, board.B2)
	b.SetPiece(board.WHITE_PAWN_C, board.C2)
	b.SetPiece(board.WHITE_PAWN_D, board.D2)
	b.SetPiece(board.WHITE_PAWN_E, board.E2)
	b.SetPiece(board.WHITE_PAWN_F, board.F2)
	b.SetPiece(board.WHITE_PAWN_G, board.G2)
	b.SetPiece(board.WHITE_PAWN_H, board.H2)

	b.SetPiece(board.WHITE_ROOK_QUEEN, board.A1)
	b.SetPiece(board.WHITE_ROOK_KING, board.H1)
	b.SetPiece(board.WHITE_KNIGHT_QUEEN, board.B1)
	b.SetPiece(board.WHITE_KNIGHT_KING, board.G1)
	b.SetPiece(board.WHITE_BISHOP_QUEEN, board.C1)
	b.SetPiece(board.WHITE_BISHOP_KING, board.F1)
	b.SetPiece(board.WHITE_QUEEN, board.D1)
	b.SetPiece(board.WHITE_KING, board.E1)
}

func setBlacks(b *board.Board) {
	b.SetPiece(board.BLACK_PAWN_A, board.A7)
	b.SetPiece(board.BLACK_PAWN_B, board.B7)
	b.SetPiece(board.BLACK_PAWN_C, board.C7)
	b.SetPiece(board.BLACK_PAWN_D, board.D7)
	b.SetPiece(board.BLACK_PAWN_E, board.E7)
	b.SetPiece(board.BLACK_PAWN_F, board.F7)
	b.SetPiece(board.BLACK_PAWN_G, board.G7)
	b.SetPiece(board.BLACK_PAWN_H, board.H7)

	b.SetPiece(board.BLACK_ROOK_QUEEN, board.A8)
	b.SetPiece(board.BLACK_ROOK_KING, board.H8)
	b.SetPiece(board.BLACK_KNIGHT_QUEEN, board.B8)
	b.SetPiece(board.BLACK_KNIGHT_KING, board.G8)
	b.SetPiece(board.BLACK_BISHOP_QUEEN, board.C8)
	b.SetPiece(board.BLACK_BISHOP_KING, board.F8)
	b.SetPiece(board.BLACK_QUEEN, board.D8)
	b.SetPiece(board.BLACK_KING, board.E8)
}
