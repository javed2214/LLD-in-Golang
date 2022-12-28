package models

type Player struct {
	Name string
}

func NewPlayer() Player {
	return Player{}
}

func (p *Player) GetPlayerName() string {
	return p.Name
}

func (p *Player) SetPlayerName(name string) {
	p.Name = name
}
