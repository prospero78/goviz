package goviz

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/size"
	"github.com/prospero78/goviz/v1/types"
)

// Widget -- базовый виджет для отрисовки
type Widget struct {
	pos         types.IPos  // Позиция виджета на экране
	size        types.ISize // Размеры виджета
	foreAttr    termbox.Attribute
	backAttr    termbox.Attribute
	BorderLeft  Line // Левая граница виджета
	BorderRight Line // Правая граница виджета
	BorderTop   Line // Верхняя граница виджета
	BorderDown  Line // Нижняя граница виджета
	litFill     rune
	IsBorder    bool // Граница виджета
}

// NewWidget -- возвращает новый виджет
func NewWidget(posX alias.APosX, posY alias.APosY,
	sizeX alias.ASizeX, sizeY alias.ASizeY,
	foreAttr, backAttr termbox.Attribute,
	litFill alias.ALit) (*Widget, error) {
	var _litFill rune
	if litFill == "" {
		_litFill = []rune(" ")[0]
	}
	size, err := size.NewSize(sizeX, sizeY)
	if err != nil {
		return nil, fmt.Errorf("NewWidget(): in create ISize, err=\n\t%w", err)
	}
	sf := &Widget{
		pos:      pos.NewPos(posX, posY),
		size:     size,
		foreAttr: foreAttr,
		backAttr: backAttr,
		litFill:  _litFill,
	}
	return sf, nil
}

// Redraw -- перерисовывает виджет на экране
func (sf *Widget) Redraw() {
	posX := sf.pos.PosX().Get()
	for x := posX; x < posX+alias.APosX(sf.size.SizeX().Get()); x++ {
		posY := sf.pos.PosY().Get()
		for y := posY; y < alias.APosY(sf.size.SizeY().Get()); y++ {
			// OutLit(x, y, sf.foreAttr, sf.backAttr, sf.litFill)
		}
	}
	if !sf.IsBorder {
		return
	}
	// Прочертить верхнюю границу
	for x := posX; x < posX+alias.APosX(sf.size.SizeX().Get()); x++ {
		// OutLit(x, sf.pos.Y, sf.BorderFore, sf.BorderBack, []rune("#")[0])
	}
	// Прочертить нижнюю границу
	for x := posX; x < posX+alias.APosX(sf.size.SizeX().Get()); x++ {
		// OutLit(x, sf.pos.Y+APosY(sf.size.Y)-2, sf.BorderFore, sf.BorderBack, []rune("#")[0])
	}
}
