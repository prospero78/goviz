// package main -- показывает пример создания экрана
//  Для запуска выполнить:
//    go run ./examples/screen/main.go
package main

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/sirupsen/logrus"

	"github.com/prospero78/goviz/v1/screen"
)

func main() {
	screen, err :=screen.GetScreen()
	if err != nil {
		logrus.WithError(err).Errorln("screen.go/main(): in get screen")
		os.Exit(1)
	}
	defer screen.Close()
	screen.Clear()
	screen.Fill("*", termbox.ColorBlack, termbox.ColorGreen)
	time.Sleep(time.Second * 3)
}
