package tricks

// import "fmt"

// sum of numbers between a range of numbers
// NOTE: this takes both the starting and the ending numbers you give into consideration
func RangedSum(startingNum int, endingNum int) int {
	if startingNum == 0 {
		return (endingNum+1)*endingNum/2
	}
	return (startingNum+endingNum)*(endingNum-startingNum)/2

}