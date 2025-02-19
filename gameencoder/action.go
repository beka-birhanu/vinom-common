package gamepb

import (
	game_i "github.com/beka-birhanu/vinom-interfaces/game"
	"github.com/google/uuid"
)

var _ game_i.Action = &Action{}

func actionFromInterface(a game_i.Action) *Action {
	return &Action{
		Id:        a.GetID().String(),
		Direction: a.GetDirection(),
		From:      cellPositionInterface(a.RetriveFrom()),
	}
}

// RetriveFrom implements game.Action.
func (x *Action) RetriveFrom() game_i.CellPosition {
	return x.From
}

// SetFrom implements game.Action.
func (x *Action) SetFrom(c game_i.CellPosition) {
	x.From = cellPositionInterface(c)
}

// SetDirection implements game.Action.
func (x *Action) SetDirection(s string) {
	x.Direction = s
}

// GetID implements game.Action.
func (x *Action) GetID() uuid.UUID {
	id, _ := uuid.Parse(x.Id)
	return id
}

// SetID implements game.Action.
func (x *Action) SetID(i uuid.UUID) {
	x.Id = i.String()
}
