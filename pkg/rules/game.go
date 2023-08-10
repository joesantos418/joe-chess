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
	Board                  *board.Board
	WhiteCaptured          []board.Piece
	BlackCaptured          []board.Piece
	NowPlays               Color
	HasWhiteKingMoved      bool `default:"false"`
	HasWhiteRookKingMoved  bool `default:"false"`
	HasWhiteRookQueenMoved bool `default:"false"`
	HasBlackKingMoved      bool `default:"false"`
	HasBlackRookKingMoved  bool `default:"false"`
	HasBlackRookQueenMoved bool `default:"false"`
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
