package round

import "fmt"

func NewRound(pins [2]int) error {
	return fmt.Errorf("Too many pins in this round %d,%d", pins[0], pins[1])
}
