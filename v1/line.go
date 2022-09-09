package goviz

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/types"
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
	IsVisible bool       // Признак видимости линии
	LitFill   *lit.Lit   // Литера заполнителя
	pos       types.IPos // Позиция линии
	Len       ALen       // Длина линии
	Direct    ADirect    // Направление линии
}

// NewLine -- возвращает новую линию
func NewLine(pos types.IPos, _len ALen, lit *lit.Lit) (*Line, error) {
	{ // Предусловия
		if pos == nil {
			return nil, fmt.Errorf("NewLine(): IPos == nil")
		}
		if lit == nil {
			return nil, fmt.Errorf("NewLine(): lit == nil")
		}
	}

	sf := &Line{
		pos:     pos,
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
	sf.LitFill.Pos().PosX().Set(sf.pos.PosX().Get())
	for y := sf.pos.PosY().Get(); y < alias.APosY(sf.Len); y++ {
		if y >= alias.APosY(scrY) {
			return
		}
		sf.LitFill.Pos().PosY().Set(y)
		sf.LitFill.Redraw()
	}
}

// Перерисовывает линию горизонтально
func (sf *Line) redrawHor() {
	scrX, _ := termbox.Size()
	sf.LitFill.Pos().PosY().Set(sf.pos.PosY().Get())
	posX:=sf.pos.PosX().Get()
	for x := posX; x < alias.APosX(sf.Len); x++ {
		if x >= alias.APosX(scrX) {
			return
		}
		sf.LitFill.Pos().PosX().Set(x)
		sf.LitFill.Redraw()
	}
}
