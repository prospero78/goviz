package goviz

import "github.com/nsf/termbox-go"

// Widget -- базовый виджет для отрисовки
type Widget struct {
	pos        Pos  // Позиция виджета на экране
	size       Size // Размеры виджета
	foreAttr   termbox.Attribute
	backAttr   termbox.Attribute
	BorderFore termbox.Attribute
	BorderBack termbox.Attribute
	litFill    rune
	IsBorder   bool // Граница виджета
}

// NewWidget -- возвращает новый виджет
func NewWidget(posX APosX, posY APosY, sizeX ASizeX, sizeY ASizeY, foreAttr, backAttr termbox.Attribute, litFill ALit) *Widget {
	var _litFill rune
	if litFill == "" {
		_litFill = []rune(" ")[0]
	}
	sf := &Widget{
		pos: Pos{
			X: posX,
			Y: posY,
		},
		size: Size{
			X: sizeX,
			Y: sizeY,
		},
		foreAttr: foreAttr,
		backAttr: backAttr,
		litFill:  _litFill,
	}
	return sf
}

// Redraw -- перерисовывает виджет на экране
func (sf *Widget) Redraw() {
	for x := sf.pos.X; x < sf.pos.X+APosX(sf.size.X); x++ {
		for y := sf.pos.Y; y < APosY(sf.size.Y); y++ {
			OutLit(x, y, sf.foreAttr, sf.backAttr, sf.litFill)
		}
	}
	if !sf.IsBorder {
		return
	}
	// Прочертить верхнюю границу
	for x := sf.pos.X; x < sf.pos.X+APosX(sf.size.X); x++ {
		OutLit(x, sf.pos.Y, sf.BorderFore, sf.BorderBack, []rune("#")[0])
	}
	// Прочертить нижнюю границу
	for x := sf.pos.X; x < sf.pos.X+APosX(sf.size.X); x++ {
		OutLit(x, sf.pos.Y+APosY(sf.size.Y)-2, sf.BorderFore, sf.BorderBack, []rune("#")[0])
	}
}
