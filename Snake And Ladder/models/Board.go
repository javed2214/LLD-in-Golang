package models

import "container/list"

type Board struct {
	Size         int
	Snakes       *list.List
	Ladders      *list.List
	Players      *list.List
	PlayerPieces map[string]int
}

func (b *Board) GetSize() int {
	return b.Size
}

func (b *Board) SetSize(size int) {
	b.Size = size
}

func (b *Board) GetSnakes() *list.List {
	return b.Snakes
}

func (b *Board) SetSnakes(snakes *list.List) {
	b.Snakes = snakes
}

func (b *Board) GetLadders() *list.List {
	return b.Ladders
}

func (b *Board) SetLadders(ladders *list.List) {
	b.Ladders = ladders
}

func (b *Board) GetPlayers() *list.List {
	return b.Players
}

func (b *Board) SetPlayers(players *list.List) {
	b.Players = players
}

func (b *Board) GetPlayerPieces() map[string]int {
	return b.PlayerPieces
}

func (b *Board) SetPlayerPieces(playerPieces map[string]int) {
	b.PlayerPieces = playerPieces
}
