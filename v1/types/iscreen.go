package types

import (
	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/alias"
)

// IScreen -- интерфейс к экрану
type IScreen interface {
	// Redraw -- перерисовывает экран
	Redraw()
	// Close -- закрывае экран, блокирующий вызов
	Close()
	// Clear -- очищает экран установленными атрибутами
	Clear()
	// Fill -- выполняет заливку экрана установленными атрибутами
	Fill(lit string, foreAttr, backAttr termbox.Attribute)
	// IsWork -- возвращает признак работы экрана
	IsWork() bool
	// Size -- возвращает размеры экрана
	Size() (alias.ASizeX, alias.ASizeY)
}
