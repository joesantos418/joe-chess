package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func isWhiteEnPassant(p board.Piece, from, to board.SquareName, g *Game) bool {
	if isWhitePawn(p) && isWhiteEnPassantMove(from, to) && isTargetVulnerableToWhiteEnPassant(to, g) {
		return true
	}

	return false
}

func isBlackEnPassant(p board.Piece, from, to board.SquareName, g *Game) bool {
	if isBlackPawn(p) && isBlackEnPassantMove(from, to) && isTargetVulnerableToBlackEnPassant(to, g) {
		return true
	}

	return false
}

func isWhitePawn(p board.Piece) bool {
	return p == board.WHITE_PAWN_A ||
		p == board.WHITE_PAWN_B ||
		p == board.WHITE_PAWN_C ||
		p == board.WHITE_PAWN_D ||
		p == board.WHITE_PAWN_E ||
		p == board.WHITE_PAWN_F ||
		p == board.WHITE_PAWN_G ||
		p == board.WHITE_PAWN_H
}

func isBlackPawn(p board.Piece) bool {
	return p == board.BLACK_PAWN_A ||
		p == board.BLACK_PAWN_B ||
		p == board.BLACK_PAWN_C ||
		p == board.BLACK_PAWN_D ||
		p == board.BLACK_PAWN_E ||
		p == board.BLACK_PAWN_F ||
		p == board.BLACK_PAWN_G ||
		p == board.BLACK_PAWN_H
}

func isWhiteEnPassantMove(from, to board.SquareName) bool {
	switch from {
	case board.A5:
		return to == board.B6
	case board.B5:
		return to == board.A6 || to == board.C6
	case board.C5:
		return to == board.B6 || to == board.D6
	case board.D5:
		return to == board.C6 || to == board.E6
	case board.E5:
		return to == board.D6 || to == board.F6
	case board.F5:
		return to == board.E6 || to == board.G6
	case board.G5:
		return to == board.F6 || to == board.H6
	case board.H5:
		return to == board.G6
	default:
		return false
	}
}

func isBlackEnPassantMove(from, to board.SquareName) bool {
	switch from {
	case board.A4:
		return to == board.B3
	case board.B4:
		return to == board.A3 || to == board.C3
	case board.C4:
		return to == board.B3 || to == board.D3
	case board.D4:
		return to == board.C3 || to == board.E3
	case board.E4:
		return to == board.D3 || to == board.F3
	case board.F4:
		return to == board.E3 || to == board.G3
	case board.G4:
		return to == board.F3 || to == board.H3
	case board.H4:
		return to == board.G3
	default:
		return false
	}
}

func isTargetVulnerableToWhiteEnPassant(to board.SquareName, g *Game) bool {
	switch to {
	case board.A6:
		return g.HasBlackPawnAJustMovedTwoSquares
	case board.B6:
		return g.HasBlackPawnBJustMovedTwoSquares
	case board.C6:
		return g.HasBlackPawnCJustMovedTwoSquares
	case board.D6:
		return g.HasBlackPawnDJustMovedTwoSquares
	case board.E6:
		return g.HasBlackPawnEJustMovedTwoSquares
	case board.F6:
		return g.HasBlackPawnFJustMovedTwoSquares
	case board.G6:
		return g.HasBlackPawnGJustMovedTwoSquares
	case board.H6:
		return g.HasBlackPawnHJustMovedTwoSquares
	default:
		return false
	}
}

func isTargetVulnerableToBlackEnPassant(to board.SquareName, g *Game) bool {
	switch to {
	case board.A3:
		return g.HasWhitePawnAJustMovedTwoSquares
	case board.B3:
		return g.HasWhitePawnBJustMovedTwoSquares
	case board.C3:
		return g.HasWhitePawnCJustMovedTwoSquares
	case board.D3:
		return g.HasWhitePawnDJustMovedTwoSquares
	case board.E3:
		return g.HasWhitePawnEJustMovedTwoSquares
	case board.F3:
		return g.HasWhitePawnFJustMovedTwoSquares
	case board.G3:
		return g.HasWhitePawnGJustMovedTwoSquares
	case board.H3:
		return g.HasWhitePawnHJustMovedTwoSquares
	default:
		return false
	}
}
