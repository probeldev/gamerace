package model

import "racegame/config"

type Player struct {
	X int
	Y int
}

func (p *Player) Left() {
	if p.X == 0 {
		return
	}
	p.X--
}

func (p *Player) Right() {
	if p.X == config.CountPointX-1 {
		return
	}
	p.X++
}

func (p *Player) Up() {
	if p.Y == 0 {
		return
	}
	p.Y--
}

func (p *Player) Down() {
	if p.Y == config.CountPointY-1 {

		return
	}
	p.Y++
}
