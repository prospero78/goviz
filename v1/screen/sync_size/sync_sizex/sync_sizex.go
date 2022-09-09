// package sync_sizex -- потокобезопасный размер объекта по оси Х
//  Размер не может быть меньше нуля
package sync_sizex

import (
	"fmt"
	"sync"

	"github.com/prospero78/goviz/v1/alias"
)

// SyncSizeX -- поткобезопасный размер объекта по оси X
type SyncSizeX struct {
	val   alias.ASizeX
	block sync.RWMutex
}

// NewSyncSizeX -- возвращает новый размер по горизонтали
func NewSyncSizeX(x alias.ASizeX) (*SyncSizeX, error) {
	sf := &SyncSizeX{}
	if err := sf.Set(x); err != nil {
		return nil, fmt.Errorf("NewSyncSizeX(): in set size, err=\n\t%w", err)
	}
	return sf, nil
}

// Get -- возвращает хранимое значение
func (sf *SyncSizeX) Get() alias.ASizeX {
	sf.block.RLock()
	defer sf.block.RUnlock()
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *SyncSizeX) Set(val alias.ASizeX) error {
	sf.block.Lock()
	defer sf.block.Unlock()
	if val < 0 {
		return fmt.Errorf("SyncSizeX.Set(): x(%v)<0", val)
	}
	sf.val = val
	return nil
}
