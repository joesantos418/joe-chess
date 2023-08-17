package rules

import (
	"testing"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestIsCastle_NoKing(t *testing.T) {
	is := isCastle(board.BLACK_BISHOP_KING, board.G4)
	assert.False(t, is)
}

func TestIsCastle_NoCastle(t *testing.T) {
	is := isCastle(board.WHITE_KING, board.E2)
	assert.False(t, is)
}

func TestIsCastle_Castle(t *testing.T) {
	is := isCastle(board.WHITE_KING, board.C1)
	assert.True(t, is)
}

func TestIsCastle_CastleBlack(t *testing.T) {
	is := isCastle(board.BLACK_KING, board.C8)
	assert.True(t, is)
}

func TestIsKing_Pawn(t *testing.T) {
	is := isKing(board.WHITE_PAWN_A)
	assert.False(t, is)
}

func TestIsKing_King(t *testing.T) {
	is := isKing(board.WHITE_KING)
	assert.True(t, is)
}

func TestHasKingMoved_NewGame(t *testing.T) {
	g := NewGame()
	has := hasKingMoved(board.WHITE_KING, g)

	assert.False(t, has)
}

func TestHasKingMoved_NotKing(t *testing.T) {
	g := NewGame()
	has := hasKingMoved(board.WHITE_PAWN_A, g)

	assert.False(t, has)
}

func TestHasKingMoved_No(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	assert.Nil(t, err)

	has := hasKingMoved(board.WHITE_KING, g)
	assert.False(t, has)
}

func TestHasKingMoved_Black(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	assert.Nil(t, err)

	has := hasKingMoved(board.BLACK_KING, g)
	assert.False(t, has)
}

func TestHasKingMoved_Yes(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
	assert.Nil(t, err)

	err = g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
	assert.Nil(t, err)

	err = g.ProcessMove(board.WHITE_KING, board.E1, board.E2)
	assert.Nil(t, err)

	has := hasKingMoved(board.WHITE_KING, g)
	assert.True(t, has)
}

func TestHasRookMovedNewGame(t *testing.T) {
	g := NewGame()
	has := hasRookMoved(board.E4, g)

	assert.False(t, has)
}

func TestHasRookMovedBlackKing(t *testing.T) {
	g := NewGame()
	has := hasRookMoved(board.G8, g)

	assert.False(t, has)
}

func TestHasRookMovedBlackQueen(t *testing.T) {
	g := NewGame()
	has := hasRookMoved(board.C8, g)

	assert.False(t, has)
}

func TestHasRookMoved_NoCastle(t *testing.T) {
	g := NewGame()

	err := g.ProcessMove(board.WHITE_PAWN_A, board.A2, board.A4)
	assert.Nil(t, err)

	err = g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
	assert.Nil(t, err)

	has := hasRookMoved(board.A2, g)
	assert.False(t, has)
}

func TestHasRookMoved_CastleMoved(t *testing.T) {
	g := NewGame()

	g.Board.SetPiece(board.NO_PIECE, board.B1)
	g.Board.SetPiece(board.NO_PIECE, board.C1)
	g.Board.SetPiece(board.NO_PIECE, board.D1)
	g.Board.SetPiece(board.NO_PIECE, board.A2)

	err := g.ProcessMove(board.WHITE_ROOK_QUEEN, board.A1, board.A4)
	assert.Nil(t, err)

	has := hasRookMoved(board.C1, g)
	assert.True(t, has)
}

func TestHasRookMoved_Castle(t *testing.T) {
	g := NewGame()

	g.Board.SetPiece(board.NO_PIECE, board.B1)
	g.Board.SetPiece(board.NO_PIECE, board.C1)
	g.Board.SetPiece(board.NO_PIECE, board.D1)
	g.Board.SetPiece(board.NO_PIECE, board.A2)

	has := hasRookMoved(board.C1, g)
	assert.False(t, has)
}

func TestIsPathFree_NoCastle(t *testing.T) {
	g := NewGame()
	p := board.WHITE_PAWN_A
	from := board.A2
	to := board.A4
	is := isPathFree(p, from, to, g)

	assert.False(t, is)
}

func TestIsPathFree_ShortWhite(t *testing.T) {
	g := NewGame()
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	is := isPathFree(p, from, to, g)

	assert.False(t, is)
}

func TestIsPathFree_LongWhite(t *testing.T) {
	g := NewGame()
	p := board.WHITE_KING
	from := board.E1
	to := board.C1
	is := isPathFree(p, from, to, g)

	assert.False(t, is)
}

func TestIsPathFree_ShortBlack(t *testing.T) {
	g := NewGame()
	p := board.BLACK_KING
	from := board.E8
	to := board.G8
	is := isPathFree(p, from, to, g)

	assert.False(t, is)
}

func TestIsPathFree_LongBlack(t *testing.T) {
	g := NewGame()
	p := board.BLACK_KING
	from := board.E8
	to := board.C8
	is := isPathFree(p, from, to, g)

	assert.False(t, is)
}

func TestIsPathFreeForShortWhiteCastle_No(t *testing.T) {
	g := NewGame()
	is := isPathFreeForShortWhiteCastle(g)

	assert.False(t, is)
}

func TestIsPathFreeForShortWhiteCastle_Yes(t *testing.T) {
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	is := isPathFreeForShortWhiteCastle(g)

	assert.True(t, is)
}

func TestIsPathFreeForLongWhiteCastle_No(t *testing.T) {
	g := NewGame()
	is := isPathFreeForLongWhiteCastle(g)

	assert.False(t, is)
}

func TestIsPathFreeForLongWhiteCastle_Yes(t *testing.T) {
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.D1)
	is := isPathFreeForLongWhiteCastle(g)

	assert.True(t, is)
}

