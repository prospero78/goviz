package types

import "github.com/nsf/termbox-go"

// ILine -- интерфейс к линии
type ILine interface {
	// Redraw -- перерисовывает линию
	Redraw()
	// LitFill -- литера заливки линии (также содержит ForeAttr)
	LitFill() ILit
	// LitBeg -- литера начала линии
	LitBeg() ILit
	// LitEnd -- литера окончания линии
	LitEnd() ILit
	// ForeAttr -- возвращает атрибуты знакоместа (из LitFill)
	ForeAttr() IAttr
	// BackAttr -- возвращает объект фона знакоместа линии
	BackAttr() termbox.Attribute
}
