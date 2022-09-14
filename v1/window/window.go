// package window - реализация окна
package window

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/safebool"
	"github.com/prospero78/goviz/v1/tstring"
	"github.com/prospero78/goviz/v1/types"
)

// Window -- реализация окна
type Window struct {
	scr   types.IScreen
	title types.IString // Название окна
	isVisible *safebool.SafeBool // Признак видимости
}

// NewWindow -- возвращает новое окно
func NewWindow(scr types.IScreen) (*Window, error) {
	if scr == nil {
		return nil, fmt.Errorf("NewWindow(): IScreen == nil")
	}
	_pos := pos.NewPos(5, 5)
	title, err := tstring.NewTString(scr, _pos, termbox.ColorYellow, termbox.ColorLightRed, "Window")
	if err != nil {
		return nil, fmt.Errorf("NewWindow(): in create IString, err=\n\t%w", err)
	}
	sf := &Window{
		scr:   scr,
		title: title,
		isVisible: safebool.NewSafeBool(),
	}
	return sf, nil
}

// Show -- показать окно
func (sf *Window)Show(){
	sf.isVisible.Set()
	sf.Redraw()
}

// Redraw -- 