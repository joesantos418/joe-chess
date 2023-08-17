package rules

import "github.com/joesantos418/joe-chess/pkg/board"

func isCastle(p board.Piece, to board.SquareName) bool {
	if p == board.WHITE_KING {
		return to == board.G1 || to == board.C1
	} else if p == board.BLACK_KING {
		return to == board.G8 || to == board.C8
	}

	return false
}

func canCastle(p board.Piece, from, to board.SquareName, g *Game) (bool, error) {
	if !isKing(p) {
		return false, nil
	}

	if hasKingMoved(p, g) {
		return false, nil
	}

	if hasRookMoved(to, g) {
		return false, nil
	}

	isInCheck := checkKingInCheck(g, g.Board)
	if isInCheck {
		return false, nil
	}

	isFree := isPathFree(p, from, to, g)
	if !isFree {
		return false, nil
	}

	return true, nil
}

func isKing(p board.Piece) bool {
	return p == board.WHITE_KING || p == board.BLACK_KING
}

func hasKingMoved(p board.Piece, g *Game) bool {
	if p == board.WHITE_KING {
		return g.HasWhiteKingMoved
	}

	if p == board.BLACK_KING {
		return g.HasBlackKingMoved
	}

	return false
}

func hasRookMoved(to board.SquareName, g *Game) bool {
	switch to {
	case board.G1:
		return g.HasWhiteRookKingMoved
	case board.C1:
		return g.HasWhiteRookQueenMoved
	case board.G8:
		return g.HasBlackRookKingMoved
	case board.C8:
		return g.HasBlackRookQueenMoved
	}

	return false
}

func isPathFree(p board.Piece, from, to board.SquareName, g *Game) bool {
	switch to {
	case board.G1:
		return isPathFreeForShortWhiteCastle(g)
	case board.C1:
		return isPathFreeForLongWhiteCastle(g)
	case board.G8:
		return isPathFreeForShortBlackCastle(g)
	case board.C8:
		return isPathFreeForLongBlackCastle(g)
	}

	return false
}

func isPathFreeForShortWhiteCastle(g *Game) bool {
	p := g.Board.GetPiece(board.F1)
	lin, col := board.GetCoordinatesBySquareName(board.F1)
	return p == board.NO_PIECE && !isSquareInCheck(lin, col, g.Board, g.NowPlays)
}

func isPathFreeForLongWhiteCastle(g *Game) bool {
	p := g.Board.GetPiece(board.D1)
	lin, col := board.GetCoordinatesBySquareName(board.D1)
	return p == board.NO_PIECE && !isSquareInCheck(lin, col, g.Board, g.NowPlays)
}

func isPathFreeForShortBlackCastle(g *Game) bool {
	p := g.Board.GetPiece(board.F8)
	lin, col := board.GetCoordinatesBySquareName(board.F8)
	return p == board.NO_PIECE && !isSquareInCheck(lin, col, g.Board, g.NowPlays)
}

func isPathFreeForLongBlackCastle(g *Game) bool {
	p := g.Board.GetPiece(board.D8)
	lin, col := board.GetCoordinatesBySquareName(board.D8)
	return p == board.NO_PIECE && !isSquareInCheck(lin, col, g.Board, g.NowPlays)
}
