package pos

import "github.com/prospero78/goviz/v1/alias"

// Pos -- позиция элемента на экране
type Pos struct {
	X alias.APosX // Позиция по X
	Y alias.APosY // Позиция по Y
}

// NewPos -- возвращает новую позицию объекта
func NewPos(x alias.APosX, y alias.APosY) *Pos {
	return &Pos{
		X: x,
		Y: y,
	}
}
