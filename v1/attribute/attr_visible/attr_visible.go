// package attr_visible -- признак видимости литеры
package attr_visible

import "github.com/nsf/termbox-go"

// Visible -- признак видимости литеры
type Visible struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewVisible -- возвращает новый признак видимости литеры
func NewVisible(fore *termbox.Attribute) *Visible {
	sf := &Visible{
		attr: fore,
	}
	// Атрибут видимости инверсный по отношению к Hidden
	if res := *fore & termbox.AttrHidden; res > 0 {
		sf.val = false
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Visible) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Visible) Set() {
	*sf.attr = *sf.attr &^ termbox.AttrHidden
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Visible) Reset() {
	*sf.attr = *sf.attr | termbox.AttrHidden
	sf.val = false
}
