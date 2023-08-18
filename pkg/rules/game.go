package rules

import (
	"github.com/arquivei/foundationkit/errors"
	"github.com/joesantos418/joe-chess/pkg/board"
)

type Color int64

const (
	WHITES Color = iota
	BLACKS
)

func (c Color) String() string {
	if c == WHITES {
		return "WHITES"
	}

	return "BLACKS"
}

type Game struct {
	Board         *board.Board
	WhiteCaptured []board.Piece
	BlackCaptured []board.Piece
	NowPlays      Color

	HasWhiteKingMoved      bool `default:"false"`
	HasWhiteRookKingMoved  bool `default:"false"`
	HasWhiteRookQueenMoved bool `default:"false"`

	HasBlackKingMoved      bool `default:"false"`
	HasBlackRookKingMoved  bool `default:"false"`
	HasBlackRookQueenMoved bool `default:"false"`

	HasWhitePawnAJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnBJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnCJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnDJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnEJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnFJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnGJustMovedTwoSquares bool `default:"false"`
	HasWhitePawnHJustMovedTwoSquares bool `default:"false"`

	HasBlackPawnAJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnBJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnCJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnDJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnEJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnFJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnGJustMovedTwoSquares bool `default:"false"`
	HasBlackPawnHJustMovedTwoSquares bool `default:"false"`
}

func NewGame() *Game {
	return &Game{
		Board:    setupBoard(),
		NowPlays: WHITES,
	}
}

func (g *Game) ProcessMove(p board.Piece, from, to board.SquareName) error {
	ok, reason, err := isValidMove(p, from, to, g)
	if err != nil {
		return err
	}

	if !ok {
		return errors.E(
			"invalid move",
			errors.KV("piece", p),
			errors.KV("from", from),
			errors.KV("to", to),
			errors.KV("reason", reason),
		)
	}

	processCapture(p, to, g)
	err = g.Board.Move(p, from, to)
	if err != nil {
		return err
	}

	switchColorTurn(g)
	updateCastlingData(g, p)
	updateEnPassantData(p, from, to, g)

	return nil
}

func processCapture(p board.Piece, to board.SquareName, g *Game) {
	if !isCapture(p, to, g) {
		return
	}

	t := g.Board.GetPiece(to)
	if isPieceWhite(t) {
		g.BlackCaptured = append(g.BlackCaptured, t)
	} else if isPieceBlack(t) {
		g.WhiteCaptured = append(g.WhiteCaptured, t)
	}
}

func switchColorTurn(g *Game) {
	if g.NowPlays == WHITES {
		g.NowPlays = BLACKS
	} else {
		g.NowPlays = WHITES
	}
}

func updateCastlingData(g *Game, p board.Piece) {
	switch p {
	case board.WHITE_KING:
		g.HasWhiteKingMoved = true
	case board.BLACK_KING:
		g.HasBlackKingMoved = true
	case board.WHITE_ROOK_KING:
		g.HasWhiteRookKingMoved = true
	case board.WHITE_ROOK_QUEEN:
		g.HasWhiteRookQueenMoved = true
	case board.BLACK_ROOK_KING:
		g.HasBlackRookKingMoved = true
	case board.BLACK_ROOK_QUEEN:
		g.HasBlackRookQueenMoved = true
	}
}

func updateEnPassantData(p board.Piece, from, to board.SquareName, g *Game) {
	switch p {
	case board.WHITE_PAWN_A:
		g.HasWhitePawnAJustMovedTwoSquares = (from == board.A2 && to == board.A4)
	case board.WHITE_PAWN_B:
		g.HasWhitePawnBJustMovedTwoSquares = (from == board.B2 && to == board.B4)
	case board.WHITE_PAWN_C:
		g.HasWhitePawnCJustMovedTwoSquares = (from == board.C2 && to == board.C4)
	case board.WHITE_PAWN_D:
		g.HasWhitePawnDJustMovedTwoSquares = (from == board.D2 && to == board.D4)
	case board.WHITE_PAWN_E:
		g.HasWhitePawnEJustMovedTwoSquares = (from == board.E2 && to == board.E4)
	case board.WHITE_PAWN_F:
		g.HasWhitePawnFJustMovedTwoSquares = (from == board.F2 && to == board.F4)
	case board.WHITE_PAWN_G:
		g.HasWhitePawnGJustMovedTwoSquares = (from == board.G2 && to == board.G4)
	case board.WHITE_PAWN_H:
		g.HasWhitePawnHJustMovedTwoSquares = (from == board.H2 && to == board.H4)
	case board.BLACK_PAWN_A:
		g.HasBlackPawnAJustMovedTwoSquares = (from == board.A7 && to == board.A5)
	case board.BLACK_PAWN_B:
		g.HasBlackPawnBJustMovedTwoSquares = (from == board.B7 && to == board.B5)
	case board.BLACK_PAWN_C:
		g.HasBlackPawnCJustMovedTwoSquares = (from == board.C7 && to == board.C5)
	case board.BLACK_PAWN_D:
		g.HasBlackPawnDJustMovedTwoSquares = (from == board.D7 && to == board.D5)
	case board.BLACK_PAWN_E:
		g.HasBlackPawnEJustMovedTwoSquares = (from == board.E7 && to == board.E5)
	case board.BLACK_PAWN_F:
		g.HasBlackPawnFJustMovedTwoSquares = (from == board.F7 && to == board.F5)
	case board.BLACK_PAWN_G:
		g.HasBlackPawnGJustMovedTwoSquares = (from == board.G7 && to == board.G5)
	case board.BLACK_PAWN_H:
		g.HasBlackPawnHJustMovedTwoSquares = (from == board.H7 && to == board.H5)
	}
}
