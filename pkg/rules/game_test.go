package rules

// func TestSwitchColorTurn(t *testing.T) {
// 	g := NewGame()

// 	assert.Equal(t, WHITES, g.NowPlays)
// 	switchColorTurn(g)
// 	assert.Equal(t, BLACKS, g.NowPlays)

// 	switchColorTurn(g)
// 	assert.Equal(t, WHITES, g.NowPlays)
// }

// func TestProcessCapture(t *testing.T) {
// 	g := NewGame()

// 	processCapture(board.WHITE_PAWN_A, board.A3, g)

// 	assert.Equal(t, 0, len(g.BlackCaptured))
// 	assert.Equal(t, 0, len(g.WhiteCaptured))

// 	g.Board.Move(board.WHITE_PAWN_A, board.A2, board.A6)
// 	processCapture(board.WHITE_PAWN_A, board.B7, g)
// 	assert.Equal(t, 0, len(g.BlackCaptured))
// 	assert.Equal(t, 1, len(g.WhiteCaptured))

// 	g.Board.Move(board.BLACK_PAWN_C, board.C7, board.C3)
// 	processCapture(board.BLACK_PAWN_C, board.B2, g)
// 	assert.Equal(t, 1, len(g.BlackCaptured))
// 	assert.Equal(t, 1, len(g.WhiteCaptured))
// }

// func TestProcessMove(t *testing.T) {
// 	g := NewGame()

// 	err := g.ProcessMove(board.WHITE_PAWN_E, board.E2, board.E4)
// 	assert.Nil(t, err)

// 	err = g.ProcessMove(board.BLACK_PAWN_E, board.E7, board.E5)
// 	assert.Nil(t, err)

// 	err = g.ProcessMove(board.WHITE_KNIGHT_KING, board.G1, board.F3)
// 	assert.Nil(t, err)

// 	err = g.ProcessMove(board.BLACK_KNIGHT_QUEEN, board.B8, board.C6)
// 	assert.Nil(t, err)

// 	err = g.ProcessMove(board.WHITE_BISHOP_KING, board.F1, board.B5)
// 	assert.Nil(t, err)

// 	err = g.ProcessMove(board.WHITE_PAWN_A, board.F1, board.B5)
// 	assert.NotNil(t, err)
// }
