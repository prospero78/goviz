package types

import "github.com/prospero78/goviz/v1/alias"

// IPosX -- интерфейс к положению X
type IPosX interface{
	// Get -- возвращает хранимое положение
	Get()alias.APosX
	// Set -- устанавливает хранимое значение
	Set(alias.APosX)
}

// IPosY -- интерфейс к положению Y
type IPosY interface{
	// Get -- возвращает хранимое положение
	Get()alias.APosY
	// Set -- устанавливает хранимое значение
	Set(alias.APosY)
}

// IPos -- интерфейс позиции на экране
type IPos interface{
	// Get -- возвращает позицию на экране
	Get()(alias.APosX,alias.APosY)
	// Set -- устанавливает позицию на экране
	Set(alias.APosX,alias.APosY)
	// PosX -- возвращает объект позиции по Х
	PosX()IPosX
	// PosY -- возвращает объект позиции по Y
	PosY()IPosY
}