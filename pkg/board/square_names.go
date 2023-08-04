package board

type SquareName int64

const (
	A1 SquareName = iota
	A2
	A3
	A4
	A5
	A6
	A7
	A8

	B1
	B2
	B3
	B4
	B5
	B6
	B7
	B8

	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8

	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8

	E1
	E2
	E3
	E4
	E5
	E6
	E7
	E8

	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8

	G1
	G2
	G3
	G4
	G5
	G6
	G7
	G8

	H1
	H2
	H3
	H4
	H5
	H6
	H7
	H8
)

func GetCoordinatesBySquareName(n SquareName) (lin, col int) {
	switch n {
	case A1:
		return 0, 0
	case A2:
		return 1, 0
	case A3:
		return 2, 0
	case A4:
		return 3, 0
	case A5:
		return 4, 0
	case A6:
		return 5, 0
	case A7:
		return 6, 0
	case A8:
		return 7, 0

	case B1:
		return 0, 1
	case B2:
		return 1, 1
	case B3:
		return 2, 1
	case B4:
		return 3, 1
	case B5:
		return 4, 1
	case B6:
		return 5, 1
	case B7:
		return 6, 1
	case B8:
		return 7, 1

	case C1:
		return 0, 2
	case C2:
		return 1, 2
	case C3:
		return 2, 2
	case C4:
		return 3, 2
	case C5:
		return 4, 2
	case C6:
		return 5, 2
	case C7:
		return 6, 2
	case C8:
		return 7, 2

	case D1:
		return 0, 3
	case D2:
		return 1, 3
	case D3:
		return 2, 3
	case D4:
		return 3, 3
	case D5:
		return 4, 3
	case D6:
		return 5, 3
	case D7:
		return 6, 3
	case D8:
		return 7, 3

	case E1:
		return 0, 4
	case E2:
		return 1, 4
	case E3:
		return 2, 4
	case E4:
		return 3, 4
	case E5:
		return 4, 4
	case E6:
		return 5, 4
	case E7:
		return 6, 4
	case E8:
		return 7, 4

	case F1:
		return 0, 5
	case F2:
		return 1, 5
	case F3:
		return 2, 5
	case F4:
		return 3, 5
	case F5:
		return 4, 5
	case F6:
		return 5, 5
	case F7:
		return 6, 5
	case F8:
		return 7, 5

	case G1:
		return 0, 6
	case G2:
		return 1, 6
	case G3:
		return 2, 6
	case G4:
		return 3, 6
	case G5:
		return 4, 6
	case G6:
		return 5, 6
	case G7:
		return 6, 6
	case G8:
		return 7, 6

	case H1:
		return 0, 7
	case H2:
		return 1, 7
	case H3:
		return 2, 7
	case H4:
		return 3, 7
	case H5:
		return 4, 7
	case H6:
		return 5, 7
	case H7:
		return 6, 7
	case H8:
		return 7, 7
	}

	panic("Unknown square name")
}

func (n SquareName) String() string {
	switch n {
	case A1:
		return "A1"
	case A2:
		return "A2"
	case A3:
		return "A3"
	case A4:
		return "A4"
	case A5:
		return "A5"
	case A6:
		return "A6"
	case A7:
		return "A7"
	case A8:
		return "A8"

	case B1:
		return "B1"
	case B2:
		return "B2"
	case B3:
		return "B3"
	case B4:
		return "B4"
	case B5:
		return "B5"
	case B6:
		return "B6"
	case B7:
		return "B7"
	case B8:
		return "B8"

	case C1:
		return "C1"
	case C2:
		return "C2"
	case C3:
		return "C3"
	case C4:
		return "C4"
	case C5:
		return "C5"
	case C6:
		return "C6"
	case C7:
		return "C7"
	case C8:
		return "C8"

	case D1:
		return "D1"
	case D2:
		return "D2"
	case D3:
		return "D3"
	case D4:
		return "D4"
	case D5:
		return "D5"
	case D6:
		return "D6"
	case D7:
		return "D7"
	case D8:
		return "D8"

	case E1:
		return "E1"
	case E2:
		return "E2"
	case E3:
		return "E3"
	case E4:
		return "E4"
	case E5:
		return "E5"
	case E6:
		return "E6"
	case E7:
		return "E7"
	case E8:
		return "E8"

	case F1:
		return "F1"
	case F2:
		return "F2"
	case F3:
		return "F3"
	case F4:
		return "F4"
	case F5:
		return "F5"
	case F6:
		return "F6"
	case F7:
		return "F7"
	case F8:
		return "F8"

	case G1:
		return "G1"
	case G2:
		return "G2"
	case G3:
		return "G3"
	case G4:
		return "G4"
	case G5:
		return "G5"
	case G6:
		return "G6"
	case G7:
		return "G7"
	case G8:
		return "G8"

	case H1:
		return "H1"
	case H2:
		return "H2"
	case H3:
		return "H3"
	case H4:
		return "H4"
	case H5:
		return "H5"
	case H6:
		return "H6"
	case H7:
		return "H7"
	case H8:
		return "H8"
	}

	panic("Unknown square name")
}
