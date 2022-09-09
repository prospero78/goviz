package sizex

import (
	"testing"
)

/*
	Тест для размера объекта по горизонтали
*/

type tester struct {
	t   *testing.T
	err error
	sz  *SizeX
}

func TestSizeX(t *testing.T) {
	test := &tester{
		t: t,
	}
	test.create()
}

// Создание объекта размера по горизонтали
func (sf *tester) create() {
	sf.t.Logf("=create=")
	sf.createBad1()
	sf.createGood1()
}

func (sf *tester) createGood1() {
	sf.t.Logf("createGood1=")
	sf.sz, sf.err = NewSizeX(25)
	if sf.err != nil {
		sf.t.Errorf("createGood1(): err=%v", sf.err)
	}
	if sf.sz == nil {
		sf.t.Errorf("size==nil")
		return
	}
	if val := sf.sz.Get(); val != 25 {
		sf.t.Errorf("createGood1(): val(%v)!=25", val)
	}
}

// Отрицательное значение размера
func (sf *tester) createBad1() {
	sf.t.Logf("createBad1=")
	sf.sz, sf.err = NewSizeX(-1)
	if sf.err == nil {
		sf.t.Errorf("createBad1(): err==nil")
	}
	if sf.sz != nil {
		sf.t.Errorf("size!=nil")
	}
}
