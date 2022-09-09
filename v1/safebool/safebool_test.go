package safebool

import "testing"

/*
	Тест для безопасного булева признака
*/

// Тестер для потокобезопасного булева признака
type tester struct {
	t  *testing.T
	sb *SafeBool
}

func TestSafeBool(t *testing.T) {
	test := &tester{
		t: t,
	}
	test.create()
	test.set()
	test.reset()
}

// Сбрасывает булев признак
func (sf *tester) reset() {
	sf.t.Logf("=reset=")
	sf.sb.Reset()
	if sf.sb.Get() {
		sf.t.Errorf("reset(): result!=false")
	}
}

// Устанавливает булев признак
func (sf *tester) set() {
	sf.t.Logf("=set=")
	sf.sb.Set()
	if !sf.sb.Get() {
		sf.t.Errorf("set(): result!=true")
	}
}

// Создание булева признака
func (sf *tester) create() {
	sf.t.Logf("=create=")
	sf.sb = NewSafeBool()
	if sf.sb == nil {
		sf.t.Errorf("create(): safeBool==nil")
	}
	if sf.sb.Get() {
		sf.t.Errorf("create(): result!=false")
	}
}
