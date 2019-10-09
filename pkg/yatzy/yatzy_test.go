package yatzy_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
	"yatzy/pkg/yatzy"
)

func TestDice_Count(t *testing.T) {
	dice := yatzy.Dice{1, 2, 3, 4, 5, 5, 6, 6, 6}

	assert.Equal(t, 1, dice.Count(4))
	assert.Equal(t, 2, dice.Count(5))
	assert.Equal(t, 3, dice.Count(6))
}

func TestDice_Score(t *testing.T) {
	dice := yatzy.Dice{1, 2, 3, 4, 5, 5, 6, 6, 6}

	assert.Equal(t, yatzy.Score(18), dice.Score(6))
	assert.Equal(t, yatzy.Score(4), dice.Score(4))
}

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

func BenchmarkGame(b *testing.B) {

	var dice yatzy.Dice
	var keep yatzy.Dice
	var err error

	rand.Seed(time.Now().UnixNano())
	game := yatzy.NewGame()

	for i := 0; i < b.N; i++ {
		for err == nil {
			dice = roll(5 - len(keep))
			dice = append(keep, dice...)
			keep, err = game.Roll(dice)
		}
		game.Results()
	}
}

func TestAverage(t *testing.T) {

	var total int

	const max = 10000000
	var dice yatzy.Dice
	var keep yatzy.Dice
	var err error

	rand.Seed(time.Now().UnixNano())
	game := yatzy.NewGame()

	for i := 0; i < max; i++ {
		for err == nil {
			dice = roll(5 - len(keep))
			dice = append(keep, dice...)
			keep, err = game.Roll(dice)
		}
		results := game.Results()
		total += int(results.Total)
	}

	fmt.Printf("Average score is %d\n", total/max)
}
