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
	if err := goviz.New(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	goviz.OutLit(2, 2, termbox.ColorRed, termbox.ColorDefault, []rune("+")[0])
	goviz.OutLit(4, 3, termbox.ColorGreen, termbox.ColorDefault, []rune("#")[0])
	goviz.OutLit(6, 4, termbox.ColorBlue, termbox.ColorDefault, []rune("*")[0])
	goviz.Update()

	time.Sleep(time.Second * 3)
	goviz.Close()
}
