package yatzy

import (
	"errors"
)

type Die int

type Dice []Die

func (dice Dice) Count(value Die) (count int) {
	for _, die := range dice {
		if die == value {
			count++
		}
	}
	return
}

func (dice Dice) Score(value Die) Score {
	return Score(dice.Count(value)) * Score(value)
}

func (dice Dice) Only(value Die) Dice {
	only := Dice{}
	for _, die := range dice {
		if die == value {
			only = append(only, die)
		}
	}
	return only
}

func (dice Dice) Without(value Die) Dice {
	without := Dice{}
	for _, die := range dice {
		if die != value {
			without = append(without, die)
		}
	}
	return without
}

type Score int

type Scorecard map[Die]*Score

type Results struct {
	Scorecard Scorecard `json:"scorecard"`
	Total     Score     `json:"total"`
}

var (
	ErrFinished = errors.New("game is finished")
)

func (scorecard Scorecard) Finished() bool {
	for _, v := range scorecard {
		if v == nil {
			return false
		}
	}
	return true
}

func (scorecard Scorecard) Total() (total Score) {
	for _, score := range scorecard {
		total += *score
	}
	return
}

type Game struct {
	scorecard Scorecard
	throws    int
	keep      Dice
	dice      Dice
}

func (game *Game) Roll(dice Dice) (keep Dice, err error) {
	if game.scorecard.Finished() {
		return nil, ErrFinished
	}

	game.throws += 1
	game.dice = dice

	bestFit := game.BestFit()
	// fmt.Printf("best fit is %d\n", bestFit)

	game.keep = game.dice.Only(bestFit)
	game.dice = game.dice.Without(bestFit)

	// Reached max number of throws; reset state
	if game.throws == 3 {
		score := game.keep.Score(bestFit)
		game.scorecard[bestFit] = &score
		// fmt.Printf("scored %d points on %d\n", score, bestFit)
		game.next()
	}

	return game.keep, nil
}

func (game *Game) next() {
	game.throws = 0
	game.keep = Dice{}
	game.dice = Dice{}
}

// NaiveFit doesn't care about the dice, just tries to finish the game
func (game *Game) NaiveFit() Die {
	for i, score := range game.scorecard {
		if score == nil {
			return Die(i)
		}
	}
	return 0
}

// BestFit decides which dies to keep and takes into account the scoreboard
// and the dice on hand.
func (game *Game) BestFit() (bestFit Die) {
	var bestCount int

	for i, score := range game.scorecard {
		if score != nil {
			continue
		}
		count := game.dice.Count(i)
		if bestCount == 0 || count > bestCount || (count == bestCount && i > bestFit) {
			bestCount = count
			bestFit = i
		}
	}

	return
}

func (game *Game) Results() Results {
	return Results{
		Scorecard: game.scorecard,
		Total:     game.scorecard.Total(),
	}
}

func NewScorecard() Scorecard {
	var i Die
	scorecard := Scorecard{}
	for i = 1; i <= 6; i++ {
		scorecard[i] = nil
	}
	return scorecard
}

func NewGame() *Game {
	return &Game{
		scorecard: NewScorecard(),
	}
}
