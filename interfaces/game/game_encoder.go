package i

import "github.com/google/uuid"

// Cell represents a cell in the maze with walls and rewards.
type Cell interface {
	HasNorthWall() bool
	SetNorthWall(bool)
	HasSouthWall() bool
	SetSouthWall(bool)
	HasEastWall() bool
	SetEastWall(bool)
	HasWestWall() bool
	SetWestWall(bool)
	GetReward() int32
	SetReward(int32)
}

// CellPosition represents the position of a cell in the maze.
type CellPosition interface {
	GetRow() int32
	SetRow(int32)
	GetCol() int32
	SetCol(int32)
}

// Move represents a move operation in the maze.
type Move interface {
	From() CellPosition
	SetFrom(CellPosition)
	To() CellPosition
	SetTo(CellPosition)
}

// Player represents a player in the maze.
type Player interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
	RetrivePos() CellPosition
	SetPos(CellPosition)
	GetReward() int32
	SetReward(int32)
}

// Maze represents a maze structure with cells, dimensions, and actions.
type Maze interface {
	NewValidMove(CellPosition, string) (Move, error)
	IsValidMove(move Move) bool
	InBound(row, col int) bool
	Move(move Move) (int32, error)
	String() string
	Height() int
	Width() int
	GetTotalReward() int32
	RemoveReward(pos CellPosition) error
	RetriveGrid() [][]Cell
	SetGrid([][]Cell)
	PopulateReward(r struct {
		RewardOne      int32
		RewardTwo      int32
		RewardTypeProb float32
	}) error
}

// Action represents an action performed in the game.
type Action interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
	GetDirection() string
	SetDirection(string)
	RetriveFrom() CellPosition
	SetFrom(CellPosition)
}

// TODO: include time left!!!
// GameState represents the state of the game at a specific version.
type GameState interface {
	GetVersion() int64
	SetVersion(int64)
	RetriveMaze() Maze
	SetMaze(Maze)
	RetrivePlayers() []Player
	SetPlayers([]Player)
}

type GameEncoder interface {
	NewCell() Cell
	NewCellPosition() CellPosition
	NewPlayer() Player
	NewMaze() Maze
	NewAction() Action
	NewGameState() GameState

	MarshalCell(Cell) ([]byte, error)
	MarshalCellPosition(CellPosition) ([]byte, error)
	MarshalPlayer(Player) ([]byte, error)
	MarshalMaze(Maze) ([]byte, error)
	MarshalAction(Action) ([]byte, error)
	MarshalGameState(GameState) ([]byte, error)

	UnmarshalCell([]byte) (Cell, error)
	UnmarshalCellPosition([]byte) (CellPosition, error)
	UnmarshalPlayer([]byte) (Player, error)
	UnmarshalMaze([]byte) (Maze, error)
	UnmarshalAction([]byte) (Action, error)
	UnmarshalGameState([]byte) (GameState, error)
}
