// package main -- демонстрация свободного расположения текста
//  Для запуска выполнить:
//    go run ./example/free_pos_text/main.go
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

	goviz.OutString(2, 2, termbox.ColorRed, termbox.ColorDefault, "Привет, терминал!")
	goviz.Update()

	time.Sleep(time.Second * 3)
	goviz.Close()
}
