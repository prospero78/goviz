package goviz

import (
	"fmt"

	"github.com/nsf/termbox-go"
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
	LitFill   *Lit    // Литера заполнителя
	Pos       Pos     // Позиция линии
	Len       ALen    // Длина линии
	Direct    ADirect // Направление линии
}

// NewLine -- возвращает новую линию
func NewLine(pos Pos, _len ALen, lit *Lit) (*Line, error) {
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
	for y := sf.Pos.Y; y < APosY(sf.Len); y++ {
		if y >= APosY(scrY) {
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
	for x := sf.Pos.X; x < APosX(sf.Len); x++ {
		if x >= APosX(scrX) {
			return
		}
		sf.LitFill.Pos.X = x
		sf.LitFill.Redraw()
	}
}
