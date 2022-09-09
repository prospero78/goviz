package types

import "github.com/prospero78/goviz/v1/alias"


// ISizeX -- интерфейс к объекту размера по X
type ISizeX interface{
	// Get -- возвращает объект размера по X
	Get()alias.ASizeX
	// Set -- устанавливает объект размера по X
	Set(alias.ASizeX)error
}

// ISizeY -- интерфейс к объекту размера по Y
type ISizeY interface{
	// Get -- возвращает объект размера по Y
	Get()alias.ASizeY
	// Set -- устанавливает объект размера по Y
	Set(alias.ASizeY)error
}

// ISize -- интерфейс к размеру объекта
type ISize interface{
	// Get -- возвращает размеры по X, Y
	Get()(alias.ASizeX, alias.ASizeY)
	// Set -- устанавливает размеры объекта
	Set(alias.ASizeX, alias.ASizeY)error
	// SizeX -- возвращает объект размера по X
	SizeX()ISizeX
	// SizeY -- возвращает объект размера по Y
	SizeY()ISizeY
}