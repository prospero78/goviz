package screen

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/safebool"
	"github.com/prospero78/goviz/v1/size"
	"github.com/prospero78/goviz/v1/ticker"
)

// Screen -- глобальный объект экрана.
type Screen struct {
	size     size.Size          // Размер экрана
	isWork   *safebool.SafeBool // Потокобезопасный признак работы экрана
	chClose  chan int           // Канал закрытия экрана
	ForeAttr termbox.Attribute  // Атрибуты литеры
	BackAttr termbox.Attribute  // Атрибуты знакоместа
	ticker   *ticker.Ticker     // Тикер обновления экрана
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
	ticker, err := ticker.NewTicker(20)
	if err != nil {
		return nil, fmt.Errorf("goviz.GetScreen(): in create ticker, err=%w", err)
	}
	scr := &Screen{
		isWork:  safebool.NewSafeBool(),
		chClose: make(chan int, 2),
		ticker:  ticker,
	}
	scr.isWork.Set()
	scrX, scrY := termbox.Size()
	scr.size.X = alias.ASizeX(scrX)
	scr.size.Y = alias.ASizeY(scrY)
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

// IsWork -- возвращает признак работы экрана
func (sf *Screen) IsWork() bool {
	return sf.isWork.Get()
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
	for sf.isWork.Get() {
		time.Sleep(time.Millisecond * 5)
	}
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
