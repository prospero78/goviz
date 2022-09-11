// package attribute -- атрибуты литеры/знакоместа
package attribute

import (
	"sync"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/attribute/attr_blink"
	"github.com/prospero78/goviz/v1/attribute/attr_bold"
	"github.com/prospero78/goviz/v1/attribute/attr_dim"
	"github.com/prospero78/goviz/v1/attribute/attr_italic"
	"github.com/prospero78/goviz/v1/attribute/attr_reverse"
	"github.com/prospero78/goviz/v1/attribute/attr_underline"
	"github.com/prospero78/goviz/v1/attribute/attr_visible"
	"github.com/prospero78/goviz/v1/types"
)

// Attribute -- атрибуты литеры/знакоместа
type Attribute struct {
	bold      types.IAttrBold      // Признак "жирности" литеры
	italic    types.IAttrItalic    // Признак курсива литеры
	underline types.IAttrUnderline // Признак подчёркивания литеры
	visible   types.IAttrVisible   // Признак видимости литеры
	dimension types.IAttrDimension // Признак размера литеры
	reverse   types.IAttrReverse   // Признак реверса литеры
	blink     types.IAttrBlink     // Признак моргания литеры
	fore      termbox.Attribute    // Атрибуты литеры
	block     sync.Mutex           // Блокировка на атрибуты
}

// NewAttribute -- возвращает новый атрибут литеры/знакоместа
func NewAttribute(fore termbox.Attribute) *Attribute {
	sf := &Attribute{
		bold:      attr_bold.NewBold(&fore),
		italic:    attr_italic.NewItalic(&fore),
		underline: attr_underline.NewUnderline(&fore),
		visible:   attr_visible.NewVisible(&fore),
		dimension: attr_dim.NewDimension(&fore),
		reverse:   attr_reverse.NewReverse(&fore),
		blink:     attr_blink.NewBlink(&fore),
		fore:      fore,
	}
	return sf
}

// Get -- возвращает атрибут литеры
func (sf *Attribute) Get() termbox.Attribute {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.fore
}

// Blink -- возвращает признак моргания литеры
func (sf *Attribute) Blink() types.IAttrBlink {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.blink
}

// Reverse -- возвращает признак реверса литеры
func (sf *Attribute) Reverse() types.IAttrReverse {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.reverse
}

// Dimension -- возвращает признак размера литеры
func (sf *Attribute) Dimension() types.IAttrDimension {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.visible
}

// Visible -- возвращает признак видмости литеры
func (sf *Attribute) Visible() types.IAttrVisible {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.visible
}

// Bold -- возвращает признак жирноты литеры
func (sf *Attribute) Bold() types.IAttrBold {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.bold
}

// Italic -- возвращает признак курсива литеры
func (sf *Attribute) Italic() types.IAttrItalic {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.italic
}

// Underline -- возвращает признак подчёркивания литеры
func (sf *Attribute) Underline() types.IAttrUnderline {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.underline
}
