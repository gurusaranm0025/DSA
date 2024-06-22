package searching

func LinearSearch(array []int, target int) int {
	for i:=0; i<len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}