package gamepb

import (
	game_i "github.com/beka-birhanu/vinom-interfaces/game"
	"github.com/google/uuid"
)

var _ game_i.Player = &Player{}

func playerFromInterface(player game_i.Player) *Player {
	return &Player{
		Pos:    cellPositionInterface(player.RetrivePos()),
		Reward: player.GetReward(),
		Id:     player.GetID().String(),
	}
}

// GetID implements game.Player.
func (x *Player) GetID() uuid.UUID {
	id, _ := uuid.Parse(x.Id)
	return id
}

// RetrivePos implements game.Player.
func (x *Player) RetrivePos() game_i.CellPosition {
	return x.Pos
}

// SetID implements game.Player.
func (x *Player) SetID(i uuid.UUID) {
	x.Id = i.String()
}

// SetPos implements game.Player.
func (x *Player) SetPos(p game_i.CellPosition) {
	x.Pos = cellPositionInterface(p)
}

// SetReward implements game.Player.
func (x *Player) SetReward(r int32) {
	x.Reward = r
}
