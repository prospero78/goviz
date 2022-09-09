package screen

import (
	"testing"

	"github.com/prospero78/goviz/v1/types"
)

/*
	Тест для экрана
*/

type tester struct {
	t   *testing.T
	err error
	scr types.IScreen
}

func TestScreen(t *testing.T) {
	test := &tester{
		t: t,
	}
	test.create()
}

// Создание экрана
func (sf *tester) create() {
	sf.t.Logf("=create=")
	sf.scr, sf.err = GetScreen()
	if sf.err == nil {
		sf.t.Errorf("TestScreen(): err==nil")
	}
	sf.scr, sf.err = GetScreen()
	if sf.err == nil {
		sf.t.Errorf("TestScreen(): err==nil")
	}
}
