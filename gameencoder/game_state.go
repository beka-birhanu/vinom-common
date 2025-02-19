package gamepb

import game_i "github.com/beka-birhanu/vinom-interfaces/game"

// Maze-related functions

// Height implements game.Maze.
func (x *Maze) Height() int {
	return len(x.Grid)
}

// Width implements game.Maze.
func (x *Maze) Width() int {
	panic("unimplemented")
}

// GetTotalReward implements game.Maze.
func (x *Maze) GetTotalReward() int32 {
	panic("unimplemented")
}

// NewValidMove implements game.Maze.
func (x *Maze) NewValidMove(game_i.CellPosition, string) (game_i.Move, error) {
	panic("unimplemented")
}

// InBound implements game.Maze.
func (x *Maze) InBound(row int, col int) bool {
	panic("unimplemented")
}

// IsValidMove implements game.Maze.
func (x *Maze) IsValidMove(move game_i.Move) bool {
	panic("unimplemented")
}

// Move implements game.Maze.
func (x *Maze) Move(move game_i.Move) (int32, error) {
	panic("unimplemented")
}

// RemoveReward implements game.Maze.
func (x *Maze) RemoveReward(pos game_i.CellPosition) error {
	panic("unimplemented")
}

// RetriveGrid implements game.Maze.
func (x *Maze) RetriveGrid() [][]game_i.Cell {
	maze := make([][]game_i.Cell, 0)
	for _, row := range x.Grid {
		new_row := make([]game_i.Cell, 0)
		for _, cell := range row.Cells {
			new_row = append(new_row, cell)
		}
		maze = append(maze, new_row)
	}
	return maze
}

// PopulateReward implements i.Maze.
func (x *Maze) PopulateReward(r struct {
	RewardOne      int32
	RewardTwo      int32
	RewardTypeProb float32
}) error {
	panic("unimplemented")
}

// SetGrid implements game.Maze.
func (x *Maze) SetGrid(g [][]game_i.Cell) {
	maze := make([]*Maze_Row, 0)
	for _, row := range g {
		maze_row := &Maze_Row{
			Cells: make([]*Cell, 0),
		}
		for _, cell := range row {
			maze_row.Cells = append(maze_row.Cells, cellFromInterface(cell))
		}
		maze = append(maze, maze_row)
	}
	x.Grid = maze
}

// GameState-related functions

// RetriveMaze implements game.GameState.
func (x *GameState) RetriveMaze() game_i.Maze {
	return x.GetMaze()
}

// RetrivePlayers implements game.GameState.
func (x *GameState) RetrivePlayers() []game_i.Player {
	players := make([]game_i.Player, 0)
	for _, p := range x.GetPlayers() {
		players = append(players, p)
	}
	return players
}

// SetMaze implements game.GameState.
func (x *GameState) SetMaze(m game_i.Maze) {
	x.Maze = mazeFromInterface(m)
}

// SetPlayers implements game.GameState.
func (x *GameState) SetPlayers(p []game_i.Player) {
	players := make([]*Player, len(x.Players))
	for _, player := range p {
		players = append(players, playerFromInterface(player))
	}
	x.Players = players
}

// SetVersion implements game.GameState.
func (x *GameState) SetVersion(v int64) {
	x.Version = v
}

// Helper functions for converting interfaces

// mazeFromInterface converts a game.Maze interface to a *Maze structure.
func mazeFromInterface(m game_i.Maze) *Maze {
	maze := &Maze{}
	maze.SetGrid(m.RetriveGrid())

	return maze
}

func gameStateFromInterface(gs game_i.GameState) *GameState {
	gameState := &GameState{}

	gameState.SetVersion(gs.GetVersion())
	gameState.SetMaze(gs.RetriveMaze())
	gameState.SetPlayers(gs.RetrivePlayers())

	return gameState
}
