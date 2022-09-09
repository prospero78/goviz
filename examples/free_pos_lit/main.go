// package main -- демонстрация свободного расположения литер
//  Для запуска выполнить:
//    go run ./examples/free_pos_lit/main.go
package main

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/sirupsen/logrus"

	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/screen"
)

func main() {
	screen, err := screen.GetScreen()
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in get screen")
		os.Exit(1)
	}
	defer screen.Close()

	lit0, err := lit.NewLit(screen, pos.Pos{X: 2, Y: 2}, termbox.ColorWhite, termbox.ColorBlue, "#")
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in create lit0")
		os.Exit(1)
	}
	lit0.Redraw()
	lit1, err := lit.NewLit(screen, pos.Pos{X: 4, Y: 3}, termbox.ColorWhite, termbox.ColorDefault, "5")
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in create lit1")
		os.Exit(1)
	}
	lit1.Redraw()
	lit2, err := lit.NewLit(screen, pos.Pos{X: 6, Y: 4}, termbox.ColorBlue, termbox.ColorDefault, "^")
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in create lit2")
		os.Exit(1)
	}
	lit2.Redraw()
	screen.Redraw()

	time.Sleep(time.Second * 3)
}
