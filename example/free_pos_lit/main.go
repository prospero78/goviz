// package main -- демонстрация свободного расположения литер
//  Для запуска выполнить:
//    go run ./example/free_pos_lit/main.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/prospero78/goviz/v1"
)

func main() {
	screen, err := goviz.GetScreen()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lit := goviz.NewLit(goviz.Pos{X: 2, Y: 2}, termbox.ColorWhite, termbox.ColorBlue, "#")
	lit.Redraw()
	lit1 := goviz.NewLit(goviz.Pos{X: 4, Y: 3}, termbox.ColorWhite, termbox.ColorDefault, "5")
	lit1.Redraw()
	lit2 := goviz.NewLit(goviz.Pos{X: 6, Y: 4}, termbox.ColorBlue, termbox.ColorDefault, "^")
	lit2.Redraw()
	screen.Redraw()

	time.Sleep(time.Second * 3)
	screen.Close()
}
