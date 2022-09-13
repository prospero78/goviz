package types

import "github.com/prospero78/goviz/v1/alias"

// IRectangleStyle -- интерфейс к типу прямоугольника
type IRectangleStyle interface {
	// LitCornerLUSet -- устанавливает литеру левого верхнего угла
	LitCornerLUSet(ILit)
	// LitCornerRUSet -- устанавливает литеру правого верхнего угла
	LitCornerRUSet(ILit)
	// LitCornerLDSet -- устанавливает литеру левого нижнего угла
	LitCornerLDSet(ILit)
	// LitCornerRDSet -- устанавливает литеру правого нижнего угла
	LitCornerRDSet(ILit)
	// Get -- возвращает хранимый стиль прямоугольника
	Get() alias.ARectangleStyle
	// Set -- устанавливает стиль прямоугольника
	Set(alias.ARectangleStyle) error
	// Redraw -- перерисовывает границы прямоугольника
	Redraw()
}
