package rules

type InvalidMoveReason int64

const (
	VALID_MOVE = iota
	NOT_COLORS_TURN
	MOVE_OBSTRUCTED
	MOVE_INCORRECT
	EMPTY_SQUARE
	KING_IN_CHECK
)

func (r InvalidMoveReason) String() string {
	switch r {
	case VALID_MOVE:
		return "VALID_MOVE"
	case NOT_COLORS_TURN:
		return "NOT_COLORS_TURN"
	case MOVE_OBSTRUCTED:
		return "MOVE_OBSTRUCTED"
	case MOVE_INCORRECT:
		return "MOVE_INCORRECT"
	case EMPTY_SQUARE:
		return "EMPTY_SQUARE"
	case KING_IN_CHECK:
		return "KING_IN_CHECK"
	}

	return "Unknown reason"
}
