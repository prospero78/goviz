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
		return
	}
	defer screen.Close()

	lit0, err := lit.NewLit(screen, pos.NewPos(2, 2), termbox.ColorWhite, termbox.ColorBlue, "#")
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in create lit0")
		return
	}
	lit0.ForeAttr().Bold().Set()
	lit0.Redraw()

	lit1, err := lit.NewLit(screen, pos.NewPos(4, 3), termbox.ColorWhite, termbox.ColorLightCyan, "5")
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in create lit1")
		return
	}
	lit1.ForeAttr().Underline().Set()
	lit1.Redraw()

	lit2, err := lit.NewLit(screen, pos.NewPos(6, 4), termbox.ColorBlue, termbox.ColorDefault, "^")
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_let: in create lit2")
		os.Exit(1)
	}
	lit2.ForeAttr().Italic().Set()
	lit2.Redraw()
	screen.Redraw()

	time.Sleep(time.Second * 3)
}
