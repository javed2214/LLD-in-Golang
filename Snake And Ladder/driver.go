package main

import (
	"container/list"
	"fmt"

	models "snake-and-ladder/models"
	services "snake-and-ladder/services"
)

func main() {

	var noOfSnakes int
	fmt.Scanf("%d", &noOfSnakes)
	var snakes = list.New()
	for i := 0; i < noOfSnakes; i++ {
		var start, end int
		fmt.Scanf("%d %d", &start, &end)
		snake := models.Snake{Start: start, End: end}
		snakes.PushBack(snake)
	}

	var noOfLadders int
	fmt.Scanf("%d", &noOfLadders)
	var ladders = list.New()
	for i := 0; i < noOfLadders; i++ {
		var start, end int
		fmt.Scanf("%d %d", &start, &end)
		ladder := models.Ladder{Start: start, End: end}
		ladders.PushBack(ladder)
	}

	var noOfPlayers int
	fmt.Scanf("%d", &noOfPlayers)
	var players = list.New()
	for i := 0; i < noOfPlayers; i++ {
		var playerName string
		fmt.Scanf("%s", &playerName)
		player := models.Player{Name: playerName}
		players.PushBack(player)
	}

	SnakeAndLadder := services.SnakeAndLadderService{}

	SnakeAndLadder.SetBoardSize(100)
	SnakeAndLadder.SetSnakes(snakes)
	SnakeAndLadder.SetLadders(ladders)
	SnakeAndLadder.SetPlayers(players)
	SnakeAndLadder.StartGame()
}
