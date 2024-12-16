package sort

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

func RunBubbleSort() {
	array := []int{12, 5, 6, 56, 89, 821, 6, 56, 23, 45}

	fmt.Println(BubbleSort(array))
}
