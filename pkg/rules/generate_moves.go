package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func generateMoves(c Color, b board.Board) []Move {
	var moves []Move
	for _, n := range board.SquareNames {
		p := b.GetPiece(n)
		if isColor(p, c) {
			moves = append(moves, generateMovesForPiece(p, n, b)...)
		}
	}

	return moves
}

func isColor(p board.Piece, c Color) bool {
	if c == WHITES {
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

	if c == BLACKS {
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

	panic("unexpected piece")
}

// we gotta consider more than one of the same piece on the board
func generateMovesForPiece(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move

	if movesVertically(p) {
		moves = append(moves, generateVerticalMoves(p, n, b)...)
	}

	if movesHorizontally(p) {
		moves = append(moves, generateHorizontalMoves(p, n, b)...)
	}

	if movesDiagonally(p) {
		moves = append(moves, generateDiagonalMoves(p, n, b)...)
	}

	if movesInL(p) {
		moves = append(moves, generateLMoves(p, n, b)...)
	}

	if isPawn(p) {
		// en passant as well
		moves = append(moves, generatePawnCaptureMoves(p, n, b)...)
	}

	if isKing(p) {
		// castle as well
		moves = append(moves, generateCastleMoves(p, n, b)...)
	}

	return moves
}

func movesVertically(p board.Piece) bool {
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
		p == board.WHITE_QUEEN ||
		p == board.WHITE_KING ||
		p == board.BLACK_PAWN_A ||
		p == board.BLACK_PAWN_B ||
		p == board.BLACK_PAWN_C ||
		p == board.BLACK_PAWN_D ||
		p == board.BLACK_PAWN_E ||
		p == board.BLACK_PAWN_F ||
		p == board.BLACK_PAWN_G ||
		p == board.BLACK_PAWN_H ||
		p == board.BLACK_ROOK_QUEEN ||
		p == board.BLACK_ROOK_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}

func movesHorizontally(p board.Piece) bool {
	return p == board.WHITE_ROOK_QUEEN ||
		p == board.WHITE_ROOK_KING ||
		p == board.WHITE_QUEEN ||
		p == board.WHITE_KING ||
		p == board.BLACK_ROOK_QUEEN ||
		p == board.BLACK_ROOK_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}

func movesDiagonally(p board.Piece) bool {
	return p == board.WHITE_BISHOP_QUEEN ||
		p == board.WHITE_BISHOP_KING ||
		p == board.WHITE_QUEEN ||
		p == board.WHITE_KING ||
		p == board.BLACK_BISHOP_QUEEN ||
		p == board.BLACK_BISHOP_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}

func movesInL(p board.Piece) bool {
	return p == board.WHITE_KNIGHT_QUEEN ||
		p == board.WHITE_KNIGHT_KING ||
		p == board.BLACK_KNIGHT_QUEEN ||
		p == board.BLACK_KNIGHT_KING
}

func isPawn(p board.Piece) bool {
	return p == board.WHITE_PAWN_A ||
		p == board.WHITE_PAWN_B ||
		p == board.WHITE_PAWN_C ||
		p == board.WHITE_PAWN_D ||
		p == board.WHITE_PAWN_E ||
		p == board.WHITE_PAWN_F ||
		p == board.WHITE_PAWN_G ||
		p == board.WHITE_PAWN_H ||
		p == board.BLACK_PAWN_A ||
		p == board.BLACK_PAWN_B ||
		p == board.BLACK_PAWN_C ||
		p == board.BLACK_PAWN_D ||
		p == board.BLACK_PAWN_E ||
		p == board.BLACK_PAWN_F ||
		p == board.BLACK_PAWN_G ||
		p == board.BLACK_PAWN_H
}

func generateVerticalMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move

	moves = append(moves, generateVerticalMovesUp(p, n, b)...)
	moves = append(moves, generateVerticalMovesDown(p, n, b)...)

	return moves
}

func generateVerticalMovesUp(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move
	return moves
}

func generateVerticalMovesDown(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move
	return moves
}

func generateHorizontalMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move

	moves = append(moves, generateHorizontalMovesLeft(p, n, b)...)
	moves = append(moves, generateHorizontalMovesRight(p, n, b)...)

	return moves
}

func generateHorizontalMovesLeft(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateHorizontalMovesRight(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateDiagonalMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move

	moves = append(moves, generateDiagonalMovesNE(p, n, b)...)
	moves = append(moves, generateDiagonalMovesNW(p, n, b)...)
	moves = append(moves, generateDiagonalMovesSE(p, n, b)...)
	moves = append(moves, generateDiagonalMovesSW(p, n, b)...)

	return moves
}

func generateDiagonalMovesNE(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateDiagonalMovesNW(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateDiagonalMovesSE(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateDiagonalMovesSW(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateLMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generatePawnCaptureMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateCastleMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}
