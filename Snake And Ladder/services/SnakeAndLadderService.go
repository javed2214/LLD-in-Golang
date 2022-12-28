package service

import (
	"container/list"
	"fmt"
	"snake-and-ladder/models"
)

type SnakeAndLadderService struct {
	SnakeAndLadderBoard models.Board
	NoOfPlayers         int
	Players             list.List
	BoardSize           int
}

func (sal *SnakeAndLadderService) SetBoardSize(boardSize int) {
	sal.BoardSize = boardSize
}

func (sal *SnakeAndLadderService) SetSnakes(snakes *list.List) {
	sal.SnakeAndLadderBoard.SetSnakes(snakes)
}

func (sal *SnakeAndLadderService) SetLadders(ladders *list.List) {
	sal.SnakeAndLadderBoard.SetLadders(ladders)
}

func (sal *SnakeAndLadderService) SetPlayers(players *list.List) {
	sal.NoOfPlayers = players.Len()
	playerPieces := make(map[string]int, sal.NoOfPlayers)
	for player := players.Front(); player != nil; player = player.Next() {
		sal.Players.PushBack(player.Value)
		playerPieces[player.Value.(models.Player).Name] = 0
	}
	sal.SnakeAndLadderBoard.SetPlayerPieces(playerPieces)
}

func (sal *SnakeAndLadderService) StartGame() {
	for !sal.isGameEnded() {
		diceValue := sal.RollDice().Roll()
		currPlayer := sal.Players.Front()
		sal.Players.Remove(currPlayer)
		sal.movePosition(diceValue, currPlayer.Value.(models.Player))
		if sal.isPlayerWon(currPlayer.Value.(models.Player)) {
			fmt.Println("Player -", currPlayer.Value.(models.Player).Name, " won the game!!!")
		} else {
			sal.Players.PushBack(currPlayer.Value)
		}
	}
}

func (sal *SnakeAndLadderService) RollDice() *DiceService {
	rollDice := new(DiceService)
	return rollDice
}

func (sal *SnakeAndLadderService) isGameEnded() bool {
	return sal.Players.Len() < sal.NoOfPlayers
}

func (sal *SnakeAndLadderService) movePosition(diceValue int, currPlayer models.Player) {
	currPosition := sal.SnakeAndLadderBoard.GetPlayerPieces()[currPlayer.GetPlayerName()]
	nextPosition := currPosition + diceValue
	if nextPosition > sal.BoardSize {
		nextPosition = currPosition
	} else {
		nextPosition = sal.getNewPositionWithSnakeAndLadders(nextPosition)
	}
	sal.SnakeAndLadderBoard.PlayerPieces[currPlayer.GetPlayerName()] = nextPosition
	fmt.Println("Player -", currPlayer.GetPlayerName(), "moved from position", currPosition, "to", nextPosition, "| Dice Value -", diceValue)
}

func (sal *SnakeAndLadderService) getNewPositionWithSnakeAndLadders(nextPosition int) int {
	for snake := sal.SnakeAndLadderBoard.GetSnakes().Front(); snake != nil; snake = snake.Next() {
		if snake.Value.(models.Snake).Start == nextPosition {
			nextPosition = snake.Value.(models.Snake).End
			return nextPosition
		}
	}
	for ladder := sal.SnakeAndLadderBoard.GetLadders().Front(); ladder != nil; ladder = ladder.Next() {
		if ladder.Value.(models.Ladder).Start == nextPosition {
			nextPosition = ladder.Value.(models.Ladder).End
			return nextPosition
		}
	}
	return nextPosition
}

func (sal *SnakeAndLadderService) isPlayerWon(player models.Player) bool {
	return sal.SnakeAndLadderBoard.PlayerPieces[player.GetPlayerName()] == sal.BoardSize
}
