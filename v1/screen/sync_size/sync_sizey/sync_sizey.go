// package sync_sizey -- потокобезопасный размер объекта по оси Y
//  Размер не может быть меньше нуля
package sync_sizey

import (
	"fmt"
	"sync"

	"github.com/prospero78/goviz/v1/alias"
)

// SyncSizeY -- поткобезопасный размер объекта по оси Y
type SyncSizeY struct {
	val   alias.ASizeY
	block sync.RWMutex
}

// NewSyncSizeY -- возвращает новый размер по горизонтали
func NewSyncSizeY(Y alias.ASizeY) (*SyncSizeY, error) {
	sf := &SyncSizeY{}
	if err := sf.Set(Y); err != nil {
		return nil, fmt.Errorf("NewSyncSizeY(): in set size, err=\n\t%w", err)
	}
	return sf, nil
}

// Get -- возвращает Yранимое значение
func (sf *SyncSizeY) Get() alias.ASizeY {
	sf.block.RLock()
	defer sf.block.RUnlock()
	return sf.val
}

// Set -- устанавливает Yранимое значение
func (sf *SyncSizeY) Set(val alias.ASizeY) error {
	sf.block.Lock()
	defer sf.block.Unlock()
	if val < 0 {
		return fmt.Errorf("SyncSizeY.Set(): Y(%v)<0", val)
	}
	sf.val = val
	return nil
}
