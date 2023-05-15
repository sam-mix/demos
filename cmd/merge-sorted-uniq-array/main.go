package main

import "fmt"

var (
	a = []int{1}
	b = []int{2}
)

func main() {
	fmt.Println(mergeSortedUniqArray(a, b))
}

// 求两个升序不重复int数组的合集, 是否存在比以下方法更高效的算法
func mergeSortedUniqArray(a, b []int) (result []int) {
	n, m := len(a), len(b)
	for i, j := 0, 0; i < n && j < m; {
		if a[i] < b[j] {
			i++
			continue
		}
		if a[i] > b[j] {
			j++
			continue
		}
		result = append(result, a[i])
		i++
		j++
	}
	return
}
