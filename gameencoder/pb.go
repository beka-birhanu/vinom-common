package gamepb

import (
	game_i "github.com/beka-birhanu/vinom-interfaces/game"
	"google.golang.org/protobuf/proto"
)

type Protobuf struct{}

func NewEncoder() (game_i.GameEncoder, error) {
	return &Protobuf{}, nil
}

// MarshalAction implements game.Encoder.
func (p *Protobuf) MarshalAction(a game_i.Action) ([]byte, error) {
	action := actionFromInterface(a)
	return proto.Marshal(action)
}

// MarshalCell implements game.Encoder.
func (p *Protobuf) MarshalCell(c game_i.Cell) ([]byte, error) {
	cell := cellFromInterface(c)
	return proto.Marshal(cell)
}

// MarshalCellPosition implements game.Encoder.
func (p *Protobuf) MarshalCellPosition(cp game_i.CellPosition) ([]byte, error) {
	cellPosition := cellPositionInterface(cp)
	return proto.Marshal(cellPosition)
}

// MarshalGameState implements game.Encoder.
func (p *Protobuf) MarshalGameState(gs game_i.GameState) ([]byte, error) {
	gameState := gameStateFromInterface(gs)
	return proto.Marshal(gameState)
}

// MarshalMaze implements game.Encoder.
func (p *Protobuf) MarshalMaze(m game_i.Maze) ([]byte, error) {
	maze := mazeFromInterface(m)
	return proto.Marshal(maze)
}

// MarshalPlayer implements game.Encoder.
func (p *Protobuf) MarshalPlayer(pl game_i.Player) ([]byte, error) {
	player := playerFromInterface(pl)
	return proto.Marshal(player)
}

// NewAction implements game.Encoder.
func (p *Protobuf) NewAction() game_i.Action {
	return &Action{}
}

// NewCell implements game.Encoder.
func (p *Protobuf) NewCell() game_i.Cell {
	return &Cell{}
}

// NewCellPosition implements game.Encoder.
func (p *Protobuf) NewCellPosition() game_i.CellPosition {
	return &Pos{}
}

// NewGameState implements game.Encoder.
func (p *Protobuf) NewGameState() game_i.GameState {
	return &GameState{}
}

// NewMaze implements game.Encoder.
func (p *Protobuf) NewMaze() game_i.Maze {
	return &Maze{}
}

// NewPlayer implements game.Encoder.
func (p *Protobuf) NewPlayer() game_i.Player {
	return &Player{}
}

// UnmarshalAction implements game.Encoder.
func (p *Protobuf) UnmarshalAction(b []byte) (game_i.Action, error) {
	action := &Action{}
	err := proto.Unmarshal(b, action)
	return action, err
}

// UnmarshalCell implements game.Encoder.
func (p *Protobuf) UnmarshalCell(b []byte) (game_i.Cell, error) {
	cell := &Cell{}
	err := proto.Unmarshal(b, cell)
	return cell, err
}

// UnmarshalCellPosition implements game.Encoder.
func (p *Protobuf) UnmarshalCellPosition(b []byte) (game_i.CellPosition, error) {
	pos := &Pos{}
	err := proto.Unmarshal(b, pos)
	return pos, err
}

// UnmarshalGameState implements game.Encoder.
func (p *Protobuf) UnmarshalGameState(b []byte) (game_i.GameState, error) {
	gameState := &GameState{}
	err := proto.Unmarshal(b, gameState)
	return gameState, err
}

// UnmarshalMaze implements game.Encoder.
func (p *Protobuf) UnmarshalMaze(b []byte) (game_i.Maze, error) {
	maze := &Maze{}
	err := proto.Unmarshal(b, maze)
	return maze, err
}

// UnmarshalPlayer implements game.Encoder.
func (p *Protobuf) UnmarshalPlayer(b []byte) (game_i.Player, error) {
	player := &Player{}
	err := proto.Unmarshal(b, player)
	return player, err
}
