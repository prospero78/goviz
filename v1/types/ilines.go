package types

// ILines -- интерфейс к линии
type ILines interface {
	// Redraw -- перерисовывает линию
	Redraw()
	// LitFill -- литера заливки линии (также содержит ForeAttr)
	LitFill() ILit
	// LitBeg -- литера начала линии
	LitBeg() ILit
	// LitEnd -- литера окончания линии
	LitEnd() ILit
	// ForeAttr -- возвращает атрибуты знакоместа (из LitFill)
	ForeAttr()IAttr
	// BackAttr -- возвращает объект фона знакоместа линии
	BackAttr()IAttr
}
