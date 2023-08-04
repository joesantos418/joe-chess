package main

import (
	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/joesantos418/joe-chess/pkg/rules"
	"github.com/joesantos418/joe-chess/pkg/utils"
)

func main() {
	// for {
	// 	reader := bufio.NewReader(os.Stdin)

	// 	text, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Default().Print(err)
	// 	}

	// 	text = strings.Trim(text, "\n")

	// 	err = utils.ParseCommand(text)
	// 	if err != nil {
	// 		log.Default().Print(err)
	// 	}
	// }

	g := rules.NewGame()

	move(g, board.E2, board.E4)
	move(g, board.E7, board.E5)
	move(g, board.G1, board.F3)
	move(g, board.B8, board.C6)
	move(g, board.F1, board.B5)
	move(g, board.D7, board.D6)
	move(g, board.B5, board.C6)
	move(g, board.B7, board.C6)
	move(g, board.E4, board.E5)

	utils.PrintGame(g)
}

func move(g *rules.Game, from, to board.SquareName) {
	err := g.ProcessMove(g.Board.GetPiece(from), from, to)
	if err != nil {
		panic(err)
	}
}
