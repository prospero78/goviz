package goviz

import "github.com/nsf/termbox-go"

// ALit -- литера
type ALit string

// Lit -- литера на экране
type Lit struct {
	Pos      Pos               // Позиция литеры на экране
	ForeAttr termbox.Attribute // Атрибуты литеры
	BackAttr termbox.Attribute // Атрибуты знакоместа
	Lit      rune              // Литера для печати
}

// NewLit -- возвращает новую литеру
func NewLit(pos Pos, foreAttr, backAttr termbox.Attribute, lit ALit) *Lit {
	var _lit rune
	switch lit {
	case "":
		_lit = []rune(" ")[0]
	default:
		_lit = []rune(lit)[0]
	}
	sf := &Lit{
		Pos:      pos,
		ForeAttr: foreAttr,
		BackAttr: backAttr,
		Lit:      _lit,
	}
	return sf
}

// Redraw -- отрисовывает литеру на экране
func (sf *Lit) Redraw() {
	if scr.isWork.val == true {
		termbox.SetCell(int(sf.Pos.X), int(sf.Pos.Y), sf.Lit,
			sf.ForeAttr, sf.BackAttr)
	}
}
