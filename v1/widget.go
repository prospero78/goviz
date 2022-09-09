package goviz

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/size"
	"github.com/prospero78/goviz/v1/types"
)

// Widget -- базовый виджет для отрисовки
type Widget struct {
	pos         pos.Pos  // Позиция виджета на экране
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
	litFill lit.ALit) (*Widget, error) {
	var _litFill rune
	if litFill == "" {
		_litFill = []rune(" ")[0]
	}
	size,err:=size.NewSize(sizeX, sizeY)
	if err!=nil{
		return nil, fmt.Errorf("NewWidget(): in create ISize, err=\n\t%w",err)
	}
	sf := &Widget{
		pos: pos.Pos{
			X: posX,
			Y: posY,
		},
		size: size,
		foreAttr: foreAttr,
		backAttr: backAttr,
		litFill:  _litFill,
	}
	return sf, nil
}

// Redraw -- перерисовывает виджет на экране
func (sf *Widget) Redraw() {
	for x := sf.pos.X; x < sf.pos.X+alias.APosX(sf.size.SizeX().Get()); x++ {
		for y := sf.pos.Y; y < alias.APosY(sf.size.SizeY().Get()); y++ {
			// OutLit(x, y, sf.foreAttr, sf.backAttr, sf.litFill)
		}
	}
	if !sf.IsBorder {
		return
	}
	// Прочертить верхнюю границу
	for x := sf.pos.X; x < sf.pos.X+alias.APosX(sf.size.SizeX().Get()); x++ {
		// OutLit(x, sf.pos.Y, sf.BorderFore, sf.BorderBack, []rune("#")[0])
	}
	// Прочертить нижнюю границу
	for x := sf.pos.X; x < sf.pos.X+alias.APosX(sf.size.SizeX().Get()); x++ {
		// OutLit(x, sf.pos.Y+APosY(sf.size.Y)-2, sf.BorderFore, sf.BorderBack, []rune("#")[0])
	}
}
