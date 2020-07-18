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


func merge(intervals [][]int) [][]int {
	compare := func(s1, s2, []int) int {
		if s1[0] < s2[0] {
			return -1
		} else if s1[0] > s2[0] {
			return 1
		} else {
			if s1[1] < s2[1] {
				return -1
			} else if s1[1] > s2[1] {
				return 1
			} else {
				return 0
			}
		}
	}

	min := func(v ...int) int {
		low := v[0]
		for _, x := range v {
			if x < low {
				low = x
			}
		}
		return low
	}

	max := func(v ...int) int {
		high := v[0]
		for _, x := range v {
			if x > high {
				high = x
			}
		}
		return high
	}

	mergesec := func (s1, s2, []int) []int{
		return []int{min(s1...), max(s2...)}
	}

	isoverlap := func(s1, s2, []int) bool {
		p1 := float64(s1[0] + s1[1]) / 2.0
		r1 := float64(s1[1] - s1[0]) / 2.0
		p2 := float64(s2[0] + s2[1]) / 2.0
		r2 := float64(s2[1] - s2[0]) / 2.0

		if math.Abs(p2 - p1) > r1 + r2 {
			return false
		}
		return true
	}

	var output [][]int

	if len(intervals) == 1 {
		return intervals
	}

	if len(intervals) == 0 {
		return nil
	}

	sec0 := []int{intervals[0][0], intervals[0][1]}
	output := append(output, sec0)

	for i:= 1; i < len(intervals); i++ {
		sec := intervals[i]
		for j := len(output) - 1; j >= 0; j-- {
			sec2 := intervals[i]
			if isoverlap(sec, sec2) {
				sec2 = mergesec(sec, sec2)
				break
			} else {
				output := append(output, sec)
			}
		}
	}

	return output
}