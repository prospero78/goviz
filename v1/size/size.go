// package size -- размеры объекта
//   Размеры не могут быть меньше 0
package size

import (
	"fmt"

	"github.com/prospero78/goviz/v1/alias"
)

// Size -- размер виджета на экране
type Size struct {
	x alias.ASizeX // Размер по X
	y alias.ASizeY // Размер по Y
}

// NewSize -- возвращает новый размер объекта
func NewSize(x alias.ASizeX, y alias.ASizeY) (*Size, error) {
	{ // Предусловия

		if y < 0 {
			return nil, fmt.Errorf("NewSize(): y(%v)<0", x)
		}
	}
	sf := &Size{
		x: x,
		y: y,
	}
	return sf, nil
}
