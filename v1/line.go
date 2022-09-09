package goviz

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
)

// ALen -- длина
type ALen int

// ADirect -- направление (горизонтальное/вертикальное)
type ADirect int

const (
	DirectHor  = ADirect(iota) // Горизонтальное направление (слева направо)
	DirectVert                 // Вертикальное направление (сверху вниз)
)

// Line -- линия на экране
type Line struct {
	IsVisible bool    // Признак видимости линии
	LitFill   *lit.Lit    // Литера заполнителя
	Pos       pos.Pos     // Позиция линии
	Len       ALen    // Длина линии
	Direct    ADirect // Направление линии
}

// NewLine -- возвращает новую линию
func NewLine(pos pos.Pos, _len ALen, lit *lit.Lit) (*Line, error) {
	if lit == nil {
		return nil, fmt.Errorf("NewLine(): lit == nil")
	}
	sf := &Line{
		Pos:     pos,
		Len:     _len,
		LitFill: lit,
	}
	return sf, nil
}

// Redraw -- перерисовывает линию
func (sf *Line) Redraw() {
	if sf.Direct == DirectHor {
		sf.redrawHor()
		return
	}
	sf.redrawVert()
}

// Перерисовывает линию вертикально
func (sf *Line) redrawVert() {
	_, scrY := termbox.Size()
	sf.LitFill.Pos.X = sf.Pos.X
	for y := sf.Pos.Y; y < alias.APosY(sf.Len); y++ {
		if y >= alias.APosY(scrY) {
			return
		}
		sf.LitFill.Pos.Y = y
		sf.LitFill.Redraw()
	}
}

// Перерисовывает линию горизонтально
func (sf *Line) redrawHor() {
	scrX, _ := termbox.Size()
	sf.LitFill.Pos.Y = sf.Pos.Y
	for x := sf.Pos.X; x < alias.APosX(sf.Len); x++ {
		if x >= alias.APosX(scrX) {
			return
		}
		sf.LitFill.Pos.X = x
		sf.LitFill.Redraw()
	}
}
