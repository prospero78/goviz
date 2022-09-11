package types

// IRectangle -- интерфейс к прямоугольнику
type IRectangle interface{
	// Redraw -- перерисовать прямоугольник
	Redraw()
	// LitFill -- литера-заполнитель
	LitFill()ILit
	// LitCornerLUSet -- устанавливает литеру левого верхнего угла
	LitCornerLUSet(ILit)
	// LitCornerRUSet -- устанавливает литеру правого верхнего угла
	LitCornerRUSet(ILit)
	// LitCornerLDSet -- устанавливает литеру левого нижнего угла
	LitCornerLDSet(ILit)
	// LitCornerRDSet -- устанавливает литеру правого нижнего угла
	LitCornerRDSet(ILit)
}