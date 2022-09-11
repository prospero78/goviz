package lit

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/attribute"
	"github.com/prospero78/goviz/v1/types"
)

// Lit -- литера на экране
type Lit struct {
	scr      types.IScreen
	pos      types.IPos        // Позиция литеры на экране
	ForeAttr types.IAttr       // Атрибуты литеры
	BackAttr termbox.Attribute // Атрибуты знакоместа
	Lit      rune              // Литера для печати
}

// NewLit -- возвращает новую литеру
func NewLit(scr types.IScreen, pos types.IPos, foreAttr, backAttr termbox.Attribute, lit alias.ALit) (*Lit, error) {
	{ // Предусловия
		if scr == nil {
			return nil, fmt.Errorf("NewLit(): IScreen == nil")
		}
		if pos == nil {
			return nil, fmt.Errorf("NewLit(): IPos == nil")
		}
	}

	var _lit rune
	switch lit {
	case "":
		_lit = []rune(" ")[0]
	default:
		_lit = []rune(lit)[0]
	}
	sf := &Lit{
		scr:      scr,
		pos:      pos,
		ForeAttr: attribute.NewAttribute(foreAttr),
		BackAttr: backAttr,
		Lit:      _lit,
	}
	return sf, nil
}

// Redraw -- отрисовывает литеру на экране
func (sf *Lit) Redraw() {
	if sf.scr.IsWork() {
		x, y := sf.pos.Get()
		termbox.SetCell(int(x), int(y), sf.Lit,
			sf.ForeAttr.ForeAttr(), sf.BackAttr)
	}
}

// Pos -- возвращает объект позиции литеры
func (sf *Lit) Pos() types.IPos {
	return sf.pos
}
