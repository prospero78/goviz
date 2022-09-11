// package main -- демонстрация свободного расположения строки
//
//	Для запуска выполнить:
//	  go run ./example/free_pos_string/main.go
package main

import (
	"time"

	// "github.com/nsf/termbox-go"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/attribute"
	"github.com/prospero78/goviz/v1/pos"
	"github.com/prospero78/goviz/v1/screen"
	"github.com/prospero78/goviz/v1/tstring"
	"github.com/sirupsen/logrus"
)

func main() {
	screen, err := screen.GetScreen()
	if err != nil {
		logrus.WithError(err).Errorln("free_pos_string: in get screen")
		return
	}
	defer screen.Close()

	{ // Первая строка
		_pos := pos.NewPos(2, 5)
		foreAttr := attribute.NewAttribute(termbox.ColorYellow)
		foreAttr.Blink().Set()
		str, err := tstring.NewTString(screen, _pos, foreAttr.Get(), termbox.ColorDefault, "Привет всем любителям пошкодить!")
		if err != nil {
			logrus.WithError(err).Errorln("free_pos_string: in create TString")
			return
		}
		str.Redraw()
	}
	{ // Вторая строка
		_pos := pos.NewPos(2, 7)
		foreAttr := attribute.NewAttribute(termbox.ColorLightGray)
		foreAttr.Blink().Set()
		foreAttr.Underline().Set()
		str, err := tstring.NewTString(screen, _pos, foreAttr.Get(), termbox.ColorDefault, "Привет всем любителям пошкодить и отладить!")
		if err != nil {
			logrus.WithError(err).Errorln("free_pos_string: in create TString")
			return
		}
		str.Redraw()
	}
	{ // Третья строка
		_pos := pos.NewPos(2, 9)
		foreAttr := attribute.NewAttribute(termbox.ColorBlue)
		foreAttr.Blink().Set()
		foreAttr.Underline().Set()
		str, err := tstring.NewTString(screen, _pos, foreAttr.Get(), termbox.ColorDefault, "Привет всем любителям пошкодить, отладить и погамать в WarZone2100!")
		if err != nil {
			logrus.WithError(err).Errorln("free_pos_string: in create TString")
			return
		}
		str.Redraw()
	}
	screen.Redraw()

	time.Sleep(time.Second * 3)
}
