package problems

import "math"

// PROBLEM
// Given two crystal balls that will break from a high enough distance,
// determine the exact spot in which it will break
// in the most optimized way.

// EXPLANATION:
// lets consider the height at which the balls break is denoted as 1 and 0 for the opposite.
// so it can be represented in a boolean array where each each index represents a certain level of height.

// Height := []bool{false, false, false, false, false, true, true, true, true, true}

func CrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount
	for ; i<len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount
	for j:=0; j<jumpAmount && i<len(breaks); i, j = i+1, j+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}