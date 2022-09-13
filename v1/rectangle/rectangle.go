// package rectangle -- объект прямоугольника
//
//	Зарисовываетсебя из линий
package rectangle

import (
	"fmt"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/cons"
	"github.com/prospero78/goviz/v1/line"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/rectangle/rectangle_style"
	"github.com/prospero78/goviz/v1/types"
)

// Rectangle -- прямоугольник из линий
type Rectangle struct {
	scr      types.IScreen
	litFill  types.ILit            // Литера заливки прямоугольника
	pos      types.IPos            // Позиция левого верхнего угла
	size     types.ISize           // Размеры прямоугольника
	lstLines []types.ILine         // Набор линий для заполнения прямоугольника
	style    types.IRectangleStyle // Стиль границ прямоугольника
}

// NewRectangle -- возвращает новый прмоугольник
func NewRectangle(scr types.IScreen, _pos types.IPos, _size types.ISize, lit types.ILit) (*Rectangle, error) {
	{ // Предусловия
		if scr == nil {
			return nil, fmt.Errorf("NewRectangle(): IScreen is nil")
		}
		if _pos == nil {
			return nil, fmt.Errorf("NewRectangle(): IPos is nil")
		}
	}
	_len := alias.ALen(_size.SizeX().Get())
	lstLines := make([]types.ILine, 0, _size.SizeX().Get())
	for y := _pos.PosY().Get(); y < _pos.PosY().Get()+alias.APosY(_size.SizeY().Get()); y++ {
		posLine := pos.NewPos(_pos.PosX().Get(), y)
		_line, err := line.NewLine(scr, posLine, _len, lit)
		if err != nil {
			return nil, fmt.Errorf("NewRectangle(): in create line y(%v), err=\n\t%w", y, err)
		}
		lstLines = append(lstLines, _line)
	}
	rectStyle, err := rectangle_style.NewRectangleStyle(scr, _pos, _size, cons.RectangleStyleNone)
	if err != nil {
		return nil, fmt.Errorf("NewRectangle(): in create IRectangleStyle, err=\n\t%w", err)
	}
	sf := &Rectangle{
		scr:      scr,
		pos:      _pos,
		size:     _size,
		lstLines: lstLines,
		style:    rectStyle,
	}
	_ = sf.style.Set(cons.RectangleStyleNone)
	return sf, nil
}

// LitFill -- возвращает литеру-заполнитель
func (sf *Rectangle) LitFill() types.ILit {
	return sf.litFill
}

// Redraw -- перерисовывает прямоугольник
func (sf *Rectangle) Redraw() {
	for _, line := range sf.lstLines {
		line.Redraw()
	}
	sf.style.Redraw()
}

// BorderStyle -- возвращает стиль границ прямоугольника
func (sf *Rectangle) BorderStyle() types.IRectangleStyle {
	return sf.style
}
