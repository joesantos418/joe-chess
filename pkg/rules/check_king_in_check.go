package rules

import (
	"math"

	"github.com/joesantos418/joe-chess/pkg/board"
)

type direction int64

const (
	NE direction = iota
	NW
	SE
	SW
)

func isKingInCheck(p board.Piece, from board.SquareName, to board.SquareName, g *Game) (bool, error) {
	b := getClonedBoard(g)
	b, err := simulateMove(p, from, to, b)
	if err != nil {
		return false, err
	}

	return checkKingInCheck(g, b), nil
}

func getClonedBoard(g *Game) *board.Board {
	return g.Board.Clone()
}

func simulateMove(p board.Piece, from board.SquareName, to board.SquareName, b *board.Board) (*board.Board, error) {
	err := b.Move(p, from, to)
	if err != nil {
		return &board.Board{}, err
	}

	return b, nil
}

func checkKingInCheck(g *Game, b *board.Board) bool {
	if g.NowPlays == WHITES {
		lin, col := getWhiteKingCoordinates(b)
		return isSquareInCheck(lin, col, b, g.NowPlays)
	} else {
		lin, col := getBlackKingCoordinates(b)
		return isSquareInCheck(lin, col, b, g.NowPlays)
	}
}

func isSquareInCheck(lin, col int, b *board.Board, c Color) bool {
	return isThreatenedVertical(lin, col, b, c) ||
		isThreatenedHorizontal(lin, col, b, c) ||
		isThreatenedDiagonal(lin, col, b, c) ||
		isThreatenedL(lin, col, b, c)
}

// wont write a generic method because some pieces can appear more than once on the board
func getWhiteKingCoordinates(b *board.Board) (int, int) {
	for lin := 0; lin < board.MAX_LIN; lin++ {
		for col := 0; col < board.MAX_COL; col++ {
			if b.Squares[lin][col].Piece == board.WHITE_KING {
				return lin, col
			}
		}
	}

	panic("could not find white king on the board")
}

func getBlackKingCoordinates(b *board.Board) (int, int) {
	for lin := 0; lin < board.MAX_LIN; lin++ {
		for col := 0; col < board.MAX_COL; col++ {
			if b.Squares[lin][col].Piece == board.BLACK_KING {
				return lin, col
			}
		}
	}

	panic("could not find white king on the board")
}

func isThreatenedVertical(lin int, col int, b *board.Board, c Color) bool {
	return isThreatenedUp(lin, col, b, c) || isThreatenedDown(lin, col, b, c)
}

func isThreatenedHorizontal(lin int, col int, b *board.Board, c Color) bool {
	return isThreatenedRight(lin, col, b, c) || isThreatenedLeft(lin, col, b, c)
}

func isThreatenedDiagonal(lin int, col int, b *board.Board, c Color) bool {
	return isThreatenedNE(lin, col, b, c) ||
		isThreatenedNW(lin, col, b, c) ||
		isThreatenedSE(lin, col, b, c) ||
		isThreatenedSW(lin, col, b, c)
}

func isThreatenedL(lin, col int, b *board.Board, c Color) bool {
	return isDangerL(lin+1, col+2, b, c) ||
		isDangerL(lin+2, col+1, b, c) ||
		isDangerL(lin-1, col+2, b, c) ||
		isDangerL(lin-2, col+1, b, c) ||
		isDangerL(lin+1, col-2, b, c) ||
		isDangerL(lin+2, col-1, b, c) ||
		isDangerL(lin-1, col-2, b, c) ||
		isDangerL(lin-2, col-1, b, c)
}

func isDangerL(lin, col int, b *board.Board, c Color) bool {
	if lin < 0 || lin >= board.MAX_COL || col < 0 || col >= board.MAX_COL {
		return false
	}

	return b.Squares[lin][col].Piece != board.NO_PIECE &&
		c != getColor(b.Squares[lin][col].Piece) &&
		isLAttacker(b.Squares[lin][col].Piece)
}

func isLAttacker(p board.Piece) bool {
	return p == board.WHITE_KNIGHT_QUEEN ||
		p == board.WHITE_KNIGHT_KING ||
		p == board.BLACK_KNIGHT_QUEEN ||
		p == board.BLACK_KNIGHT_KING
}

