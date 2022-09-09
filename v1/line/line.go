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
	scr types.IScreen
	// isVisible bool          // Признак видимости линии
	LitFill  *lit.Lit      // Литера заполнителя
	LitBegin *lit.Lit      // Литера начала
	LitEnd   *lit.Lit      // Линия окончания
	pos      types.IPos    // Позиция линии
	Len      alias.ALen    // Длина линии
	Direct   alias.ADirect // Направление линии
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
	selfPosX := sf.pos.PosX()
	selfPosY := sf.pos.PosY()
	litPosX := sf.LitFill.Pos().PosX()
	litPosY := sf.LitFill.Pos().PosY()

	litPosX.Set(selfPosX.Get())
	for y := selfPosY.Get(); y < selfPosY.Get()+alias.APosY(sf.Len); y++ {
		if y >= alias.APosY(scrY) {
			return
		}
		litPosY.Set(y)
		sf.LitFill.Redraw()
	}
	sf.redrawBeg()
	sf.redrawEnd()
}

// Перерисовывает линию горизонтально
func (sf *Line) redrawHor() {
	scrX, _ := termbox.Size()
	selfPosX := sf.pos.PosX()
	selfPosY := sf.pos.PosY()
	litPosX := sf.LitFill.Pos().PosX()
	litPosY := sf.LitFill.Pos().PosY()

	litPosY.Set(selfPosY.Get())
	for x := selfPosX.Get(); x < selfPosX.Get()+alias.APosX(sf.Len); x++ {
		if x >= alias.APosX(scrX) {
			return
		}
		litPosX.Set(x)
		sf.LitFill.Redraw()
	}
	sf.redrawBeg()
	sf.redrawEnd()
}

// Перерисовывает литеру начала (если есть)
func (sf *Line) redrawBeg() {
	if sf.LitBegin == nil {
		return
	}
	x, y := sf.pos.Get()
	sf.LitBegin.Pos().Set(x, y)
	sf.LitBegin.Redraw()
}

// Перерисовывает литеру конца (если есть)
func (sf *Line) redrawEnd() {
	if sf.LitEnd == nil {
		return
	}
	x, y := sf.pos.Get()
	if sf.Direct == cons.DirectHor {
		x += alias.APosX(sf.Len) - 1
	} else {
		y += alias.APosY(sf.Len) - 1
	}
	sf.LitEnd.Pos().Set(x, y)
	sf.LitEnd.Redraw()
}
