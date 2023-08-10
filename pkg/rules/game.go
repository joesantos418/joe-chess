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
	Board              *board.Board
	WhiteCaptured      []board.Piece
	BlackCaptured      []board.Piece
	NowPlays           Color
	IsWhiteKingInCheck bool
	IsBlackKingInCheck bool
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
