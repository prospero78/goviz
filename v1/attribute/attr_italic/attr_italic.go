// package attr_italic -- признак "курсива" литеры
package attr_italic

import "github.com/nsf/termbox-go"

// Italic -- признак "курсива" литеры
type Italic struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewItalic -- возвращает новый признак "курсива" литеры
func NewItalic(fore *termbox.Attribute) *Italic {
	sf := &Italic{
		attr: fore,
	}
	if res := *fore & termbox.AttrCursive; res > 0 {
		sf.val = true
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Italic) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Italic) Set() {
	*sf.attr = *sf.attr | termbox.AttrCursive
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Italic) Reset() {
	*sf.attr = *sf.attr & ^termbox.AttrCursive
	sf.val = false
}
