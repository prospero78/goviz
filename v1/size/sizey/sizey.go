// package sizey -- размер объекта по оси Y
//  Размер не может быть меньше нуля
package sizey

import (
	"fmt"

	"github.com/prospero78/goviz/v1/alias"
)

// SizeY -- размер объекта по оси Y
type SizeY struct {
	val alias.ASizeY
}

// NewSizeY -- возвращает новый размер по горизонтали
func NewSizeY(y alias.ASizeY) (*SizeY, error) {
	sf := &SizeY{}
	if err := sf.Set(y); err != nil {
		return nil, fmt.Errorf("NewSizeY(): in set size, err=\n\t%w", err)
	}
	return sf, nil
}

// Get -- возвращает хранимое значение
func (sf *SizeY) Get() alias.ASizeY {
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *SizeY) Set(val alias.ASizeY) error {
	if val < 0 {
		return fmt.Errorf("SizeY.Set(): y(%v)<0", val)
	}
	sf.val = val
	return nil
}
