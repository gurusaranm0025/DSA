package searching

func BinarySearch(array []int, target int) int {
	low := 0
	high := len(array)

	for low < high {
		mid := (low+high)/2
		midVal := array[mid]

		if midVal == target {
			return mid
		}

		if midVal > target {
			high = mid
		}

		if midVal < target {
			low = mid
		}
	}

	return -1
}