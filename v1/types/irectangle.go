package types

// IRectangle -- интерфейс к прямоугольнику
type IRectangle interface {
	// Redraw -- перерисовать прямоугольник
	Redraw()
	// LitFill -- литера-заполнитель
	LitFill() ILit
	// BorderStyle -- возвращает стиль границ прямоуголника
	BorderStyle() IRectangleStyle
}
