// package cons -- константы пакета
package cons

import (
	"github.com/prospero78/goviz/v1/alias"
)

const (
	DirectHor  = alias.ADirect(iota) // Горизонтальное направление (слева направо)
	DirectVert                       // Вертикальное направление (сверху вниз)
)

const ( // Стили границпрямоугольника
	RectangleStyleNone= alias.ARectangleStyle(iota+2)
)
