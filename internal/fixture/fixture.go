package fixture

import (
	"github.com/elaxer/chess/internal/chess/model"
)

var (
	Elaxer, _ = model.NewUser("elaxer", model.Password("elaxer"))
	John, _   = model.NewUser("john", model.Password("john"))
	Paul, _   = model.NewUser("paul", model.Password("paul"))
)

var Users = [...]*model.User{Elaxer, John, Paul}

var Games = [...]*model.Game{
	model.NewGameExisted(Elaxer.ID, John.ID, model.GameResultWinWhite, []string{
		"e2e4", "e7e5",
		"Ng1f3", "Nb8c6",
		"Bf1c3", "Ng8f6",
		"0-0", "d7d5",
		"e4xd5", "Nf6xd5",
		"Qd1e1", "Bf8c5",
		"Nf3xe5", "0-0",
		"Ne5xc6", "b7xc6",
		"c2c4", "Nd5f4",
		"Nb1c3", "Bc8g4",
		"g2g3", "Nf4h3+",
		"Kg1g2", "Qd8xd3",
		"Rf1h1", "Qd3f3+",
		"Kg2f1", "Qf3xh1#",
	}),
	model.NewGameExisted(John.ID, Elaxer.ID, model.GameResultWinBlack, []string{
		// https://www.chess.com/game/live/136458596805
		"e2e4", "e7e5",
		"Bf1c4", "Nb8c6",
		"Qd1h5", "Qd8e7",
		"Bc4f7+", "Qe7f7",
		"Qh5f7+", "Ke8f7",
		"d2d4", "e5d4",
		"Ng1f3", "Bf8b4+",
		"Nf3d2", "Ng8f6",
		"Nd2c3", "d7d6",
		"a2a3", "Bb4c3",
		"b2c3", "Bc8g4",
		"Nf3h3", "Rh8e8",
		"Bc1f4", "Re8e4+",
		"Ke1d2", "Rf4f4",
		"Nh3g5+", "Kf7f8",
		"Ng5h7+", "Nf6h7",
		"f2f3", "Bg4f5",
		"Rh1e1", "d6c3+",
		"Kd2c3", "Nh7f6",
		"g2g3", "Rf4f3+",
		"Kc3d2", "Nf6d4",
		"c2c3", "Nd4b3+",
		"Ke1e2", "Rc8c3",
		"Ra1d1", "Re8e8+",
		"Kd2f2", "Bc8c2",
		"Rd1d5", "Nb3d5",
		"Re1e8+", "Kf8e8",
		"g3g4", "Nd5d2",
		"h2h4", "Rf3f3+",
		"Kf2e2", "Re3e3+",
		"Kd2d3", "Rc3c3",
		"g4g5", "Ra8a3",
		"Kd3c2", "g7g6",
		"h4h5", "g6h5",
		"g5g6", "Ke8f8",
		"Kc2d2", "c7c5",
		"Ke2e2", "Ra3a1",
		"Kf2f3", "Rd8d1",
		"Ke2e4", "Nc3e3+",
		"Kf4f4", "d7d5",
		"Kg5g5", "d5d4",
		"Kh5h5", "d4d3",
		"Kh6h6", "d3d2",
		"Kh7h7", "Rh1h1#",
	}),
}
