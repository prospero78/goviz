package pos

import (
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/posx"
	"github.com/prospero78/goviz/v1/posy"
	"github.com/prospero78/goviz/v1/types"
)

// Pos -- позиция элемента на экране
type Pos struct {
	x types.IPosX // Позиция по X
	y types.IPosY // Позиция по Y
}

// NewPos -- возвращает новую позицию объекта
func NewPos(x alias.APosX, y alias.APosY) *Pos {
	return &Pos{
		x: posx.NewPosX(x),
		y: posy.NewPosY(y),
	}
}

// Get -- возвращает позицию объекта
func (sf *Pos) Get() (alias.APosX, alias.APosY) {
	return sf.x.Get(), sf.y.Get()
}

// Set -- устанавливает позицию объекта
func (sf *Pos) Set(x alias.APosX, y alias.APosY) {
	sf.x.Set(x)
	sf.y.Set(y)
}

// PosX -- возвращает объект положения по X
func (sf *Pos)PosX()types.IPosX{
	return sf.x
}

// PosY -- возвращает объект положения по Y
func (sf *Pos)PosY()types.IPosY{
	return sf.y
}