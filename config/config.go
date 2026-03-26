package config

import (
	"golang.org/x/image/font"
)

const (
	PointSize    = 60
	WindowWidth  = 1280
	WindowHeight = 860
)

var GameFont font.Face

type ScreenType int

const (
	ScreenTypeStart    ScreenType = 1
	ScreenTypeGame     ScreenType = 2
	ScreenTypeGameOver ScreenType = 3
)

const (
	CountLine  = 4
	AreaWidth  = PointSize * CountLine
	AreaHeight = WindowHeight

	AreaMarginLeft   = 300
	AreaMarginTop    = 0
	AreaMarginBottom = 50

	AreaBorderSize = 20

	CountPointX = AreaWidth / PointSize
	CountPointY = (AreaHeight - AreaMarginTop - AreaMarginBottom) / PointSize
)

const CountScopeForLevel = 20
