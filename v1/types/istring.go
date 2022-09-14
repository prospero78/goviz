package types

// IString -- интерфейс к графической строке
type IString interface{
	// Get -- возвращает строку
	Get()string
	// Redraw -- перерисовывает строку
	Redraw()
}