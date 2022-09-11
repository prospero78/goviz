// package attr_dim -- признак размера литеры
package attr_dim

import "github.com/nsf/termbox-go"

// Dimension -- признак размера литеры
type Dimension struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewDimension -- возвращает новый признак размера литеры
func NewDimension(fore *termbox.Attribute) *Dimension {
	sf := &Dimension{
		attr: fore,
	}
	if res := *fore & termbox.AttrDim; res > 0 {
		sf.val = true
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Dimension) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Dimension) Set() {
	*sf.attr = *sf.attr | termbox.AttrDim
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Dimension) Reset() {
	*sf.attr = *sf.attr &^ termbox.AttrDim
	sf.val = false
}
