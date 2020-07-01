package algorithm

func SelectionSort(v ...int) ([]int) {
	length := len(v)

	if length <= 1 {
		return v
	}

	for i := 0; i < length - 1; i++ {
		minpos := i
		for j := i + 1; j < length; j++ {
			if v[j] <= v[minpos] {
				minpos = j
			}
		}
		v[i], v[minpos] = v[minpos], v[i]
	}

	return v
}