func TestIsPathFreeForShortBlackCastle_No(t *testing.T) {
	g := NewGame()
	is := isPathFreeForShortBlackCastle(g)

	assert.False(t, is)
}

func TestIsPathFreeForShortBlackCastle_Yes(t *testing.T) {
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F8)
	g.NowPlays = BLACKS
	is := isPathFreeForShortBlackCastle(g)

	assert.True(t, is)
}

func TestIsPathFreeForLongBlackCastle_No(t *testing.T) {
	g := NewGame()
	is := isPathFreeForLongBlackCastle(g)

	assert.False(t, is)
}

func TestIsPathFreeForLongBlackCastle_Yes(t *testing.T) {
	g := NewGame()
	g.NowPlays = BLACKS
	g.Board.SetPiece(board.NO_PIECE, board.D8)
	is := isPathFreeForLongBlackCastle(g)

	assert.True(t, is)
}

func TestCanCastle_NotKing(t *testing.T) {
	p := board.WHITE_PAWN_A
	from := board.A2
	to := board.A4
	g := NewGame()
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.False(t, can)
}

func TestCanCastle_KingHasMoved(t *testing.T) {
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	g.Board.SetPiece(board.NO_PIECE, board.G1)
	g.HasWhiteKingMoved = true
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.False(t, can)
}

func TestCanCastle_RookKingHasMoved(t *testing.T) {
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	g.Board.SetPiece(board.NO_PIECE, board.G1)
	g.HasWhiteRookKingMoved = true
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.False(t, can)
}

func TestCanCastle_RookQueenHasMoved(t *testing.T) {
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	g.Board.SetPiece(board.NO_PIECE, board.G1)
	g.HasWhiteRookQueenMoved = true
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.True(t, can)
}

func TestCanCastle_KingInCheck(t *testing.T) {
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	g.Board.SetPiece(board.NO_PIECE, board.G1)
	g.Board.SetPiece(board.BLACK_QUEEN, board.E2)
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.False(t, can)
}

func TestCanCastle_PathNotFree(t *testing.T) {
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	g.Board.SetPiece(board.NO_PIECE, board.G1)
	g.Board.SetPiece(board.NO_PIECE, board.E2)
	g.Board.SetPiece(board.BLACK_QUEEN, board.C4)
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.False(t, can)
}

func TestCanCastle_Can(t *testing.T) {
	p := board.WHITE_KING
	from := board.E1
	to := board.G1
	g := NewGame()
	g.Board.SetPiece(board.NO_PIECE, board.F1)
	g.Board.SetPiece(board.NO_PIECE, board.G1)
	can, err := canCastle(p, from, to, g)

	assert.Nil(t, err)
	assert.True(t, can)
}
