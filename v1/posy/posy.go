// package posy -- положение объекта по оси Х
//  Положение может быть меньше нуля
package posy

import (
	"github.com/prospero78/goviz/v1/alias"
)

// PosY -- положение объекта по оси Y
type PosY struct {
	val alias.ASizeY
}

// NewPosY -- возвращает новое положение по горизонтали
func NewPosY(Y alias.ASizeY) *PosY {
	sf := &PosY{
		val: Y,
	}
	return sf
}

// Get -- возвращает хранимое значение
func (sf *PosY) Get() alias.ASizeY {
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *PosY) Set(val alias.ASizeY) {
	sf.val = val
}
