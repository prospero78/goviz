package posx

import (
	"testing"
)

/*
	Тест для положения объекта по горизонтали
*/

type tester struct {
	t   *testing.T
	sz  *PosX
}

func TestSizeX(t *testing.T) {
	test := &tester{
		t: t,
	}
	test.create()
	test.set()
}

// Установка положения
func (sf *tester) set() {
	sf.t.Logf("=set=")
	sf.sz.Set(12)
	if val := sf.sz.Get(); val != 12 {
		sf.t.Errorf("set(): val(%v)!=12", val)
	}
}

// Создание объекта положения по горизонтали
func (sf *tester) create() {
	sf.t.Logf("=create=")
	sf.createGood1()
}

func (sf *tester) createGood1() {
	sf.t.Logf("createGood1=")
	sf.sz = NewPosX(25)
	if sf.sz == nil {
		sf.t.Errorf("size==nil")
		return
	}
	if val := sf.sz.Get(); val != 25 {
		sf.t.Errorf("createGood1(): val(%v)!=25", val)
	}
}
