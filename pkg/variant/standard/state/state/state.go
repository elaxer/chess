package state

import "github.com/elaxer/chess/pkg/chess"

var (
	Check = chess.NewState("check", chess.StateTypeThreat)
	Mate  = chess.NewState("mate", chess.StateTypeTerminal)

	Stalemate      = chess.NewState("stalemate", chess.StateTypeDraw)
	DrawFiftyMoves = chess.NewState("draw by fifty moves", chess.StateTypeDraw)
)
