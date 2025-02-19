package gamepb

import game_i "github.com/beka-birhanu/vinom-interfaces/game"

func cellFromInterface(cell game_i.Cell) *Cell {
	return &Cell{
		NorthWall: cell.HasNorthWall(),
		SouthWall: cell.HasSouthWall(),
		EastWall:  cell.HasEastWall(),
		WestWall:  cell.HasWestWall(),
		Reward:    cell.GetReward(),
	}
}

func cellPositionInterface(cp game_i.CellPosition) *Pos {
	return &Pos{
		Row: cp.GetRow(),
		Col: cp.GetCol(),
	}

}

// HasEastWall implements game.Cell.
func (x *Cell) HasEastWall() bool {
	return x.EastWall
}

// HasNorthWall implements game.Cell.
func (x *Cell) HasNorthWall() bool {
	return x.NorthWall
}

// HasSouthWall implements game.Cell.
func (x *Cell) HasSouthWall() bool {
	return x.SouthWall
}

// HasWestWall implements game.Cell.
func (x *Cell) HasWestWall() bool {
	return x.WestWall
}

// SetEastWall implements game.Cell.
func (x *Cell) SetEastWall(value bool) {
	x.EastWall = value
}

// SetNorthWall implements game.Cell.
func (x *Cell) SetNorthWall(value bool) {
	x.NorthWall = value
}

// SetReward implements game.Cell.
func (x *Cell) SetReward(value int32) {
	x.Reward = value
}

// SetSouthWall implements game.Cell.
func (x *Cell) SetSouthWall(value bool) {
	x.SouthWall = value
}

// SetWestWall implements game.Cell.
func (x *Cell) SetWestWall(value bool) {
	x.WestWall = value
}

// SetCol implements game.CellPosition.
func (x *Pos) SetCol(c int32) {
	x.Col = c
}

// SetRow implements game.CellPosition.
func (x *Pos) SetRow(r int32) {
	x.Row = r
}
