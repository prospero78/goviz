// package posy -- положение объекта по оси Х
//  Положение может быть меньше нуля
package posy

import (
	"github.com/prospero78/goviz/v1/alias"
)

// PosY -- положение объекта по оси Y
type PosY struct {
	val alias.APosY
}

// NewPosY -- возвращает новое положение по горизонтали
func NewPosY(Y alias.APosY) *PosY {
	sf := &PosY{
		val: Y,
	}
	return sf
}

// Get -- возвращает хранимое значение
func (sf *PosY) Get() alias.APosY {
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *PosY) Set(val alias.APosY) {
	sf.val = val
}
