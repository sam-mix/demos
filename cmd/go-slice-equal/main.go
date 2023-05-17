package main

import (
	"fmt"
	"reflect"
)

var (
	_ reflect.Kind
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	if reflect.DeepEqual(a, b) {
		fmt.Println("a 和 b 相等")
	} else {
		fmt.Println("a 和 b 不相等")
	}
}
