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
	litFill  types.ILit    // Литера заливки прямоугольника
	pos      types.IPos    // Позиция левого верхнего угла
	size     types.ISize   // Размеры прямоугольника
	lstLines []types.ILine // Набор линий для заполнения прямоугольника
	cornerLU types.ILit    // Литера верхнего левого угла
	cornerRU types.ILit    // Литера правого верхнего угла
	cornerLD types.ILit    // Литера нижнего левого угла
	cornerRD types.ILit    // Литера нижнего правого угла
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
		pos:      _pos,
		size:     _size,
		lstLines: lstLines,
	}
	return sf, nil
}

// LitCornerRDSet -- устанавливает литеру нижнего правого угла
func (sf *Rectangle) LitCornerRDSet(lit types.ILit) {
	sf.cornerRD = lit
	if lit == nil {
		return
	}
	posX := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
	posY := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
	pos := pos.NewPos(posX, posY)
	sf.cornerRD.PosSet(pos)
	sf.Redraw()
}

// LitCornerRUSet -- устанавливает литеру верхнего правого угла
func (sf *Rectangle) LitCornerRUSet(lit types.ILit) {
	sf.cornerRU = lit
	if lit == nil {
		return
	}
	posX := sf.pos.PosX().Get() + alias.APosX(sf.size.SizeX().Get()) - 1
	posY := sf.pos.PosY().Get()
	pos := pos.NewPos(posX, posY)
	sf.cornerRU.PosSet(pos)
	sf.Redraw()
}

// LitCornerLDSet -- устанавливает литеру нижнего левого угла
func (sf *Rectangle) LitCornerLDSet(lit types.ILit) {
	sf.cornerLD = lit
	if lit == nil {
		return
	}
	posX := sf.pos.PosX().Get()
	posY := sf.pos.PosY().Get() + alias.APosY(sf.size.SizeY().Get()) - 1
	pos := pos.NewPos(posX, posY)
	sf.cornerLD.PosSet(pos)
	sf.Redraw()
}

// LitCornerLUSet -- устанавливает литеру верхнего левого угла
func (sf *Rectangle) LitCornerLUSet(lit types.ILit) {
	sf.cornerLU = lit
	if lit == nil {
		return
	}
	sf.cornerLU.PosSet(sf.pos)
	sf.Redraw()
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
	if sf.cornerLU != nil {
		sf.cornerLU.Redraw()
	}
	if sf.cornerRU != nil {
		sf.cornerRU.Redraw()
	}
	if sf.cornerLD != nil {
		sf.cornerLD.Redraw()
	}
	if sf.cornerRD != nil {
		sf.cornerRD.Redraw()
	}
}
