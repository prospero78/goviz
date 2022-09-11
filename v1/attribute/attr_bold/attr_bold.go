// package attr_bold -- признак "жирноты" литеры
package attr_bold

import "github.com/nsf/termbox-go"

// Bold -- признак "жирноты" литеры
type Bold struct {
	val  bool               // Хранимое значение
	attr *termbox.Attribute // Ссылка на родительский атрибут
}

// NewBold -- возвращает новый признак "жирноты" литеры
func NewBold(fore *termbox.Attribute) *Bold {
	sf := &Bold{
		attr: fore,
	}
	if res := *fore & termbox.AttrBold; res > 0 {
		sf.val = true
	}
	return sf
}

// Get -- возвращает хранимый признак
func (sf *Bold) Get() bool {
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *Bold) Set() {
	*sf.attr = *sf.attr | termbox.AttrBold
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *Bold) Reset() {
	*sf.attr = *sf.attr &^ termbox.AttrBold
	sf.val = false
}
