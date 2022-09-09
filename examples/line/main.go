// package main -- демонстрация свободного расположения линии
//  Для запуска выполнить:
//    go run ./example/line/main.go
package main

import (
	"time"

	// "github.com/nsf/termbox-go"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/line"
	"github.com/prospero78/goviz/v1/lit"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/screen"
	"github.com/sirupsen/logrus"
)

func main() {
	screen, err := screen.GetScreen()
	if err != nil {
		logrus.WithError(err).Errorln("line: in get screen")
		return
	}
	defer screen.Close()

	{ // Первая линия
		_posLine := pos.NewPos(2, 2)
		_posLit := pos.NewPos(2, 2)
		lit0, err := lit.NewLit(screen, _posLit, termbox.ColorLightYellow, termbox.ColorBlue, "")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		line0, err := line.NewLine(screen, _posLine, 10, lit0)
		if err != nil {
			logrus.WithError(err).Errorln("line: in create TString")
			return
		}
		line0.Redraw()
	}
	{ // Вторая линия
		_posLine2 := pos.NewPos(2, 4)
		_posLit2 := pos.NewPos(2, 8)
		lit0, err := lit.NewLit(screen, _posLit2, termbox.ColorLightYellow, termbox.ColorCyan, "")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		line0, err := line.NewLine(screen, _posLine2, 12, lit0)
		if err != nil {
			logrus.WithError(err).Errorln("line: in create TString")
			return
		}
		litBeg, err := lit.NewLit(screen, _posLit2, termbox.ColorWhite, termbox.ColorYellow, " ")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		litEnd, err := lit.NewLit(screen, _posLit2, termbox.ColorRed, termbox.ColorYellow, " ")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		line0.LitBegin = litBeg
		line0.LitEnd = litEnd
		line0.Redraw()
	}
	{ // Третья линия
		_posLine3 := pos.NewPos(2, 6)
		_posLit3 := pos.NewPos(2, 8)
		lit0, err := lit.NewLit(screen, _posLit3, termbox.ColorLightYellow, termbox.ColorCyan, "-")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		line0, err := line.NewLine(screen, _posLine3, 14, lit0)
		if err != nil {
			logrus.WithError(err).Errorln("line: in create TString")
			return
		}
		litBeg, err := lit.NewLit(screen, _posLit3, termbox.ColorWhite, termbox.ColorYellow, "o")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		litEnd, err := lit.NewLit(screen, _posLit3, termbox.ColorRed, termbox.ColorYellow, "x")
		if err != nil {
			logrus.WithError(err).Errorln("line: in create ILit")
			return
		}
		line0.LitBegin = litBeg
		line0.LitEnd = litEnd
		line0.Redraw()
	}
	screen.Redraw()

	time.Sleep(time.Second * 7)
}
