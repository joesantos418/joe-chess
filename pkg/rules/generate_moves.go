package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func generateMoves(c Color, b board.Board) []Move {
	var moves []Move
	for _, n := range board.SquareNames {
		p := b.GetPiece(n)
		if isColor(p, c) {
			moves = append(moves, generateMovesForPiece(p, b))
		}
	}

	return moves
}

func isColor(p board.Piece, c Color) bool {
	return false
}

func generateMovesForPiece(p board.Piece, b board.Board) []Move {
	var moves []Move

	if movesVertically(p) {
		moves = append(moves, generateVericalMoves(p, b))
	}

	if movesHorizontally(p) {
		moves = append(moves, generateHorizontalMoves(p, b))
	}

	if movesDiagonally(p) {
		moves = append(moves, generateDiagonalMoves(p, b))
	}

	if movesInL(p) {
		moves = append(moves, generateLMoves(p, b))
	}

	if isPawn(p) {
		// en passant as well
		moves = append(moves, generatePawnCaptureMoves(p, b))
	}

	if isKing(p) {
		// castle as well
		moves = append(moves, generateCastleMoves(p, b))
	}

	return moves
}
