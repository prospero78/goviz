// package attr_underline -- признак подчёркивания литеры
package attr_underline

import "github.com/nsf/termbox-go"

// Underline -- признак подчёркивания литеры
type Underline struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewUnderline -- возвращает новый признак подчёркивания литеры
func NewUnderline(fore *termbox.Attribute) *Underline {
	sf := &Underline{
		attr: fore,
	}
	if res := *fore & termbox.AttrUnderline; res > 0 {
		sf.val = true
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Underline) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Underline) Set() {
	*sf.attr = *sf.attr | termbox.AttrUnderline
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Underline) Reset() {
	*sf.attr = *sf.attr &^ termbox.AttrUnderline
	sf.val = false
}
