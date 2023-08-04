package utils

import (
	"fmt"
	"regexp"

	"github.com/arquivei/foundationkit/errors"
	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/joesantos418/joe-chess/pkg/rules"
)

var g *rules.Game

func ParseCommand(cmd string) error {
	switch cmd {
	case "start":
		g = rules.NewGame()
		PrintGame(g)
	case "end":
		PrintGame(g)
	default:
		err := parseMove(cmd)
		if err != nil {
			return err
		}
		PrintGame(g)
	}

	return nil
}

func parseMove(m string) error {
	r := regexp.MustCompile("^[a-h][1-8][a-h][1-8]$")
	ok := r.MatchString(m)
	if !ok {
		return errors.E("invalid move", errors.KV("move", m))
	}

	from := getSquareName(m[0:2])
	to := getSquareName(m[2:4])

	return g.ProcessMove(g.Board.GetPiece(from), from, to)
}

func getSquareName(s string) board.SquareName {
	switch s {
	case "a1":
		return board.A1
	case "a2":
		return board.A2
	case "a3":
		return board.A3
	case "a4":
		return board.A4
	case "a5":
		return board.A5
	case "a6":
		return board.A6
	case "a7":
		return board.A7
	case "a8":
		return board.A8
	case "b1":
		return board.B1
	case "b2":
		return board.B2
	case "b3":
		return board.B3
	case "b4":
		return board.B4
	case "b5":
		return board.B5
	case "b6":
		return board.B6
	case "b7":
		return board.B7
	case "b8":
		return board.B8
	case "c1":
		return board.C1
	case "c2":
		return board.C2
	case "c3":
		return board.C3
	case "c4":
		return board.C4
	case "c5":
		return board.C5
	case "c6":
		return board.C6
	case "c7":
		return board.C7
	case "c8":
		return board.C8
	case "d1":
		return board.D1
	case "d2":
		return board.D2
	case "d3":
		return board.D3
	case "d4":
		return board.D4
	case "d5":
		return board.D5
	case "d6":
		return board.D6
	case "d7":
		return board.D7
	case "d8":
		return board.D8
	case "e1":
		return board.E1
	case "e2":
		return board.E2
	case "e3":
		return board.E3
	case "e4":
		return board.E4
	case "e5":
		return board.E5
	case "e6":
		return board.E6
	case "e7":
		return board.E7
	case "e8":
		return board.E8
	case "f1":
		return board.F1
	case "f2":
		return board.F2
	case "f3":
		return board.F3
	case "f4":
		return board.F4
	case "f5":
		return board.F5
	case "f6":
		return board.F6
	case "f7":
		return board.F7
	case "f8":
		return board.F8
	case "g1":
		return board.G1
	case "g2":
		return board.G2
	case "g3":
		return board.G3
	case "g4":
		return board.G4
	case "g5":
		return board.G5
	case "g6":
		return board.G6
	case "g7":
		return board.G7
	case "g8":
		return board.G8
	case "h1":
		return board.H1
	case "h2":
		return board.H2
	case "h3":
		return board.H3
	case "h4":
		return board.H4
	case "h5":
		return board.H5
	case "h6":
		return board.H6
	case "h7":
		return board.H7
	case "h8":
		return board.H8
	}

	panic(fmt.Sprintf("Could not parse square name: %s~~~", s))
}
