package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"yatzy/pkg/yatzy"
)

func random() yatzy.Die {
	return yatzy.Die(rand.Intn(6) + 1)
}

func roll(n int) yatzy.Dice {
	dice := make(yatzy.Dice, n)
	for i := range dice {
		dice[i] = random()
	}
	return dice
}

func main() {
	var dice yatzy.Dice
	var keep yatzy.Dice
	var err error

	rand.Seed(time.Now().UnixNano())
	game := yatzy.NewGame()

	for err == nil {
		dice = roll(5 - len(keep))
		// fmt.Printf("rolled: %+v\n", dice)
		dice = append(keep, dice...)
		// fmt.Printf("next hand is: %+v\n", dice)
		keep, err = game.Roll(dice)
		// fmt.Printf("keeping: %+v\n", keep)
	}

	results := game.Results()

	data, err := json.MarshalIndent(results, "", "    ")
	fmt.Println(string(data))
}
