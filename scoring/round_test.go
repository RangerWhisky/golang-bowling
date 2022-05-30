package round

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrike(t *testing.T) {
	t.Fail()
}

func TestTooManyPins(t *testing.T) {
	pins := [2]int{10, 1}
	error := NewRound(pins)
	assert.NotNil(t, error)
}
