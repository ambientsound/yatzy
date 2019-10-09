package yatzy_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