func isThreatenedUp(lin int, col int, b *board.Board, c Color) bool {
	if lin == (board.MAX_LIN - 1) {
		return false
	}

	for i := 1; i < (board.MAX_LIN - lin); i++ {
		if b.Squares[lin+i][col].Piece != board.NO_PIECE {
			return isVerticalDanger(b.Squares[lin+i][col].Piece, lin, lin+i, c)
		}
	}

	return false
}

func isThreatenedDown(lin int, col int, b *board.Board, c Color) bool {
	if lin == 0 {
		return false
	}

	for i := 1; i <= lin; i++ {
		if b.Squares[lin-i][col].Piece != board.NO_PIECE {
			return isVerticalDanger(b.Squares[lin-i][col].Piece, lin, lin-i, c)
		}
	}

	return false
}

func isVerticalDanger(atk board.Piece, from int, to int, c Color) bool {
	return c != getColor(atk) &&
		isAttackVertical(atk) &&
		isInRangeForVerticalAttack(atk, from, to)
}

func isAttackVertical(p board.Piece) bool {
	return p == board.WHITE_ROOK_QUEEN ||
		p == board.WHITE_ROOK_KING ||
		p == board.WHITE_QUEEN ||
		p == board.WHITE_KING ||
		p == board.BLACK_ROOK_QUEEN ||
		p == board.BLACK_ROOK_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}

func isInRangeForVerticalAttack(p board.Piece, from, to int) bool {
	switch p {
	case board.WHITE_ROOK_QUEEN,
		board.WHITE_ROOK_KING,
		board.WHITE_QUEEN,
		board.BLACK_ROOK_QUEEN,
		board.BLACK_ROOK_KING,
		board.BLACK_QUEEN:
		return true
	case board.WHITE_KING,
		board.BLACK_KING:
		return math.Abs(float64(to-from)) == 1
	default:
		return false
	}
}

func isThreatenedRight(lin int, col int, b *board.Board, c Color) bool {
	if col == (board.MAX_COL - 1) {
		return false
	}

	for i := 1; i < (board.MAX_COL - col); i++ {
		if b.Squares[lin][col+i].Piece != board.NO_PIECE {
			return isHorizontalDanger(b.Squares[lin][col+i].Piece, col, col+i, c)
		}
	}

	return false
}

// isThreatenedLeft checks if a piece in a square with coordinates @lin, @col
// in a board @b has an horizontal threat to the left while being @c colors turn
func isThreatenedLeft(lin int, col int, b *board.Board, c Color) bool {
	if col == 0 {
		return false
	}

	for i := 1; i <= col; i++ {
		if b.Squares[lin][col-i].Piece != board.NO_PIECE {
			return isHorizontalDanger(b.Squares[lin][col-i].Piece, col, col-i, c)
		}
	}

	return false
}

func isHorizontalDanger(atk board.Piece, from int, to int, c Color) bool {
	return c != getColor(atk) &&
		isAttackHorizontal(atk) &&
		isInRangeForHorizontalAttack(atk, from, to)
}

func isAttackHorizontal(p board.Piece) bool {
	return p == board.WHITE_ROOK_QUEEN ||
		p == board.WHITE_ROOK_KING ||
		p == board.WHITE_QUEEN ||
		p == board.WHITE_KING ||
		p == board.BLACK_ROOK_QUEEN ||
		p == board.BLACK_ROOK_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}

func isInRangeForHorizontalAttack(p board.Piece, from, to int) bool {
	switch p {
	case board.WHITE_ROOK_QUEEN,
		board.WHITE_ROOK_KING,
		board.WHITE_QUEEN,
		board.BLACK_ROOK_QUEEN,
		board.BLACK_ROOK_KING,
		board.BLACK_QUEEN:
		return true
	case board.WHITE_KING,
		board.BLACK_KING:
		return math.Abs(float64(to-from)) == 1
	default:
		return false
	}
}

func isSquareEmpty(s board.Square) bool {
	return s.Piece == board.NO_PIECE
}

// isThreatenedNE checks if a piece in a square with coordinates @lin, @col
// in a board @b is threatened diagonally from the NE direction during @c colors
// turn
func isThreatenedNE(lin int, col int, b *board.Board, c Color) bool {
	if lin == (board.MAX_LIN-1) || col == (board.MAX_COL-1) {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, NE); i++ {
		// exercise: why not just use lin + i and col - i?
		// hint: mind beacons and avoiding errors
		nextL := lin + i
		nextC := col + i

		// exercise: this condition is repeated for all 4 methods
		// can we encapsulate it in a function?
		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(atk, lin, col, nextL, nextC, c)
		}
	}

	return false
}

