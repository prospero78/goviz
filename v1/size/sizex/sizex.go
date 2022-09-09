// package sizex -- размер объекта по оси Х
//  Размер не может быть меньше нуля
package sizex

import (
	"fmt"

	"github.com/prospero78/goviz/v1/alias"
)

// SizeX -- размер объекта по оси X
type SizeX struct {
	val alias.ASizeX
}

// NewSizeX -- возвращает новый размер по горизонтали
func NewSizeX(x alias.ASizeX) (*SizeX, error) {
	sf := &SizeX{}
	if err := sf.Set(x); err != nil {
		return nil, fmt.Errorf("NewSizeX(): in set size, err=\n\t%w", err)
	}
	return sf, nil
}

// Get -- возвращает хранимое значение
func (sf *SizeX) Get() alias.ASizeX {
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *SizeX) Set(val alias.ASizeX) error {
	if val < 0 {
		return fmt.Errorf("SizeX.Set(): x(%v)<0", val)
	}
	sf.val = val
	return nil
}
