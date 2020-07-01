package algorithm

func BubbleSort(v []int) {
	len := len(v)
	for i := 0; i < len - 1; i++ {
		for j := 0; j < len - 1 - i; j++ {
			if v[j] > v[j + 1] {
				v[j], v[j + 1] = v[j + 1], v[j]
			}
		}
	}
}


func BubbleSortOrderly(v []int) {
	orderly := false
	len := len(v)

	for i := 0; i < len - 1 && orderly == false; i++ {
		orderly = true
		for j := 0; j < len - 1 - i; j++ {
			if v[j] > v[j + 1] {
				v[j], v[j + 1] = v[j + 1], v[j]
				orderly = false
			}
		}
	}
}