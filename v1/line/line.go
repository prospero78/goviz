package line

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/cons"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/types"
)

// Line -- линия на экране
type Line struct {
	scr       types.IScreen
	IsVisible bool          // Признак видимости линии
	LitFill   *lit.Lit      // Литера заполнителя
	pos       types.IPos    // Позиция линии
	Len       alias.ALen    // Длина линии
	Direct    alias.ADirect // Направление линии
}

// NewLine -- возвращает новую линию
func NewLine(scr types.IScreen, pos types.IPos, _len alias.ALen, lit *lit.Lit) (*Line, error) {
	{ // Предусловия
		if scr == nil {
			return nil, fmt.Errorf("NewLine(): IScreen == nil")
		}
		if pos == nil {
			return nil, fmt.Errorf("NewLine(): IPos == nil")
		}
		if lit == nil {
			return nil, fmt.Errorf("NewLine(): lit == nil")
		}
	}

	sf := &Line{
		scr:     scr,
		pos:     pos,
		Len:     _len,
		LitFill: lit,
	}
	return sf, nil
}

// Redraw -- перерисовывает линию
func (sf *Line) Redraw() {
	if sf.Direct == cons.DirectHor {
		sf.redrawHor()
		return
	}
	sf.redrawVert()
}

// Перерисовывает линию вертикально
func (sf *Line) redrawVert() {
	_, scrY := sf.scr.Size()
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
	posX := sf.pos.PosX().Get()
	for x := posX; x < alias.APosX(sf.Len); x++ {
		if x >= alias.APosX(scrX) {
			return
		}
		sf.LitFill.Pos().PosX().Set(x)
		sf.LitFill.Redraw()
	}
}
