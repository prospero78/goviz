// package attr_reverse -- признак "реверса" литеры
package attr_reverse

import "github.com/nsf/termbox-go"

// Reverse -- признак "реверса" литеры
type Reverse struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewReverse -- возвращает новый признак "реверса" литеры
func NewReverse(fore *termbox.Attribute) *Reverse {
	sf := &Reverse{
		attr: fore,
	}
	if res := *fore & termbox.AttrReverse; res > 0 {
		sf.val = true
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Reverse) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Reverse) Set() {
	*sf.attr = *sf.attr | termbox.AttrReverse
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Reverse) Reset() {
	*sf.attr = *sf.attr &^ termbox.AttrReverse
	sf.val = false
}