func isThreatenedNW(lin int, col int, b *board.Board, c Color) bool {
	if lin == board.MAX_LIN || col == 0 {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, NW); i++ {
		nextL := lin + i
		nextC := col - i

		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(atk, lin, col, nextL, nextC, c)
		}
	}

	return false
}

func isThreatenedSE(lin int, col int, b *board.Board, c Color) bool {
	if lin == 0 || col == board.MAX_COL {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, SE); i++ {
		nextL := lin - i
		nextC := col + i

		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(atk, lin, col, nextL, nextC, c)
		}
	}

	return false
}

func isThreatenedSW(lin int, col int, b *board.Board, c Color) bool {
	if lin == 0 || col == 0 {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, SW); i++ {
		nextL := lin - i
		nextC := col - i

		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(atk, lin, col, nextL, nextC, c)
		}
	}

	return false
}

func isDiagonalDanger(atk board.Piece, lf int, cf int, lt int, ct int, c Color) bool {
	return c != getColor(atk) &&
		isAttackDiagonal(atk) &&
		isInRangeForDiagonalAttack(atk, lf, cf, lt, ct)
}

// is this a good name? maybe is diagonalAttacker
func isAttackDiagonal(p board.Piece) bool {
	return p == board.WHITE_PAWN_A ||
		p == board.WHITE_PAWN_B ||
		p == board.WHITE_PAWN_C ||
		p == board.WHITE_PAWN_D ||
		p == board.WHITE_PAWN_E ||
		p == board.WHITE_PAWN_F ||
		p == board.WHITE_PAWN_G ||
		p == board.WHITE_PAWN_H ||
		p == board.WHITE_BISHOP_QUEEN ||
		p == board.WHITE_BISHOP_KING ||
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
		p == board.BLACK_BISHOP_QUEEN ||
		p == board.BLACK_BISHOP_KING ||
		p == board.BLACK_QUEEN ||
		p == board.BLACK_KING
}

func isInRangeForDiagonalAttack(atk board.Piece, lf int, cf int, lt int, ct int) bool {
	switch atk {
	case board.WHITE_PAWN_A,
		board.WHITE_PAWN_B,
		board.WHITE_PAWN_C,
		board.WHITE_PAWN_D,
		board.WHITE_PAWN_E,
		board.WHITE_PAWN_F,
		board.WHITE_PAWN_G,
		board.WHITE_PAWN_H:
		return lt-lf == 1 && math.Abs(float64(ct-cf)) == 1
	case board.WHITE_BISHOP_QUEEN,
		board.WHITE_BISHOP_KING,
		board.WHITE_QUEEN:
		return true
	case board.WHITE_KING:
		return math.Abs(float64(lt-lf)) == 1 && math.Abs(float64(ct-cf)) == 1

	case board.BLACK_PAWN_A,
		board.BLACK_PAWN_B,
		board.BLACK_PAWN_C,
		board.BLACK_PAWN_D,
		board.BLACK_PAWN_E,
		board.BLACK_PAWN_F,
		board.BLACK_PAWN_G,
		board.BLACK_PAWN_H:
		return lt-lf == -1 && math.Abs(float64(ct-cf)) == 1
	case board.BLACK_BISHOP_QUEEN,
		board.BLACK_BISHOP_KING,
		board.BLACK_QUEEN:
		return true
	case board.BLACK_KING:
		return math.Abs(float64(lt-lf)) == 1 && math.Abs(float64(ct-cf)) == 1
	default:
		return false
	}
}

func getMaxSteps(lin, col int, d direction) int {
	var l, c int

	switch d {
	case NE:
		l = (board.MAX_LIN - 1) - lin
		c = (board.MAX_COL - 1) - col
	case NW:
		l = (board.MAX_LIN - 1) - lin
		c = col
	case SE:
		l = lin
		c = (board.MAX_COL - 1) - col
	case SW:
		l = lin
		c = col
	}

	if l > c {
		return c
	}

	return l
}

func getColor(p board.Piece) Color {
	if isPieceWhite(p) {
		return WHITES
	}

	if isPieceBlack(p) {
		return BLACKS
	}

	panic("Cannot get color of piece")
}
