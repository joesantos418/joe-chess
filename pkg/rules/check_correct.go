package rules

import (
	"math"

	"github.com/joesantos418/joe-chess/pkg/board"
)

func isMoveCorrect(p board.Piece, from, to board.SquareName, g *Game) bool {
	if from == to {
		return true
	}

	switch p {
	case board.WHITE_PAWN_A,
		board.WHITE_PAWN_B,
		board.WHITE_PAWN_C,
		board.WHITE_PAWN_D,
		board.WHITE_PAWN_E,
		board.WHITE_PAWN_F,
		board.WHITE_PAWN_G,
		board.WHITE_PAWN_H:
		return isCorrectForWhitePawn(p, from, to, g)

	case board.BLACK_PAWN_A,
		board.BLACK_PAWN_B,
		board.BLACK_PAWN_C,
		board.BLACK_PAWN_D,
		board.BLACK_PAWN_E,
		board.BLACK_PAWN_F,
		board.BLACK_PAWN_G,
		board.BLACK_PAWN_H:
		return isCorrectForBlackPawn(p, from, to, g)

	case board.WHITE_ROOK_QUEEN,
		board.WHITE_ROOK_KING:
		return isCorrectForRook(from, to, g)
	case board.BLACK_ROOK_QUEEN,
		board.BLACK_ROOK_KING:
		return isCorrectForRook(from, to, g)

	case board.BLACK_KNIGHT_QUEEN,
		board.BLACK_KNIGHT_KING:
		return isCorrectForKnight(from, to, g)
	case board.WHITE_KNIGHT_QUEEN,
		board.WHITE_KNIGHT_KING:
		return isCorrectForKnight(from, to, g)

	case board.BLACK_BISHOP_QUEEN,
		board.BLACK_BISHOP_KING,
		board.WHITE_BISHOP_QUEEN,
		board.WHITE_BISHOP_KING:
		return isCorrectForBishop(from, to, g)

	case board.BLACK_QUEEN,
		board.WHITE_QUEEN:
		return isCorrectForQueen(from, to, g)

	case board.BLACK_KING,
		board.WHITE_KING:
		return isCorrectForKing(from, to, g)
	}

	return false
}

func isCorrectForWhitePawn(p board.Piece, from, to board.SquareName, g *Game) bool {
	if isCapture(p, to, g) {
		return isDiagonalUpOne(from, to)
	} else {
		if isFirstSquareForWhitePawn(from) {
			return isVerticalUpOne(from, to) || isVerticalUpTwo(from, to)
		} else {
			return isVerticalUpOne(from, to)
		}
	}
}

func isCorrectForBlackPawn(p board.Piece, from, to board.SquareName, g *Game) bool {
	if isCapture(p, to, g) {
		return isDiagonalDownOne(from, to)
	} else {
		if isFirstSquareForBlackPawn(from) {
			return isVerticalDownOne(from, to) || isVerticalDownTwo(from, to)
		} else {
			return isVerticalDownOne(from, to)
		}
	}
}

func isCorrectForRook(from, to board.SquareName, g *Game) bool {
	return isVerticalN(from, to) || isHorizontalN(from, to)
}

func isCorrectForKnight(from, to board.SquareName, g *Game) bool {
	return isL(from, to)
}

func isCorrectForBishop(from, to board.SquareName, g *Game) bool {
	return isDiagonalN(from, to)
}

func isCorrectForQueen(from, to board.SquareName, g *Game) bool {
	return isDiagonalN(from, to) || isVerticalN(from, to) || isHorizontalN(from, to)
}

func isCorrectForKing(from, to board.SquareName, g *Game) bool {
	return isVerticalOne(from, to) || isHorizontalOne(from, to) || isDiagonalOne(from, to)
}

func isCapture(p board.Piece, to board.SquareName, g *Game) bool {
	if isPieceWhite(p) {
		return isPieceBlack(g.Board.GetPiece(to))
	} else if isPieceBlack(p) {
		return isPieceWhite(g.Board.GetPiece(to))
	}

	return false
}

func isDiagonalUpOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (lt-lf == 1) && (math.Abs(float64(ct-cf)) == 1)
}

func isFirstSquareForWhitePawn(s board.SquareName) bool {
	return s == board.A2 ||
		s == board.B2 ||
		s == board.C2 ||
		s == board.D2 ||
		s == board.E2 ||
		s == board.F2 ||
		s == board.G2 ||
		s == board.H2
}

func isVerticalUpOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (lt-lf == 1) && (ct-cf == 0)
}

func isVerticalUpTwo(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (lt-lf == 2) && (ct-cf == 0)
}

func isDiagonalDownOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (lt-lf == -1) && (math.Abs(float64(ct-cf)) == 1)
}

func isFirstSquareForBlackPawn(s board.SquareName) bool {
	return s == board.A7 ||
		s == board.B7 ||
		s == board.C7 ||
		s == board.D7 ||
		s == board.E7 ||
		s == board.F7 ||
		s == board.G7 ||
		s == board.H7
}

func isVerticalDownOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (lt-lf == -1) && (ct-cf == 0)
}

func isVerticalDownTwo(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (lt-lf == -2) && (ct-cf == 0)
}

func isL(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (math.Abs(float64(lt-lf)) == 2 && math.Abs(float64(ct-cf)) == 1) ||
		(math.Abs(float64(lt-lf)) == 1 && math.Abs(float64(ct-cf)) == 2)

}

func isVerticalOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (math.Abs(float64(lt-lf)) == 1 && ct-cf == 0)
}

func isHorizontalOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (math.Abs(float64(ct-cf)) == 1 && lt-lf == 0)
}

func isDiagonalOne(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (math.Abs(float64(lt-lf)) == 1 && math.Abs(float64(ct-cf)) == 1)
}

func isVerticalN(from, to board.SquareName) bool {
	_, cf := board.GetCoordinatesBySquareName(from)
	_, ct := board.GetCoordinatesBySquareName(to)

	return cf == ct
}

func isHorizontalN(from, to board.SquareName) bool {
	lf, _ := board.GetCoordinatesBySquareName(from)
	lt, _ := board.GetCoordinatesBySquareName(to)

	return lf == lt
}

func isDiagonalN(from, to board.SquareName) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return math.Abs(float64(lt-lf)) == math.Abs(float64(ct-cf))
}
