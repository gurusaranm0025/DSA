package main

import (
	"fmt"

	"github.com/gurusaranm0025/DSATP/algorithms/linkedList"
)

func main(){
	// nums := []int{3,7,4,9,12,45,0,65,32,23}
	// sortedNums := []int{0,3,4,7,9,12,23,32,45,65}
	
	// nums = sorting.BubbleSort(nums)
	// i := searching.BinarySearch(sortedNums, 32)
	// i := tricks.RangedSum(10,30)

	list := linkedList.DoublyLinkedList{}
	list.Append(3)
	list.Append(67)
	list.Append("Hi")
		
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))

	list.Remove(3)
	fmt.Println("Remoed")
	fmt.Println(list.Get(0))

	// for i:=0; i<list.GetLength(); i++{
	// 	fmt.Println(list.Get(i))
	// }

}