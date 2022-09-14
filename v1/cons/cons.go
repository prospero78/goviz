// package cons -- константы пакета
package cons

import (
	"github.com/prospero78/goviz/v1/alias"
)

const (
	DirectHor  = alias.ADirect(iota) // Горизонтальное направление (слева направо)
	DirectVert                       // Вертикальное направление (сверху вниз)
)

const ( // Стили границ прямоугольника
	RectangleStyleNone       = alias.ARectangleStyle(iota + 2) // Нет никакого стиля, форма одежды №8 -- что надели, то и носим
	RectangleStyleSimple                                       // Граница без заполнителя
	RectangleStyleSingle                                       // Одиночная линия по всему периметру
	RectangleStyleDouble                                       // Двойная линия по всему периметру
)
