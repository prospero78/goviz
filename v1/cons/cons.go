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
	RectangleStyleNone       = alias.ARectangleStyle(iota + 2) // Нет никакого стиля
	RectangleStyleSimple                                       // Граница без заполнителя
	RectangleStyleSingleVert                                   // Одиночная линия вертикально
	RectangleStyleSingleHor                                    // Одиночная линия горизонтально
	RectangleStyleSingle                                       // Одиночная линия по всему периметру
	RectangleStyleDoubleVert                                   // Двойная линия вертикально
	RectangleStyleDoubleHor                                    // Двойная линия горизонтально
	RectangleStyleDouble                                       // Двойная линия по всему периметру
)
