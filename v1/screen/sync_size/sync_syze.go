// package sync_syze -- потокобезопасный размер
package sync_size

import (
	"fmt"
	"sync"

	"github.com/prospero78/goviz/v1/alias"
	"github.com/prospero78/goviz/v1/screen/sync_size/sync_sizex"
	"github.com/prospero78/goviz/v1/screen/sync_size/sync_sizey"
	"github.com/prospero78/goviz/v1/types"
)

// Syncsize -- потокобезопасный размер
type SyncSize struct {
	sizeX types.ISizeX
	sizeY types.ISizeY
	block sync.RWMutex
}

// NewSyncSize -- возвращает новый потокобезопасный размер
func NewSyncSize(x alias.ASizeX, y alias.ASizeY) (*SyncSize, error) {
	sizeX, err := sync_sizex.NewSyncSizeX(x)
	if err != nil {
		return nil, fmt.Errorf("NewSyncSize(): in create ISyncSizeX, err=\n\t%w", err)
	}
	sizeY, err := sync_sizey.NewSyncSizeY(y)
	if err != nil {
		return nil, fmt.Errorf("NewSyncSize(): in create ISyncSizeX, err=\n\t%w", err)
	}
	sf := &SyncSize{
		sizeX: sizeX,
		sizeY: sizeY,
	}
	return sf, nil
}

// Get -- возвращает размер объекта
func (sf *SyncSize) Get() (alias.ASizeX, alias.ASizeY) {
	sf.block.RLock()
	defer sf.block.RUnlock()
	return sf.sizeX.Get(), sf.sizeY.Get()
}

// Set -- устанавливает размер объекта
func (sf *SyncSize) Set(x alias.ASizeX, y alias.ASizeY) error {
	sf.block.Lock()
	defer sf.block.Unlock()
	if !(x > 0 && y > 0) {
		return fmt.Errorf("Size.Set(): x(%v)<0; y(%v)<0", x, y)
	}
	_ = sf.sizeX.Set(x)
	_ = sf.sizeY.Set(y)
	return nil
}

// SizeX -- устанавливает объект размера объекта по Х
func (sf *SyncSize) SizeX() types.ISizeX {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.sizeX
}

// SizeY -- устанавливает объект размера объекта по Y
func (sf *SyncSize) SizeY() types.ISizeY {
	sf.block.Lock()
	defer sf.block.Unlock()
	return sf.sizeY
}
