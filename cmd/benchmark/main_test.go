package main

import (
	"math/rand"
	"testing"
)

// 生成一个包含1000个随机整数的切片
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

// 对切片进行冒泡排序
func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

// 基准测试函数
func BenchmarkBubbleSort1000(b *testing.B) {
	slice := generateRandomSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubbleSort(slice)
	}
}

func main() {}
