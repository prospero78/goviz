package line

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/cons"
	"github.com/prospero78/goviz/v1/types"
)

// Line -- линия на экране
type Line struct {
	scr types.IScreen
	// isVisible bool          // Признак видимости линии
	litFill  types.ILit    // Литера заполнителя
	litBegin types.ILit    // Литера начала
	litEnd   types.ILit    // Линия окончания
	pos      types.IPos    // Позиция линии
	Len      alias.ALen    // Длина линии
	Direct   alias.ADirect // Направление линии
}

// NewLine -- возвращает новую линию
func NewLine(scr types.IScreen, pos types.IPos, _len alias.ALen, lit types.ILit) (*Line, error) {
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
		litFill: lit,
	}
	return sf, nil
}

// LitFill -- возвращает литеру заполнения
func (sf *Line)LitFill()types.ILit{
	return sf.litFill
}

// LitBegin -- возвращает литеру начала
func (sf *Line)LitBeg()types.ILit{
		return sf.litBegin
}

// LitEnd -- возвращает литеру конца
func (sf *Line)LitEnd()types.ILit{
	return sf.litEnd
}

// LitEndSet -- устанавливает литеру конца
func (sf *Line) LitEndSet(lit types.ILit) error {
	if lit == nil {
		return fmt.Errorf("Line.LitEndSet(): lit == nil")
	}
	sf.litEnd = lit
	sf.Redraw()
	return nil
}

// LitBeginSet -- устанавливает литеру начала
func (sf *Line) LitBeginSet(lit types.ILit) error {
	if lit == nil {
		return fmt.Errorf("Line.LitBeginSet(): lit == nil")
	}
	sf.litBegin = lit
	sf.Redraw()
	return nil
}

// ForeAttr -- возвращает атрибуты литер заливки
func (sf *Line) ForeAttr() types.IAttr {
	return sf.litFill.ForeAttr()
}

// BackAttr -- возвращает атрибут фона линии
func (sf *Line) BackAttr() termbox.Attribute {
	return sf.litFill.BackAttr()
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
	litPosX := sf.litFill.Pos().PosX()
	litPosY := sf.litFill.Pos().PosY()

	litPosX.Set(selfPosX.Get())
	for y := selfPosY.Get(); y < selfPosY.Get()+alias.APosY(sf.Len); y++ {
		if y >= alias.APosY(scrY) {
			return
		}
		litPosY.Set(y)
		sf.litFill.Redraw()
	}
	sf.redrawBeg()
	sf.redrawEnd()
}

// Перерисовывает линию горизонтально
func (sf *Line) redrawHor() {
	scrX, _ := sf.scr.Size()
	selfPosX := sf.pos.PosX()
	selfPosY := sf.pos.PosY()
	litPosX := sf.litFill.Pos().PosX()
	litPosY := sf.litFill.Pos().PosY()

	litPosY.Set(selfPosY.Get())
	for x := selfPosX.Get(); x < selfPosX.Get()+alias.APosX(sf.Len); x++ {
		if x >= alias.APosX(scrX) {
			return
		}
		litPosX.Set(x)
		sf.litFill.Redraw()
	}
	sf.redrawBeg()
	sf.redrawEnd()
}

// Перерисовывает литеру начала (если есть)
func (sf *Line) redrawBeg() {
	if sf.litBegin == nil {
		return
	}
	x, y := sf.pos.Get()
	sf.litBegin.Pos().Set(x, y)
	sf.litBegin.Redraw()
}

// Перерисовывает литеру конца (если есть)
func (sf *Line) redrawEnd() {
	if sf.litEnd == nil {
		return
	}
	x, y := sf.pos.Get()
	if sf.Direct == cons.DirectHor {
		x += alias.APosX(sf.Len) - 1
	} else {
		y += alias.APosY(sf.Len) - 1
	}
	sf.litEnd.Pos().Set(x, y)
	sf.litEnd.Redraw()
}
