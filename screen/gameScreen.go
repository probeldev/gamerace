package screen

import (
	"image/color"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/probeldev/gamerace/config"
	"github.com/probeldev/gamerace/model"
	"github.com/probeldev/gamerace/scope"
)

type gameScreen struct {
	Player model.Player
	Cars   []model.Car
	timer  int
	Coins  []model.Player
	scope  *scope.Scope

	changeScreenFunc func(config.ScreenType)
}

func NewGameScreen(
	changeScreenFunc func(config.ScreenType),
	scope *scope.Scope,
) *gameScreen {
	gs := &gameScreen{}

	gs.Player.X = 1
	gs.Player.Y = config.CountPointY - 1

	gs.changeScreenFunc = changeScreenFunc

	gs.scope = scope
	return gs
}

func (gs *gameScreen) addNewCar() {

	minX := 0
	maxX := config.CountPointX - 1

	x := rand.Intn(maxX-minX+1) + minX

	gs.Cars = append(gs.Cars, model.Car{
		Y:       -1,
		VisualY: -1,
		X:       x,
	})
}

func (gs *gameScreen) deleteCars() {

	cars := []model.Car{}
	for _, c := range gs.Cars {
		if c.Y > config.CountPointY {
			gs.scope.Value++
		} else {
			cars = append(cars, c)
		}
	}

	gs.Cars = cars
}

func (gs *gameScreen) Update() error {

	gs.moveCars()

	if gs.needsToAddCar() {
		gs.addNewCar()
	}

	if gs.isStopGame() {
		gs.changeScreenFunc(config.ScreenTypeGameOver)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyH) {
		gs.Player.Left()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyL) {
		gs.Player.Right()
	}

	gs.deleteCars()

	gs.timer++
	return nil
}

func (gs *gameScreen) Draw(
	screenH *ebiten.Image,
) {
	// Обычный HUD с нормальным шрифтом

	screenH.Fill(color.RGBA{0x00, 0x33, 0x00, 0xFF})

	pointSize := float32(config.PointSize)

	for _, c := range gs.Cars {
		vector.FillRect(screenH, float32(c.X)*pointSize+config.AreaMarginLeft, float32(c.VisualY)*pointSize+config.AreaMarginTop, pointSize, pointSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)
	}

	vector.FillRect(screenH, float32(gs.Player.X)*pointSize+config.AreaMarginLeft, float32(gs.Player.Y)*pointSize+config.AreaMarginTop, pointSize, pointSize, color.RGBA{0x00, 0xFF, 0x00, 0xff}, false)

	scoreText := "Score: " + strconv.Itoa(gs.scope.Value)
	text.Draw(screenH, scoreText, config.GameFont, 10, 30, color.White)

	levelText := "Level: " + strconv.Itoa(gs.getCurrentLevel())
	text.Draw(screenH, levelText, config.GameFont, 10, 90, color.White)

	// border area
	vector.FillRect(screenH, config.AreaMarginLeft-config.AreaBorderSize, 0, config.AreaBorderSize, config.AreaHeight, color.RGBA{0xFF, 0xFF, 0xFF, 0xff}, false)
	vector.FillRect(screenH, config.AreaMarginLeft+config.AreaWidth, 0, config.AreaBorderSize, config.AreaHeight, color.RGBA{0xFF, 0xFF, 0xFF, 0xff}, false)
}

func (gs *gameScreen) getCurrentLevel() int {
	level := gs.scope.Value/config.CountScopeForLevel + 1
	if level == 0 {
		level = 1
	}
	return level
}

func (gs *gameScreen) getCurrentMoveTime() int {
	switch gs.getCurrentLevel() {
	case 1:
		return 20
	case 2:
		return 18
	case 3:
		return 16
	case 4:
		return 14
	case 5:
		return 12
	}

	return 10
}

func (gs *gameScreen) needsToAddCar() bool {
	return gs.timer%(gs.getCurrentMoveTime()*2) == 0
}

func (gs *gameScreen) isStopGame() bool {

	for _, c := range gs.Cars {
		deltaX := gs.Player.X - c.X
		deltaY := gs.Player.Y - c.Y

		if deltaX == 0 && deltaY == 0 {
			return true
		}

	}

	return false
}

func (gs *gameScreen) moveCars() {
	for i := range gs.Cars {
		gs.Cars[i].Down(gs.getCurrentMoveTime())
	}

}
