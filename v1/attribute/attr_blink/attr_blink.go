// package attr_blink -- признак моргания литеры
package attr_blink

import "github.com/nsf/termbox-go"

// Blink -- признак моргания литеры
type Blink struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewBlink -- возвращает новый признак моргания литеры
func NewBlink(fore *termbox.Attribute) *Blink {
	sf := &Blink{
		attr: fore,
	}
	if res := *fore & termbox.AttrBlink; res > 0 {
		sf.val = true
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Blink) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Blink) Set() {
	*sf.attr = *sf.attr | termbox.AttrBlink
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Blink) Reset() {
	*sf.attr = *sf.attr &^ termbox.AttrBlink
	sf.val = false
}
