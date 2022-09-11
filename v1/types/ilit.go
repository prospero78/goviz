package types

import "github.com/nsf/termbox-go"

// ILIt -- интерфейс к литере
type ILit interface {
	// Redraw -- перерисовывает своё отображение
	Redraw()
	// Pos -- возвращает позицию литеры
	Pos() IPos
	// PosSet -- устанавливает позицию литеры
	PosSet(IPos) error
	// ForeAttr -- возвращает атрибуты литеры
	ForeAttr() IAttr
	// BackAttr -- возвращает атрибуты фона
	BackAttr() termbox.Attribute
}
