// package main -- пример отрисовки прямоугольника
//  Для запуска выполнить:
//    go run ./example/rectangle/main.go
package main

import (
	"time"

	"github.com/nsf/termbox-go"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/cons"
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
		logrus.WithError(err).Errorln("rectangle.go/main(): in get IScreen, err=\n\t%w", err)
		return
	}
	defer scr.Close()
	{ // Первый прямоугольник
		rectPos := pos.NewPos(6, 6)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorCyan, termbox.ColorBlue, alias.ALit(rune(0x025e6)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle1, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle1.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Второй прямоугольник
		rectPos := pos.NewPos(4, 4)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorYellow, termbox.ColorLightGray, alias.ALit(rune(0x02591)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle2, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle2.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Третий прямоугольник
		rectPos := pos.NewPos(2, 2)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorLightGreen, termbox.ColorYellow, alias.ALit(rune(0x025ec)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		litCornLU, err := lit.NewLit(scr, posLit, termbox.ColorWhite, termbox.ColorLightCyan, "_")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerLU, err=\n\t%w", err)
			return
		}
		posRU:= pos.NewPos(2, 5)
		litCornRU, err := lit.NewLit(scr, posRU, termbox.ColorWhite, termbox.ColorYellow, "*")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerRU, err=\n\t%w", err)
			return
		}
		posLD:= pos.NewPos(2, 5)
		litCornLD, err := lit.NewLit(scr, posLD, termbox.ColorWhite, termbox.ColorRed, "x")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerRU, err=\n\t%w", err)
			return
		}
		posRD:= pos.NewPos(2, 5)
		litCornRD, err := lit.NewLit(scr, posRD, termbox.ColorWhite, termbox.ColorRed, "#")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerRU, err=\n\t%w", err)
			return
		}
		rectangle3, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle3.BorderStyle().LitCornerLUSet(litCornLU)
		rectangle3.BorderStyle().LitCornerRUSet(litCornRU)
		rectangle3.BorderStyle().LitCornerLDSet(litCornLD)
		rectangle3.BorderStyle().LitCornerRDSet(litCornRD)
		rectangle3.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Четвёртый прямоугольник (со стилем)
		rectPos := pos.NewPos(30, 10)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorYellow, termbox.ColorLightGray, alias.ALit(rune(0x02591)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		litCornLU, err := lit.NewLit(scr, posLit, termbox.ColorWhite, termbox.ColorLightCyan, "_")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerLU, err=\n\t%w", err)
			return
		}
		posRU:= pos.NewPos(2, 5)
		litCornRU, err := lit.NewLit(scr, posRU, termbox.ColorWhite, termbox.ColorYellow, "*")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerRU, err=\n\t%w", err)
			return
		}
		posLD:= pos.NewPos(2, 5)
		litCornLD, err := lit.NewLit(scr, posLD, termbox.ColorWhite, termbox.ColorRed, "x")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerRU, err=\n\t%w", err)
			return
		}
		posRD:= pos.NewPos(2, 5)
		litCornRD, err := lit.NewLit(scr, posRD, termbox.ColorWhite, termbox.ColorRed, "#")
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litCornerRU, err=\n\t%w", err)
			return
		}
		rectangle4, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle4.BorderStyle().LitCornerLUSet(litCornLU)
		rectangle4.BorderStyle().LitCornerRUSet(litCornRU)
		rectangle4.BorderStyle().LitCornerLDSet(litCornLD)
		rectangle4.BorderStyle().LitCornerRDSet(litCornRD)
		rectangle4.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Пятый прямоугольник (с простым стилем)
		rectPos := pos.NewPos(30, 17)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorYellow, termbox.ColorLightGray, alias.ALit(rune(0x02591)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle5, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle5.BorderStyle().Set(cons.RectangleStyleSimple)
		rectangle5.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Шестой прямоугольник (с одиночной линией)
		rectPos := pos.NewPos(55, 10)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorYellow, termbox.ColorLightGray, alias.ALit(rune(0x02591)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle6, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle6.BorderStyle().Set(cons.RectangleStyleSingle)
		rectangle6.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 1)
	}
	{ // Седьмой прямоугольник (с одиночной линией)
		rectPos := pos.NewPos(55, 17)
		rectSize, err := size.NewSize(20, 5)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectSize, err=\n\t%w", err)
			return
		}
		posLit := pos.NewPos(2, 5)
		litFill, err := lit.NewLit(scr, posLit, termbox.ColorYellow, termbox.ColorLightGray, alias.ALit(rune(0x02591)))
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create litFill, err=\n\t%w", err)
			return
		}
		rectangle6, err := rectangle.NewRectangle(scr, rectPos, rectSize, litFill)
		if err != nil {
			logrus.WithError(err).Errorln("rectangle.go/main(): in create rectangle, err=\n\t%w", err)
			return
		}
		rectangle6.BorderStyle().Set(cons.RectangleStyleDouble)
		rectangle6.Redraw()
		scr.Redraw()
		time.Sleep(time.Second * 3)
	}
	scr.Redraw()
	time.Sleep(time.Second * 3)
}
