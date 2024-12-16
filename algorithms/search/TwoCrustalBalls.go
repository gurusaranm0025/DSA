package search

import (
	"fmt"
	"math"
)

func TwoCrustalBalls(breaks []bool) int {
	jmpAmt := int(math.Floor(math.Sqrt(float64(len((breaks))))))
	i := jmpAmt
	for ; i < len(breaks); i += jmpAmt {
		if breaks[i] {
			break
		}
	}

	i -= jmpAmt

	for j := 0; j < jmpAmt && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}

func RunTwoCrystallBalls() {
	breaks := []bool{false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true}

	fmt.Println(TwoCrustalBalls((breaks)))
}
