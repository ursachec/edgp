package edgp

import (
	"crypto/rand"
	"errors"
	"io"
	"math"
	"math/big"
	"os"
)

// Thrower holds the source of randomness and the words for generating passphrases
type Thrower struct {
	source       io.Reader
	words        map[int64]string
	dicePerThrow int
	maxDiceValue int
}

// NewThrower does what it says
func NewThrower(source io.Reader, words map[int64]string, dicePerThrow, maxDiceValue int) *Thrower {
	// TODO: this should panic or something if the entries do not cocontain all the permutations required for the dicePerThrow/maxDiceValue combination
	return &Thrower{
		source:       source,
		words:        words,
		dicePerThrow: dicePerThrow,
		maxDiceValue: maxDiceValue,
	}
}

// Throw returns numThrows entries from the Thrower's words
func (g *Thrower) Throw(numThrows int) ([]string, error) {
	chosen := []string{}
	for i := 0; i < numThrows; i++ {
		diceThrows, err := generateThrows(rand.Reader, g.dicePerThrow, g.maxDiceValue)
		if err != nil {
			os.Exit(1)
		}

		entryID := throwsToID(diceThrows)
		if word, ok := g.words[entryID]; ok {
			chosen = append(chosen, word)
		} else {
			// This should be a test, not a damn panic
			panic("Could not find word for dice entry")
		}
	}
	return chosen, nil
}

func generateThrows(randSource io.Reader, dice, maxDiceVal int) ([]int64, error) {
	diceThrows := []int64{}
	for j := 0; j < dice; j++ {
		five := new(big.Int).SetInt64(int64(maxDiceVal))
		res, err := rand.Int(randSource, five)
		if err != nil {
			return []int64{}, errors.New("Could not generate random integer")
		}
		one := new(big.Int).SetInt64(1)
		res.Add(res, one)

		diceThrows = append(diceThrows, res.Int64())
	}
	return diceThrows, nil
}

func throwsToID(throws []int64) int64 {
	final := int64(0)
	for idx, entry := range throws {
		if idx == 0 {
			final += entry
		} else {
			final += int64(math.Pow10(idx)) * entry
		}
	}
	return final
}
