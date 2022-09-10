// package rectangle -- объект прямоугольника
//
//	Зарисовываетсебя из линий
package rectangle

import (
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/types"
)

// Rectangle -- прямоугольник из линий
type Rectangle struct {
	scr      types.IScreen
	LitFill  *lit.Lit       // Литера заливки прямоугольника
	Pos      types.IPos     // Позиция левого верхнего угла
	Size     types.ISize    // Размеры прямоугольника
	lstLines []types.ILines // Набор линий длязаполнения прямоугольника
	CornerLU *lit.Lit       // Литера верхнего левого угла
	CornerRU *lit.Lit       // Литера правого верхнего угла
	CornerLD *lit.Lit       // Литера нижнего левого угла
	CornerRD *lit.Lit       // Литера нижнего правого угла
}
