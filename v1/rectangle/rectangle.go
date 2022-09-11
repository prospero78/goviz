// package rectangle -- объект прямоугольника
//
//	Зарисовываетсебя из линий
package rectangle

import (
	"fmt"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/line"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/types"
)

// Rectangle -- прямоугольник из линий
type Rectangle struct {
	scr      types.IScreen
	LitFill  types.ILit     // Литера заливки прямоугольника
	Pos      types.IPos    // Позиция левого верхнего угла
	Size     types.ISize   // Размеры прямоугольника
	lstLines []types.ILine // Набор линий для заполнения прямоугольника
	CornerLU types.ILit       // Литера верхнего левого угла
	CornerRU types.ILit      // Литера правого верхнего угла
	CornerLD types.ILit       // Литера нижнего левого угла
	CornerRD types.ILit      // Литера нижнего правого угла
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
	sf := &Rectangle{
		scr:      scr,
		Pos:      _pos,
		lstLines: lstLines,
	}
	return sf, nil
}

// LitBeg -- возвращает литеру начала линии
func (sf *Rectangle)LitBeg()types.ILit{
	return sf.LitBeg()
}

// Redraw -- перерисовывает прямоугольник
func (sf *Rectangle) Redraw() {

}
