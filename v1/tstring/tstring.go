// package tstring -- строка для вывода в произвольной позиции на экран
package tstring

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/types"
)

// TString -- реализует строку
type TString struct {
	lstLit []types.ILit // Набор литер для вывода на экран
	pos    types.IPos   // Позиция строки на экране
	val    string       // Значение строки
}

// NewTString -- возвращает новую строку
func NewTString(scr types.IScreen, _pos types.IPos, foreAttr, backAttr termbox.Attribute, strVal string) (*TString, error) {
	{ // Предусловия
		if scr == nil {
			return nil, fmt.Errorf("NewTString(): IScreen == nil")
		}
		if _pos == nil {
			return nil, fmt.Errorf("NewTString(): IPos == nil")
		}
	}
	lstLit := make([]types.ILit, 0)
	posX, posY := _pos.Get()
	ind := 0
	for _, _rune := range strVal {
		_pos := pos.NewPos(posX+alias.APosX(ind), posY)
		lit, err := lit.NewLit(scr, _pos, foreAttr, backAttr, alias.ALit(_rune))
		if err != nil {
			return nil, fmt.Errorf("NewTString(): in create ILit, err=\n\t%w", err)
		}
		lstLit = append(lstLit, lit)
		ind++
	}
	sf := &TString{
		pos:    _pos,
		lstLit: lstLit,
		val:    strVal,
	}
	return sf, nil
}

// Get -- возвращает хранимое значение
func (sf *TString) Get() string {
	return sf.val
}

// Redraw -- перерисовывает строку на экране
func (sf *TString) Redraw() {
	for _, lit := range sf.lstLit {
		lit.Redraw()
	}
}
