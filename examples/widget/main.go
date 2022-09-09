// package main -- показывает базовые свойства виджета
//  Для запуска выполнить:
//    go run ./example/widget/main.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/prospero78/goviz/v1"
	"github.com/prospero78/goviz/v1/screen"
)

func main() {
	screen, err :=screen.GetScreen()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w1 := goviz.NewWidget(2, 2, 20, 5, termbox.ColorRed, termbox.ColorBlue, "")
	w1.IsBorder = true
	w1.Redraw()
	screen.Redraw()

	time.Sleep(time.Second * 10)
	screen.Close()
}
