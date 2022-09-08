package goviz

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// ALit -- литера
type ALit string

// New -- инициализирует терминал
func New() error {
	err := termbox.Init()
	if err != nil {
		return fmt.Errorf("goviz.Init(): in init termbox=\n\t%w", err)
	}
	return nil
}

// Close -- закрывает работу терминала
func Close() {
	termbox.Close()
}

// Update -- обновляет терминал по требованию
//  ВНИМАНИЕ!
//  Вызывать после всех отрисовок
func Update() {
	termbox.Flush()
}

// OutLit -- выводит литеру в указанную позицию
func OutLit(posX APosX, posY APosY, foreAttr, beckAttr termbox.Attribute, lit rune) {
	termbox.SetCell(int(posX), int(posY), lit, foreAttr, beckAttr)
}

// OutString -- выводит строку в указанную позицию
func OutString(posX APosX, posY APosY, foreAttr, beckAttr termbox.Attribute, msg string) {
	for _, lit := range msg {
		OutLit(posX, posY, foreAttr, beckAttr, lit)
		posX += APosX(runewidth.RuneWidth(lit))
	}
}
