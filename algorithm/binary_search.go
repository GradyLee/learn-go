package algorithm

func BinarySearch(v []int, value, low, high int) (int) {
	if len(v) <= 0 {
		return -1
	}

	for low < high {
		mid := low + (high - low) / 2
		if v[mid] == value {
			return mid
		} else if v[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearch2(v []int, value, low, high int) (int) {
	if low > high {
		return -1
	}

	mid := low + (high - low) / 2
	if v[mid] == value {
		return mid
	} else if v[mid] > value {
		return BinarySearch2(v, value, low, mid - 1)
	} else {
		return BinarySearch2(v, value, mid + 1 , high)
	}
}

