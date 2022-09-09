package lit

import (
	"testing"

	"github.com/nsf/termbox-go"
	"github.com/prospero78/goviz/v1/pos"
)

/*
	Тест для литеры
*/

type tester struct {
	t   *testing.T
	err error
	lit *Lit
}

var (
	_pos = pos.NewPos(1, 1)
	fa   = termbox.ColorGreen
	ba   = termbox.ColorBlue
)

func TestLit(t *testing.T) {
	test := &tester{
		t: t,
	}
	test.create()
}

// Создаёт новую литеру
func (sf *tester) create() {
	sf.t.Logf("=create=")
	sf.createBad1()
}

// Нет объекта экрана
func (sf *tester) createBad1() {
	sf.t.Logf("=createBad1=")
	sf.lit, sf.err = NewLit(nil, _pos, fa, ba, " ")
}
