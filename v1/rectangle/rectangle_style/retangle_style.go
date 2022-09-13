// package rectangle_style -- стиль границы прямоугольника
package rectangle_style

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/cons"
	"github.com/prospero78/goviz/v1/line"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/types"
)

// RectangleStyle -- стиль границы прямоугольника
type RectangleStyle struct {
	scr      types.IScreen
	pos      types.IPos
	size     types.ISize
	style    alias.ARectangleStyle // Стиль прямоугольника
	cornerLU types.ILit            // Литера верхнего левого угла
	cornerRU types.ILit            // Литера правого верхнего угла
	cornerLD types.ILit            // Литера нижнего левого угла
	cornerRD types.ILit            // Литера нижнего правого угла
}

// NewRectangleStyle -- возвращает новый стиль границ прямоугольника
func NewRectangleStyle(scr types.IScreen, pos types.IPos, size types.ISize, style alias.ARectangleStyle) (*RectangleStyle, error) {
	{ // Предусловия
		if scr == nil {
			return nil, fmt.Errorf("NewRectangleStyle(): IScreen is empty")
		}
		if pos == nil {
			return nil, fmt.Errorf("NewRectangleStyle(): IPos is empty")
		}
		if size == nil {
			return nil, fmt.Errorf("NewRectangleStyle(): ISize is empty")
		}
		if !(cons.RectangleStyleNone <= style && style <= cons.RectangleStyleDouble) {
			return nil, fmt.Errorf("NewRectangleStyle(): style(%v) not in list style", style)
		}
	}
	sf := &RectangleStyle{
		scr:   scr,
		pos:   pos,
		size:  size,
		style: style,
	}
	return sf, nil
}

// Get -- возвращает хранимый стиль прямоугольника
func (sf *RectangleStyle) Get() alias.ARectangleStyle {
	return sf.style
}

// Set -- устанавливает хранимый стиль прямоугольника
func (sf *RectangleStyle) Set(val alias.ARectangleStyle) error {
	if !(cons.RectangleStyleNone <= val && val <= cons.RectangleStyleDouble) {
		return fmt.Errorf("RectangleStyle.Set(): val(%v) not in list style", val)
	}
	sf.style = val
	return nil
}

// LitCornerRDSet -- устанавливает литеру нижнего правого угла
func (sf *RectangleStyle) LitCornerRDSet(lit types.ILit) {
	sf.cornerRD = lit
	if lit == nil {
		return
	}
	posX := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
	posY := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
	pos := pos.NewPos(posX, posY)
	sf.cornerRD.PosSet(pos)
	sf.Redraw()
}

// LitCornerRUSet -- устанавливает литеру верхнего правого угла
func (sf *RectangleStyle) LitCornerRUSet(lit types.ILit) {
	sf.cornerRU = lit
	if lit == nil {
		return
	}
	posX := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
	posY := sf.pos.PosY().Get()
	pos := pos.NewPos(posX, posY)
	sf.cornerRU.PosSet(pos)
	sf.Redraw()
}

// LitCornerLDSet -- устанавливает литеру нижнего левого угла
func (sf *RectangleStyle) LitCornerLDSet(lit types.ILit) {
	sf.cornerLD = lit
	if lit == nil {
		return
	}
	posX := sf.pos.PosX().Get()
	posY := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
	pos := pos.NewPos(posX, posY)
	sf.cornerLD.PosSet(pos)
	sf.Redraw()
}

// LitCornerLUSet -- устанавливает литеру верхнего левого угла
func (sf *RectangleStyle) LitCornerLUSet(lit types.ILit) {
	sf.cornerLU = lit
	if lit == nil {
		return
	}
	sf.cornerLU.PosSet(sf.pos)
	sf.Redraw()
}

// Redraw -- перерисовывает прямоугольник
func (sf *RectangleStyle) Redraw() {
	switch sf.style {
	case cons.RectangleStyleNone: // Нет стиля
		sf.redrawCornes()
	case cons.RectangleStyleSimple: // Простая граница
		sf.redrawSimple()
	case cons.RectangleStyleSingle: // Одиночная граница
		sf.redrawSingle()
	}
}

