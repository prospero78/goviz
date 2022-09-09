package pos

import "testing"

/*
	Тест для позиции объекта
*/

func TestPos(t *testing.T) {
	pos := NewPos(5, 6)
	if pos.X != 5 {
		t.Errorf("TstPos(): x(%v)!=5", pos.X)
	}
	if pos.Y != 6 {
		t.Errorf("TstPos(): y(%v)!=6", pos.Y)
	}
}
