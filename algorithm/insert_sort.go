package algorithm


func InsertSort(v ...int) ([]int){
	length:= len(v)
	var j int

	if length <= 1 {
		return v
	}

	for i := 1; i < length; i++ {
		temp := v[i]
		for j = i - 1; j >= 0; j-- {
			if v[j] > temp {
				v[j + 1] = v[j]
			} else {
				break
			}
		}
		v[j + 1] = temp
	}

	return v
}