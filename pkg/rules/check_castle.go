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

	isInCheck, err := isKingInCheck(p, from, to, g)
	if err != nil {
		return false, err
	}

	if isInCheck {
		return false, nil
	}

	isFree, err := isPathFree(p, from, to, g)
	if err != nil {
		return false, err
	}

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

func isPathFree(p board.Piece, from, to board.SquareName, g *Game) (bool, error) {
	return false, nil
}
