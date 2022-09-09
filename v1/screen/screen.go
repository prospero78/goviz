package screen

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/sirupsen/logrus"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/safebool"
	"github.com/prospero78/goviz/v1/size"
	"github.com/prospero78/goviz/v1/ticker"
	"github.com/prospero78/goviz/v1/types"
)

// Screen -- глобальный объект экрана.
type Screen struct {
	size     types.ISize        // Размер экрана
	isWork   *safebool.SafeBool // Потокобезопасный признак работы экрана
	chClose  chan int           // Канал закрытия экрана
	ForeAttr termbox.Attribute  // Атрибуты литеры
	BackAttr termbox.Attribute  // Атрибуты знакоместа
	ticker   *ticker.Ticker     // Тикер обновления экрана
}

var (
	Scr types.IScreen // Глобальный объект экрана
)

// GetScreen -- возвращает новый экран
func GetScreen() (types.IScreen, error) {
	if Scr != nil {
		return Scr, nil
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
	size, err := size.NewSize(alias.ASizeX(scrX), alias.ASizeY(scrY))
	if err != nil {
		scr.Close()
		return nil, fmt.Errorf("GetScreen(): in create ISize, err=\n\t%w", err)
	}
	scr.size = size
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

// Fill -- заливает экран указанными атрибутами
func (sf *Screen) Fill(_lit string, foreAttr, backAttr termbox.Attribute) {
	lit, err := lit.NewLit(sf, pos.NewPos(0, 0), foreAttr, backAttr, alias.ALit(_lit))
	if err != nil {
		logrus.WithError(err).Errorln("Screen.Fill(): in create ILit")
		return
	}
	for x := 0; x < int(sf.size.SizeX().Get()); x++ {
		for y := 0; y < int(sf.size.SizeY().Get()); y++ {
			lit.Pos().Set(alias.APosX(x), alias.APosY(y))
			lit.Redraw()
		}
	}
	sf.Redraw()
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
