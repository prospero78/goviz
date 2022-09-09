package size

import "testing"

/*
	Тест для размера объекта
*/

type tester struct {
	t   *testing.T
	err error
	sz  *Size
}

func TestSize(t *testing.T) {
	test := &tester{
		t: t,
	}
	test.create()
	test.getX()
	test.getY()
	test.set()
}

// Установка размеров
func (sf *tester) set() {
	sf.t.Logf("=set=")
	sf.setBad1()
	sf.setBad2()
	sf.setGood1()
}

func (sf *tester) setGood1() {
	sf.t.Logf("=setGood1=")
	if err := sf.sz.Set(7, 8); err != nil {
		sf.t.Errorf("setGood1(): err=%v", err)
	}
	x := sf.sz.SizeX().Get()
	if x != 7 {
		sf.t.Errorf("setBad1(): x(%v)!=7", x)
	}
	y := sf.sz.SizeY().Get()
	if y != 8 {
		sf.t.Errorf("setGood1(): y(%v)!=8", y)
	}
}

// Отрицательное значение высоты
func (sf *tester) setBad2() {
	sf.t.Logf("=setBad2=")
	if err := sf.sz.Set(0, -1); err == nil {
		sf.t.Errorf("setBad1(): err==nil")
	}
	y := sf.sz.SizeY().Get()
	if y != 6 {
		sf.t.Errorf("setBad2(): y(%v)!=6", y)
	}
}

// Отрицательное значение ширины
func (sf *tester) setBad1() {
	sf.t.Logf("=setBad1=")
	if err := sf.sz.Set(-1, 0); err == nil {
		sf.t.Errorf("setBad1(): err==nil")
	}
	x := sf.sz.SizeX().Get()
	if x != 5 {
		sf.t.Errorf("setBad1(): x(%v)!=5", x)
	}
}

// получение объекта высоты
func (sf *tester) getY() {
	sf.t.Logf("=getY=")
	szY := sf.sz.SizeY()
	if szY == nil {
		sf.t.Errorf("getY(): sizeY == nil")
	}
}

// получение объекта ширины
func (sf *tester) getX() {
	sf.t.Logf("=getX=")
	szX := sf.sz.SizeX()
	if szX == nil {
		sf.t.Errorf("getX(): sizeX == nil")
	}
}

// Создание размера
func (sf *tester) create() {
	sf.t.Logf("=create=")
	sf.createBad1()
	sf.createBad2()
	sf.createGood1()
}

func (sf *tester) createGood1() {
	sf.t.Logf("=createGood1=")
	sf.sz, sf.err = NewSize(5, 6)
	if sf.err != nil {
		sf.t.Errorf("createGood1(): err=%v", sf.err)
	}
	if sf.sz == nil {
		sf.t.Errorf("createGood1(): size==nil")
		return
	}
	x, y := sf.sz.Get()
	if x != 5 {
		sf.t.Errorf("createGood1(): x(%v)!=5", x)
	}
	if y != 6 {
		sf.t.Errorf("createGood1(): y(%v)!=6", y)
	}
}

// отрицательная высота
func (sf *tester) createBad2() {
	sf.t.Logf("=createBad2=")
	sf.sz, sf.err = NewSize(0, -1)
	if sf.err == nil {
		sf.t.Errorf("createBad2(): err==nil")
	}
	if sf.sz != nil {
		sf.t.Errorf("createBad2(): size!=nil")
	}
}

// отрицательная ширина
func (sf *tester) createBad1() {
	sf.t.Logf("=createBad1=")
	sf.sz, sf.err = NewSize(-1, 0)
	if sf.err == nil {
		sf.t.Errorf("createBad1(): err==nil")
	}
	if sf.sz != nil {
		sf.t.Errorf("createBad1(): size!=nil")
	}
}
