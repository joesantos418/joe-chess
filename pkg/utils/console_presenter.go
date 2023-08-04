package utils

import (
	"fmt"

	"github.com/joesantos418/joe-chess/pkg/board"
	"github.com/joesantos418/joe-chess/pkg/rules"
)

func PrintBoard(b *board.Board) {
	var white = false
	for lin := board.MAX_LIN - 1; lin >= 0; lin-- {
		fmt.Printf("%d ", lin+1)
		for col := 0; col < board.MAX_COL; col++ {
			printSquare(b.Squares[lin][col], white)
			white = !white
			fmt.Printf(" ")
		}

		white = !white
		fmt.Printf("\n")
	}

	fmt.Println("  a b c d e f g h")
}

func PrintGame(g *rules.Game) {
	PrintBoard(g.Board)

	fmt.Printf("\nWhite captured:")
	if len(g.WhiteCaptured) == 0 {
		fmt.Printf(" none")
	} else {
		for i := 0; i < len(g.WhiteCaptured); i++ {
			fmt.Printf(" ")
			printPiece(g.WhiteCaptured[i], true)
		}
	}
	fmt.Println("")

	fmt.Printf("Black captured:")
	if len(g.BlackCaptured) == 0 {
		fmt.Printf(" none")
	} else {
		for i := 0; i < len(g.BlackCaptured); i++ {
			fmt.Printf(" ")
			printPiece(g.BlackCaptured[i], false)
		}
	}
	fmt.Println("")

	fmt.Printf("Now plays: %s\n", g.NowPlays)
}

func printSquare(s board.Square, white bool) {
	printPiece(s.Piece, white)
}

func printPiece(p board.Piece, white bool) {
	switch p {
	case board.WHITE_PAWN_A:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_B:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_C:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_D:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_E:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_F:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_G:
		fmt.Printf("\u265F")
	case board.WHITE_PAWN_H:
		fmt.Printf("\u265F")
	case board.WHITE_ROOK_QUEEN:
		fmt.Printf("\u265C")
	case board.WHITE_ROOK_KING:
		fmt.Printf("\u265C")
	case board.WHITE_KNIGHT_QUEEN:
		fmt.Printf("\u265E")
	case board.WHITE_KNIGHT_KING:
		fmt.Printf("\u265E")
	case board.WHITE_BISHOP_QUEEN:
		fmt.Printf("\u265D")
	case board.WHITE_BISHOP_KING:
		fmt.Printf("\u265D")
	case board.WHITE_QUEEN:
		fmt.Printf("\u265B")
	case board.WHITE_KING:
		fmt.Printf("\u265A")

	case board.BLACK_PAWN_A:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_B:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_C:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_D:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_E:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_F:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_G:
		fmt.Printf("\u2659")
	case board.BLACK_PAWN_H:
		fmt.Printf("\u2659")
	case board.BLACK_ROOK_QUEEN:
		fmt.Printf("\u2656")
	case board.BLACK_ROOK_KING:
		fmt.Printf("\u2656")
	case board.BLACK_KNIGHT_QUEEN:
		fmt.Printf("\u2658")
	case board.BLACK_KNIGHT_KING:
		fmt.Printf("\u2658")
	case board.BLACK_BISHOP_QUEEN:
		fmt.Printf("\u2657")
	case board.BLACK_BISHOP_KING:
		fmt.Printf("\u2657")
	case board.BLACK_QUEEN:
		fmt.Printf("\u2655")
	case board.BLACK_KING:
		fmt.Printf("\u2654")

	default:
		if white {
			fmt.Printf("\u25A1")
		} else {
			fmt.Printf("\u25A0")
		}
	}
}
