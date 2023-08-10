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

func isKingInCheck(p board.Piece, from board.SquareName, to board.SquareName, g *Game) bool {
	b := getClonedBoard(g)
	b = simulateMove(p, from, to, b)
	return checkKingInCheck(g, b)
}

func getClonedBoard(g *Game) *board.Board {
	return g.Board.Clone()
}

func simulateMove(p board.Piece, from board.SquareName, to board.SquareName, b *board.Board) *board.Board {
	err := b.Move(p, from, to)
	if err != nil {
		return &board.Board{}
	}

	return b
}

func checkKingInCheck(g *Game, b *board.Board) bool {
	if g.NowPlays == WHITES {
		return isPieceInCheck(board.WHITE_KING, b)
	} else {
		return isPieceInCheck(board.BLACK_KING, b)
	}
}

func isPieceInCheck(p board.Piece, b *board.Board) bool {
	lin, col := getCoordinates(p, b)
	return isThreatenedVertical(p, lin, col, b) ||
		isThreatenedHorizontal(p, lin, col, b) ||
		isThreatenedDiagonal(p, lin, col, b) ||
		isThreatenedL(p, lin, col, b)
}

func getCoordinates(p board.Piece, b *board.Board) (int, int) {
	for lin := 0; lin < board.MAX_LIN; lin++ {
		for col := 0; col < board.MAX_COL; col++ {
			if b.Squares[lin][col].Piece == p {
				return lin, col
			}
		}
	}

	panic("could not find piece on the board")
}

func isThreatenedVertical(p board.Piece, lin int, col int, b *board.Board) bool {
	return isThreatenedUp(p, lin, col, b) || isThreatenedDown(p, lin, col, b)
}

func isThreatenedHorizontal(p board.Piece, lin int, col int, b *board.Board) bool {
	return isThreatenedRight(p, lin, col, b) || isThreatenedLeft(p, lin, col, b)
}

func isThreatenedDiagonal(p board.Piece, lin int, col int, b *board.Board) bool {
	return isThreatenedNE(p, lin, col, b) ||
		isThreatenedNW(p, lin, col, b) ||
		isThreatenedSE(p, lin, col, b) ||
		isThreatenedSW(p, lin, col, b)
}

func isThreatenedL(p board.Piece, lin, col int, b *board.Board) bool {
	return isDangerL(p, lin+1, col+2, b) ||
		isDangerL(p, lin+2, col+1, b) ||
		isDangerL(p, lin-1, col+2, b) ||
		isDangerL(p, lin-2, col+1, b) ||
		isDangerL(p, lin+1, col-2, b) ||
		isDangerL(p, lin+2, col-1, b) ||
		isDangerL(p, lin-1, col-2, b) ||
		isDangerL(p, lin-2, col-1, b)
}

func isDangerL(p board.Piece, lin, col int, b *board.Board) bool {
	if lin < 0 || lin > board.MAX_COL || col < 0 || col > board.MAX_COL {
		return false
	}

	return b.Squares[lin][col].Piece != board.NO_PIECE &&
		getColor(p) != getColor(b.Squares[lin][col].Piece) &&
		isLAttacker(b.Squares[lin][col].Piece)
}

func isLAttacker(p board.Piece) bool {
	return p == board.WHITE_KNIGHT_QUEEN ||
		p == board.WHITE_KNIGHT_KING ||
		p == board.BLACK_KNIGHT_QUEEN ||
		p == board.BLACK_KNIGHT_KING
}

func isThreatenedUp(p board.Piece, lin int, col int, b *board.Board) bool {
	if lin == board.MAX_LIN {
		return false
	}

	for i := 1; i < (board.MAX_LIN - lin); i++ {
		if b.Squares[lin+i][col].Piece != board.NO_PIECE {
			return isVerticalDanger(p, b.Squares[lin+i][col].Piece, lin, lin+i)
		}
	}

	return false
}

func isThreatenedDown(p board.Piece, lin int, col int, b *board.Board) bool {
	if lin == 0 {
		return false
	}

	for i := 1; i <= lin; i++ {
		if b.Squares[lin-i][col].Piece != board.NO_PIECE {
			return isVerticalDanger(p, b.Squares[lin-i][col].Piece, lin, lin-i)
		}
	}

	return false
}

func isVerticalDanger(def, atk board.Piece, from int, to int) bool {
	return getColor(def) != getColor(atk) &&
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

func isThreatenedRight(p board.Piece, lin int, col int, b *board.Board) bool {
	if col == (board.MAX_COL - 1) {
		return false
	}

	for i := 1; i < (board.MAX_COL - col); i++ {
		if b.Squares[lin][col+i].Piece != board.NO_PIECE {
			return isHorizontalDanger(p, b.Squares[lin][col+i].Piece, col, col+i)
		}
	}

	return false
}

// isThreatenedLeft checks if a piece @p in a square with coordinates @lin, @col
// in a board @b has an horizontal threat to the left
func isThreatenedLeft(p board.Piece, lin int, col int, b *board.Board) bool {
	if col == 0 {
		return false
	}

	for i := 1; i <= col; i++ {
		if b.Squares[lin][col-i].Piece != board.NO_PIECE {
			return isHorizontalDanger(p, b.Squares[lin][col-i].Piece, col, col-i)
		}
	}

	return false
}

func isHorizontalDanger(def, atk board.Piece, from int, to int) bool {
	return getColor(def) != getColor(atk) &&
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

// isThreatenedNE checks if a piece @p in a square with coordinates @lin, @col
// in a board @b is threatened diagonally from the NE direction
func isThreatenedNE(p board.Piece, lin int, col int, b *board.Board) bool {
	if lin == board.MAX_LIN || col == board.MAX_COL {
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
			return isDiagonalDanger(p, atk, lin, col, nextL, nextC)
		}
	}

	return false
}

func isThreatenedNW(p board.Piece, lin int, col int, b *board.Board) bool {
	if lin == board.MAX_LIN || col == 0 {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, NW); i++ {
		nextL := lin + i
		nextC := col - i

		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(p, atk, lin, col, nextL, nextC)
		}
	}

	return false
}

func isThreatenedSE(p board.Piece, lin int, col int, b *board.Board) bool {
	if lin == 0 || col == board.MAX_COL {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, SE); i++ {
		nextL := lin - i
		nextC := col + i

		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(p, atk, lin, col, nextL, nextC)
		}
	}

	return false
}

func isThreatenedSW(p board.Piece, lin int, col int, b *board.Board) bool {
	if lin == 0 || col == 0 {
		return false
	}

	for i := 1; i <= getMaxSteps(lin, col, SW); i++ {
		nextL := lin - i
		nextC := col - i

		if !isSquareEmpty(b.Squares[nextL][nextC]) {
			atk := b.Squares[nextL][nextC].Piece
			return isDiagonalDanger(p, atk, lin, col, nextL, nextC)
		}
	}

	return false
}

func isDiagonalDanger(def board.Piece, atk board.Piece, lf int, cf int, lt int, ct int) bool {
	return getColor(def) != getColor(atk) &&
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

	panic("trying to get color of an unknown piece")
}
