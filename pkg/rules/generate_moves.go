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
	return false
}

// we gotta consider more than one of the same piece on the board
func generateMovesForPiece(p board.Piece, n board.SquareName, b board.Board) []Move {
	var moves []Move

	if movesVertically(p) {
		moves = append(moves, generateVericalMoves(p, n, b)...)
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
	return false
}

func movesHorizontally(p board.Piece) bool {
	return false
}

func movesDiagonally(p board.Piece) bool {
	return false
}

func movesInL(p board.Piece) bool {
	return false
}

func isPawn(p board.Piece) bool {
	return false
}

func generateVericalMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateHorizontalMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
	return []Move{}
}

func generateDiagonalMoves(p board.Piece, n board.SquareName, b board.Board) []Move {
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
