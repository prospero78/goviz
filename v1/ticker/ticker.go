package ticker

import (
	"fmt"
	"time"

	"github.com/prospero78/goviz/v1/safebool"
)

// ATick -- интервал в миллисекундах
type ATick int

// Ticker -- генерирует тики с заданным интервалом
type Ticker struct {
	val    ATick              // Интервал в миллисекундах
	chTick chan int           // Канал для отсчёта тиков
	isWork *safebool.SafeBool // Признак работы тикера
	chEnd  chan int           // Признак окончания работы тикера
	chWait chan int           // Канал ожидания вызова
}

// NewTicker -- возвращает новый тикер
func NewTicker(msec ATick) (*Ticker, error) {
	if msec < 1 {
		return nil, fmt.Errorf("NewTicker(): msec(%v)<1", msec)
	}
	sf := &Ticker{
		val:    msec,
		chTick: make(chan int, 2),
		isWork: safebool.NewSafeBool(),
		chEnd:  make(chan int, 2),
		chWait: make(chan int, 1),
	}
	go sf.gener()
	return sf, nil
}

// Генерирует тики в отдельном потоке
func (sf *Ticker) gener() {
	for sf.isWork.Get() {
		time.Sleep(time.Millisecond * time.Duration(sf.val))
		sf.chTick <- 1
	}
}

// Wait -- ожидание тика (блокирующий вызов)
func (sf *Ticker) Wait() <-chan int {
	return sf.chWait
}

// Close -- закрывает тикер по требованию
func (sf *Ticker) Close() {
	if sf.isWork.Get() {
		sf.chEnd <- 1
	}
}

// Run -- главный цикл тикера, должен работать в отдельном потоке
func (sf *Ticker) Run() {
	for sf.isWork.Get() {
		select {
		case <-sf.chEnd: // Остановка тикера
			sf.isWork.Reset()
		case <-sf.chTick: // пора дёрнуть вызов
			sf.chWait <- 1
		}
	}
}
