package types

import "github.com/nsf/termbox-go"

// IAttrLit -- базовый атрибут литеры
type IAttrLit interface {
	// Get -- возвращает хранимый атрибут признак
	Get() bool
	// Set -- устанавливает признак
	Set()
	// Reset -- сбросить признак
	Reset()
}

// IAttrBold -- атрибут жирности литеры
type IAttrBold interface {
	IAttrLit
}

// IAttrItalic -- атрибут курсив литеры
type IAttrItalic interface {
	IAttrLit
}

// IAttrUnderline -- атрибут подчёркнутости литеры
type IAttrUnderline interface {
	IAttrLit
}

// IAttrBlink -- атрибут моргания литеры литеры
type IAttrBlink interface {
	IAttrLit
}

// IAttrVisible -- атрибут видимости литеры литеры
type IAttrVisible interface {
	IAttrLit
}

// IAttrDimension -- атрибут размера литеры литеры
type IAttrDimension interface {
	IAttrLit
}

// IAttrReverse -- атрибут реверса литеры литеры(???)
type IAttrReverse interface {
	IAttrLit
}

// IAttr -- атрибуты литеры
type IAttr interface {
	// Bold -- возвращает атрибут "жирноты"
	Bold() IAttrBold
	// Italic -- возвращает атрибут "курсив"
	Italic() IAttrItalic
	// Underline -- возвращает атрибут "подчёркнуто"
	Underline() IAttrUnderline
	// Blink -- возвращает атрибут мерцания
	Blink() IAttrBlink
	// Visible -- возвращает атрибут видимости
	Visible() IAttrVisible
	// Dimension -- возвращает атрибут размера
	Dimension() IAttrDimension
	// Reverse -- возвращает атрибут вывернутости на изнанку (???)
	Reverse() IAttrReverse
	// Get -- возвращает атрибуты литеры
	Get() termbox.Attribute
}
