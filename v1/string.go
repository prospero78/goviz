package goviz

import (
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
)

// String -- реализует строку
type String struct {
	lstLit []*lit.Lit // Набор литер для вывода на экран
	Pos    pos.Pos    // Позиция строки на экране
}
