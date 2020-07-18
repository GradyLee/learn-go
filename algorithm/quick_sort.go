package algorithm

func QuickSort(v []int, low, high int) {
	if low >= high {
		return
	}

	first := low
	last := high
	base := v[low]

	for first < last {
		for first < last && v[last] >= base {
			last--
		}
		if (first < last) {
			v[first] = v[last]
			first++
		}

		for first < last && v[first] <= base {
			first++
		}
		if (first < last) {
			v[last] = v[first]
			last--
		}
	}
	v[first] = base

	QuickSort(v, low, first - 1)
	QuickSort(v, first + 1, high)
}
