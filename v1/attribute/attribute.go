// package attribute -- атрибуты литеры/знакоместа
package attribute

import (
	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/attribute/attr_bold"
	"github.com/prospero78/goviz/v1/attribute/attr_italic"
	"github.com/prospero78/goviz/v1/types"
)

// Attribute -- атрибуты литеры/знакоместа
type Attribute struct {
	bold   types.IBold       // Признак "жирности" литеры
	italic types.IItalic     // Признак курсива литеры
	fore   termbox.Attribute // Атрибуты литеры
	back   termbox.Attribute // Атрибуты фона
}

// NewAttribute -- возвращает новый атрибут литеры/знакоместа
func NewAttribute(fore, back termbox.Attribute) *Attribute {
	sf := &Attribute{
		bold: attr_bold.NewBold(&fore),
		italic:attr_italic.NewItalic(&fore),
		fore: fore,
		back: back,
	}
	return sf
}

// Bold -- возвращает признак жирноты литеры
func (sf *Attribute) Bold() types.IBold {
	return sf.bold
}

// Italic -- возвращает признак курсива литеры
func (sf *Attribute) Italic() types.IItalic {
	return sf.italic
}
