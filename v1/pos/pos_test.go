package pos

import "testing"

/*
	Тест для позиции объекта
*/

func TestPos(t *testing.T) {
	pos := NewPos(5, 6)
	x := pos.PosX().Get()
	if x != 5 {
		t.Errorf("TstPos(): x(%v)!=5", x)
	}
	y := pos.PosY().Get()
	if y != 6 {
		t.Errorf("TstPos(): y(%v)!=6", y)
	}

	pos.Set(7, 8)
	x,y = pos.Get()
	if x != 7 {
		t.Errorf("TstPos(): x(%v)!=7", x)
	}
	if y != 8 {
		t.Errorf("TstPos(): y(%v)!=8", y)
	}
}
