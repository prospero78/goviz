package types

import "github.com/nsf/termbox-go"

// IBold -- атрибут жирности литеры
type IBold interface{
	// Get -- возвращает признак жирности
	Get()bool
	// Set -- устанавливает признак
	Set()
	// Reset -- сбросить признак
	Reset()
}

// IItalic -- атрибут курсив литеры
type IItalic interface{
	// Get -- возвращает признак курсива
	Get()bool
	// Set -- устанавливает признак
	Set()
	// Reset -- сбросить признак
	Reset()
}

// IUnderline -- атрибут подчёркнутости литеры
type IUnderline interface{
	// Get -- возвращает признак подчёркнутости
	Get()bool
	// Set -- устанавливает признак
	Set()
	// Reset -- сбросить признак
	Reset()
}

// IAttr -- атрибуты литеры и фона
type IAttr interface{
	// Bold -- возвращает атрибут "жирноты"
	Bold()IBold
	// Italic -- возвращает атрибут "курсив"
	Italic()IItalic
	// Underline -- возвращает атрибут "подчёркнуто"
	Underline()IUnderline
	// Blink -- возвращает атрибут мерцания
	// Visible -- возвращает атрибут видимости
	// Dimension -- возвращает атрибут размера
	// ForeAttr -- возвращает атрибуты литеры
	// Reverse -- возвращает атрибутвывернутости на изнанку
	ForeAttr()termbox.Attribute
	// BackAttr -- возвращает атрибуты фона
	BackAttr()termbox.Attribute
}