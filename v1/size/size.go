// package size -- размеры объекта
//   Размеры не могут быть меньше 0
package size

import (
	"fmt"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/size/sizex"
	"github.com/prospero78/goviz/v1/size/sizey"
	"github.com/prospero78/goviz/v1/types"
)

// Size -- размер виджета на экране
type Size struct {
	x types.ISizeX // Размер по X
	y types.ISizeY // Размер по Y
}

// NewSize -- возвращает новый размер объекта
func NewSize(x alias.ASizeX, y alias.ASizeY) (*Size, error) {
	sizeX, err := sizex.NewSizeX(x)
	if err != nil {
		return nil, fmt.Errorf("NewSize(): in create ISizeX, err=\n\t%w", err)
	}
	sizeY, err := sizey.NewSizeY(y)
	if err != nil {
		return nil, fmt.Errorf("NewSize(): in create ISizeY, err=\n\t%w", err)
	}
	sf := &Size{
		x: sizeX,
		y: sizeY,
	}
	return sf, nil
}

// Get -- возвращает размеры объекта
func (sf *Size) Get() (alias.ASizeX, alias.ASizeY) {
	return sf.x.Get(), sf.y.Get()
}

// SizeX -- возвращает объекта размера по горизонтали
func (sf *Size) SizeX() types.ISizeX {
	return sf.x
}

// SizeY -- возвращает объекта размера по горизонтали
func (sf *Size) SizeY() types.ISizeY {
	return sf.y
}

// Set -- устанавливает размера объекта
func (sf *Size) Set(x alias.ASizeX, y alias.ASizeY) error {
	if !(x > 0 && y > 0) {
		return fmt.Errorf("Size.Set(): x(%v)<0; y(%v)<0", x, y)
	}
	_ = sf.x.Set(x)
	_ = sf.y.Set(y)
	return nil
}