// Перерисовывает границу с одиночной линией
func (sf *RectangleStyle) redrawSingle() {
	{ // Верхняя граница
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get()
		posLit := pos.NewPos(x, y)
		sizeX := alias.ALen(sf.size.SizeX().Get())
		// fmt.Printf("\nx=%v, y=%v, sizeX=%v\n", x, y, sizeX)
		lit, _ := lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9472}))
		posLine := pos.NewPos(x, y)
		lineUp, _ := line.NewLine(sf.scr, posLine, sizeX, lit)
		lineUp.Redraw()
	}
	{ // Нижняя граница
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
		posLit := pos.NewPos(x, y)
		sizeX := alias.ALen(sf.size.SizeX().Get())
		lit, _ := lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9472}))
		posLine := pos.NewPos(x, y)
		lineDown, _ := line.NewLine(sf.scr, posLine, sizeX, lit)
		lineDown.Redraw()
	}
	{ // Левая граница
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get()
		posLit := pos.NewPos(x, y)
		sizeY := alias.ALen(sf.size.SizeY().Get())
		lit, _ := lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9474}))
		posLine := pos.NewPos(x, y)
		lineLeft, _ := line.NewLine(sf.scr, posLine, sizeY, lit)
		lineLeft.Direct = cons.DirectVert
		lineLeft.Redraw()
	}
	{ // Правая граница
		x := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
		y := sf.pos.PosY().Get()
		posLit := pos.NewPos(x, y)
		sizeY := alias.ALen(sf.size.SizeY().Get())
		lit, _ := lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9474}))
		posLine := pos.NewPos(x, y)
		lineRight, _ := line.NewLine(sf.scr, posLine, sizeY, lit)
		lineRight.Direct = cons.DirectVert
		lineRight.Redraw()
	}
	// Прочертить угловые литеры
	{ // LU
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get()
		posLit := pos.NewPos(x, y)
		sf.cornerLU, _ = lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9484}))
	}
	{ // RU
		x := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
		y := sf.pos.PosY().Get() //+alias.APosY(sf.size.SizeY().Get())
		posLit := pos.NewPos(x, y)
		sf.cornerRU, _ = lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9488}))
	}
	{ // LD
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
		posLit := pos.NewPos(x, y)
		sf.cornerLD, _ = lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9492}))
	}
	{ // RD
		x := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
		y := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
		posLit := pos.NewPos(x, y)
		sf.cornerRD, _ = lit.NewLit(sf.scr, posLit, termbox.ColorWhite, termbox.ColorBlue, alias.ALit([]rune{9496}))
	}
	sf.redrawCornes()
}

// Перерисовывает границу по стилю фона верхнего левого угла
func (sf *RectangleStyle) redrawSimple() {
	if sf.cornerLU == nil {
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get()
		pos := pos.NewPos(x, y)
		sf.cornerLU, _ = lit.NewLit(sf.scr, pos, termbox.ColorWhite, termbox.ColorBlue, " ")
	}
	{ // Верхняя граница
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get()
		pos := pos.NewPos(x, y)
		sizeX := alias.ALen(sf.size.SizeX().Get())
		lineUp, _ := line.NewLine(sf.scr, pos, sizeX, sf.cornerLU)
		lineUp.Redraw()
	}
	{ // Нижняя граница
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get())
		pos := pos.NewPos(x, y)
		sizeX := alias.ALen(sf.size.SizeX().Get())
		lineDown, _ := line.NewLine(sf.scr, pos, sizeX, sf.cornerLU)
		lineDown.Redraw()
	}
	{ // Левая граница
		x := sf.pos.PosX().Get()
		y := sf.pos.PosY().Get()
		pos := pos.NewPos(x, y)
		sizeY := alias.ALen(sf.size.SizeY().Get())
		lineLeft, _ := line.NewLine(sf.scr, pos, sizeY, sf.cornerLU)
		lineLeft.Direct = cons.DirectVert
		lineLeft.Redraw()
	}
	{ // Правая граница
		x := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
		y := sf.pos.PosY().Get() //+alias.APosY(sf.size.SizeY().Get())
		pos := pos.NewPos(x, y)
		sizeY := alias.ALen(sf.size.SizeY().Get())
		lineRight, _ := line.NewLine(sf.scr, pos, sizeY, sf.cornerLU)
		lineRight.Direct = cons.DirectVert
		lineRight.Redraw()
	}
}

// Перерисовывает углы границы
func (sf *RectangleStyle) redrawCornes() {
	if sf.cornerLU != nil {
		sf.cornerLU.Redraw()
	}
	if sf.cornerRU != nil {
		sf.cornerRU.Redraw()
	}
	if sf.cornerLD != nil {
		sf.cornerLD.Redraw()
	}
	if sf.cornerRD != nil {
		sf.cornerRD.Redraw()
	}
}
