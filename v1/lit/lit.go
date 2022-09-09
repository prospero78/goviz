package lit

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/screen"
)

// ALit -- литера
type ALit string

// Lit -- литера на экране
type Lit struct {
	scr      *screen.Screen
	Pos      pos.Pos           // Позиция литеры на экране
	ForeAttr termbox.Attribute // Атрибуты литеры
	BackAttr termbox.Attribute // Атрибуты знакоместа
	Lit      rune              // Литера для печати
}

// NewLit -- возвращает новую литеру
func NewLit(pos pos.Pos, foreAttr, backAttr termbox.Attribute, lit ALit) (*Lit, error) {
	var _lit rune
	switch lit {
	case "":
		_lit = []rune(" ")[0]
	default:
		_lit = []rune(lit)[0]
	}
	scr, err := screen.GetScreen()
	if err != nil {
		return nil, fmt.Errorf("NewLit(): in get screen, err=%w", err)
	}
	sf := &Lit{
		scr:      scr,
		Pos:      pos,
		ForeAttr: foreAttr,
		BackAttr: backAttr,
		Lit:      _lit,
	}
	return sf, nil
}

// Redraw -- отрисовывает литеру на экране
func (sf *Lit) Redraw() {
	if sf.scr.IsWork() {
		termbox.SetCell(int(sf.Pos.X), int(sf.Pos.Y), sf.Lit,
			sf.ForeAttr, sf.BackAttr)
	}
}