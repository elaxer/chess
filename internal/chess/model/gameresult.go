package model

// GameResult представляет исход игры.
type GameResult int

const (
	// GameResultInProcess - игра в процессе и еще не завершена.
	GameResultInProcess GameResult = 1 << iota
	// GameResultWinWhite - игра завершена и белые выиграли.
	// todo пересмотреть логику
	GameResultWinWhite
	// GameResultWinBlack - игра завершена и черные выиграли.
	GameResultWinBlack
	// GameResultLeaved - игра завершена потому, что один из игроков покинул игру.
	GameResultLeaved
	// GameResultDraw - игра завершена и это ничья (или пат).
	GameResultDraw
)

func (r GameResult) IsInProcess() bool {
	return r&GameResultInProcess != 0
}

func (r GameResult) IsWinWhite() bool {
	return r&GameResultWinWhite != 0
}

func (r GameResult) IsWinBlack() bool {
	return r&GameResultWinBlack != 0
}

func (r GameResult) IsLeaved() bool {
	return r&GameResultLeaved != 0
}

func (r GameResult) IsDraw() bool {
	return r&GameResultDraw != 0
}

func (r GameResult) IsGameOver() bool {
	return r != GameResultInProcess
}
