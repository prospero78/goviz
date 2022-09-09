package goviz

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// Screen -- глобальный объект экрана.
type Screen struct {
	size     Size              // Размер экрана
	isWork   *SafeBool         // Потокобезопасный признак работы экрана
	chClose  chan int          // Канал закрытия экрана
	ForeAttr termbox.Attribute // Атрибуты литеры
	BackAttr termbox.Attribute // Атрибуты знакоместа
	ticker   *Ticker           // Тикер обновления экрана
}

var (
	scr *Screen // Глобальный объект экрана
)

// GetScreen -- возвращает новый экран
func GetScreen() (*Screen, error) {
	if scr != nil {
		return scr, nil
	}
	err := termbox.Init()
	if err != nil {
		return nil, fmt.Errorf("goviz.GetScreen(): in init termbox=\n\t%w", err)
	}
	ticker, err := NewTicker(20)
	if err != nil {
		return nil, fmt.Errorf("goviz.GetScreen(): in create ticker, err=%w", err)
	}
	scr := &Screen{
		isWork:  NewSafeBool(),
		chClose: make(chan int, 2),
		ticker:  ticker,
	}
	scr.isWork.Set()
	scrX, scrY := termbox.Size()
	scr.size.X = ASizeX(scrX)
	scr.size.Y = ASizeY(scrY)
	// Тут надо запустить тикер!!!!
	go scr.run()
	return scr, nil
}

// Redraw -- перерисовывает экран по требованию
func (sf *Screen) Redraw() {
	if sf.isWork.Get() {
		termbox.Flush()
	}
}

// Clear -- очищает экран
func (sf *Screen) Clear() {
	termbox.Clear(sf.ForeAttr, sf.BackAttr)
}

// Close -- закрывает экран
func (sf *Screen) Close() {
	if !sf.isWork.Get() {
		return
	}
	sf.chClose <- 1
}

// Главный цикл работы экрана
func (sf *Screen) run() {
	for sf.isWork.Get() {
		<-sf.chClose // Сигнал закрытия экрана
		sf.isWork.Reset()
		termbox.Close()
		return
	}
}
