package main

import (
	"fmt"

	"github.com/gurusaranm0025/DSATP/algorithms/tricks"
)

func main(){
	// nums := []int{3,7,4,9,12,45,0,65,32,23}
	// sortedNums := []int{0,3,4,7,9,12,23,32,45,65}
	
	// nums = sorting.BubbleSort(nums)
	// i := searching.BinarySearch(sortedNums, 32)
	i := tricks.RangedSum(10,30)
		
	fmt.Println(i)
}