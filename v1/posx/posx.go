// package posx -- положение объекта по оси Х
//  Положение может быть меньше нуля
package posx

import (
	"github.com/prospero78/goviz/v1/alias"
)

// PosX -- положение объекта по оси X
type PosX struct {
	val alias.ASizeX
}

// NewPosX -- возвращает новое положение по горизонтали
func NewPosX(x alias.ASizeX) *PosX {
	sf := &PosX{
		val:x,
	}
	return sf
}

// Get -- возвращает хранимое значение
func (sf *PosX) Get() alias.ASizeX {
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *PosX) Set(val alias.ASizeX) {
	sf.val = val
}
