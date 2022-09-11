// package main -- пример отрисовки прямоугольника
//  Для запуска выполнить:
//    go run ./example/rectangle/main.go
package main

import (
	"time"

	"github.com/nsf/termbox-go"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/rectangle"
	"github.com/prospero78/goviz/v1/screen"
	"github.com/prospero78/goviz/v1/size"
	"github.com/sirupsen/logrus"
)

func main() {
	scr, err := screen.GetScreen()
	if err != nil {
		logrus.WithError(err).Errorln("rectangle.go/maain(): in get IScreen, err=\n\t%w", err)
		return
	}
	defer scr.Close()
	{ // Первый прямоугольник
		rectPos := pos.NewPos(6, 6)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorCyan, termbox.ColorBlue, alias.ALit(rune(0x025e6)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Второй прямоугольник
		rectPos := pos.NewPos(4, 4)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorYellow, termbox.ColorLightGray, alias.ALit(rune(0x02591)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Третий прямоугольник
		rectPos := pos.NewPos(2, 2)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorLightGreen, termbox.ColorLightMagenta, alias.ALit(rune(0x025ec)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/maain(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle.Redraw()
	}
	scr.Redraw()
	time.Sleep(time.Second * 3)
}
