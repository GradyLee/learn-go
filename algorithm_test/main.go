package main

import (
	"fmt"
	"lizq.dev/learn-go/algorithm"
)


func main() {
	var exist algorithm.Status
	v := []int{1, 3, 6, 4, 7, 9, 0, 2}
	x := algorithm.BinarySearch(v, 9, 0, len(v) - 1)
	y := algorithm.BinarySearch2(v, 9, 0, len(v) - 1)
	fmt.Println(x)
	fmt.Println(y)

	var t *algorithm.BiTree
	for _, d := range v {
		t = algorithm.AddBST(t, algorithm.KeyType(d))
	}
	exist = algorithm.SearchBST(t, 9)
	fmt.Println(exist)
	exist = algorithm.SearchBST(t, 10)
	fmt.Println(exist)

	algorithm.BubbleSortOrderly(v)
	fmt.Println(v)

	v2 := []int{1, 3, 6, 4, 7, 9, 0, 2, 11, 32, 18, 56, 22}
	algorithm.BucketSort(v2)
	fmt.Println(v2)

	//可变参数实现
	v3 := algorithm.InsertSort(1, 3, 6, 4, 7, 9, 0, 2, 11, 32, 18, 56, 22)
	fmt.Println(v3)

	v4 := algorithm.SelectionSort(1, 3, 6, 4, 7, 9, 0, 2, 11, 32, 18, 56, 22)
	fmt.Println(v4)

}