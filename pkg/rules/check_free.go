package rules

import (
	"math"

	"github.com/joesantos418/joe-chess/pkg/board"
)

func isMoveFree(p board.Piece, from, to board.SquareName, g *Game) bool {
	if g.Board.GetPiece(to) != board.NO_PIECE {
		if !isCapture(p, to, g) {
			return false
		}
	}

	switch p {
	case board.WHITE_PAWN_A,
		board.WHITE_PAWN_B,
		board.WHITE_PAWN_C,
		board.WHITE_PAWN_D,
		board.WHITE_PAWN_E,
		board.WHITE_PAWN_F,
		board.WHITE_PAWN_G,
		board.WHITE_PAWN_H,
		board.BLACK_PAWN_A,
		board.BLACK_PAWN_B,
		board.BLACK_PAWN_C,
		board.BLACK_PAWN_D,
		board.BLACK_PAWN_E,
		board.BLACK_PAWN_F,
		board.BLACK_PAWN_G,
		board.BLACK_PAWN_H:
		return isFreeForPawn(p, from, to, g)

	case board.WHITE_ROOK_QUEEN,
		board.WHITE_ROOK_KING,
		board.BLACK_ROOK_QUEEN,
		board.BLACK_ROOK_KING:
		return isFreeForRook(p, from, to, g)

	case board.BLACK_KNIGHT_QUEEN,
		board.BLACK_KNIGHT_KING,
		board.WHITE_KNIGHT_QUEEN,
		board.WHITE_KNIGHT_KING:
		return isFreeForKnight(p, from, to, g)

	case board.BLACK_BISHOP_QUEEN,
		board.BLACK_BISHOP_KING,
		board.WHITE_BISHOP_QUEEN,
		board.WHITE_BISHOP_KING:
		return isFreeForBishop(p, from, to, g)

	case board.BLACK_QUEEN,
		board.WHITE_QUEEN:
		return isFreeForQueen(p, from, to, g)

	case board.BLACK_KING,
		board.WHITE_KING:
		return isFreeForKing(p, from, to, g)

	}

	return false
}

func isEmpty(n board.SquareName, g *Game) bool {
	return g.Board.GetPiece(n) == board.NO_PIECE
}

func canOcuppy(p board.Piece, to board.SquareName, g *Game) bool {
	return isEmpty(to, g) || isCapture(p, to, g)
}

func isFreeForPawn(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return (isVertical(lf, cf, lt, ct) && isFreeVertical(cf, lf, lt, g) && isEmpty(to, g)) ||
		(isDiagonal(lf, cf, lt, ct) && isFreeDiagonal(lf, cf, lt, ct, g) && isCapture(p, to, g))
}

func isFreeForRook(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return canOcuppy(p, to, g) &&
		((isVertical(lf, cf, lt, ct) && isFreeVertical(cf, lf, lt, g)) ||
			(isHorizontal(lf, cf, lt, ct) && isFreeHorizontal(lf, cf, ct, g)))
}

func isFreeForKnight(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	return canOcuppy(p, to, g)
}

func isFreeForBishop(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return canOcuppy(p, to, g) && (isDiagonal(lf, cf, lt, ct) && isFreeDiagonal(lf, cf, lt, ct, g))
}

func isFreeForQueen(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return canOcuppy(p, to, g) &&
		((isVertical(lf, cf, lt, ct) && isFreeVertical(cf, lf, lt, g)) ||
			(isHorizontal(lf, cf, lt, ct) && isFreeHorizontal(lf, cf, ct, g)) ||
			(isDiagonal(lf, cf, lt, ct) && isFreeDiagonal(lf, cf, lt, ct, g)))
}

func isFreeForKing(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	lf, cf := board.GetCoordinatesBySquareName(from)
	lt, ct := board.GetCoordinatesBySquareName(to)

	return canOcuppy(p, to, g) &&
		((isVertical(lf, cf, lt, ct) && isFreeVertical(cf, lf, lt, g)) ||
			(isHorizontal(lf, cf, lt, ct) && isFreeHorizontal(lf, cf, ct, g)) ||
			(isDiagonal(lf, cf, lt, ct) && isFreeDiagonal(lf, cf, lt, ct, g)))
}

func isVertical(lf, cf, lt, ct int) bool {
	return cf == ct && lf != lt
}

func isHorizontal(lf, cf, lt, ct int) bool {
	return cf != ct && lf == lt
}

func isDiagonal(lf, cf, lt, ct int) bool {
	return math.Abs(float64(ct-cf)) == math.Abs(float64(lt-lf))
}

func isFreeVertical(col, from, to int, g *Game) bool {
	f := from
	t := to
	if from > to {
		f = to
		t = from
	}

	for i := f + 1; i < t; i++ {
		if g.Board.Squares[i][col].Piece != board.NO_PIECE {
			return false
		}
	}

	return true
}

func isFreeHorizontal(lin, from, to int, g *Game) bool {
	f := from
	t := to
	if from > to {
		f = to
		t = from
	}

	for i := f + 1; i < t; i++ {
		if g.Board.Squares[lin][i].Piece != board.NO_PIECE {
			return false
		}
	}

	return true
}

func isFreeDiagonal(lf, cf, lt, ct int, g *Game) bool {
	if lt > lf {
		if ct > cf {
			return isFreeDiagonalNE(lf, cf, lt, g)
		} else {
			return isFreeDiagonalNW(lf, cf, lt, g)
		}
	} else {
		if ct > cf {
			return isFreeDiagonalSE(lf, cf, lt, g)
		} else {
			return isFreeDiagonalSW(lf, cf, lt, g)
		}
	}
}

func isFreeDiagonalNE(lf, cf, lt int, g *Game) bool {
	steps := lt - lf
	for i := 1; i < steps; i++ {
		if g.Board.Squares[lf+i][cf+i].Piece != board.NO_PIECE {
			return false
		}
	}

	return true
}

func isFreeDiagonalNW(lf, cf, lt int, g *Game) bool {
	steps := lt - lf
	for i := 1; i < steps; i++ {
		if g.Board.Squares[lf+i][cf-i].Piece != board.NO_PIECE {
			return false
		}
	}

	return true
}

func isFreeDiagonalSE(lf, cf, lt int, g *Game) bool {
	steps := lf - lt
	for i := 1; i < steps; i++ {
		if g.Board.Squares[lf-i][cf+i].Piece != board.NO_PIECE {
			return false
		}
	}

	return true
}

func isFreeDiagonalSW(lf, cf, lt int, g *Game) bool {
	steps := lf - lt
	for i := 1; i < steps; i++ {
		if g.Board.Squares[lf-i][cf-i].Piece != board.NO_PIECE {
			return false
		}
	}

	return true
}
