package safebool

import "sync"

// SafeBool -- потокобезопасный признак
type SafeBool struct {
	val   bool // Хранимое значение признака
	block sync.RWMutex
}

// NewSafeBool -- возвращает новый потокобезопасный признак
func NewSafeBool() *SafeBool {
	return &SafeBool{}
}

// Get -- возвращает хранимый признак
func (sf *SafeBool) Get() bool {
	sf.block.RLock()
	defer sf.block.RUnlock()
	return sf.val
}

// Set -- устанавливает хранимый признак
func (sf *SafeBool) Set() {
	sf.block.Lock()
	defer sf.block.Unlock()
	sf.val = true
}

// Reset -- сбрасывает хранимый признак
func (sf *SafeBool) Reset() {
	sf.block.Lock()
	defer sf.block.Unlock()
	sf.val = false
}
