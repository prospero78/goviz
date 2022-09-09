// package main -- демонстрация свободного расположения строки
//  Для запуска выполнить:
//    go run ./example/free_pos_string/main.go
package main

import (
	"time"

	// "github.com/nsf/termbox-go"

	"github.com/nsf/termbox-go"
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

	_pos:=pos.NewPos(2,5)
	str,err:=tstring.NewTString(screen, _pos, termbox.ColorYellow, termbox.ColorDefault, "Привет всем любителям пошкодить!")
	if err!=nil{
		logrus.WithError(err).Errorln("free_pos_string: in create TString")
		return
	}
	str.Redraw()
	screen.Redraw()

	time.Sleep(time.Second * 3)
}
