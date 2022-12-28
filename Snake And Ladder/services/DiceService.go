package service

import (
	"math/rand"
	"time"
)

type DiceService struct{}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (d *DiceService) Roll() int {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, 7)
	return randomNum
}
