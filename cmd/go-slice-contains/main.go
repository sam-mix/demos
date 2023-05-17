package main

import (
	"fmt"
	"sort"
)

var (
	_ sort.Interface
)

func main() {
	fmt.Println(containsSample([]string{"", "hello"}, "hello"))
	fmt.Println(contains([]string{"", "hello1"}, "hello"))
}

func containsSample(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func contains(s []string, e string) bool {
	i := sort.SearchStrings(s, e)
	return i < len(s) && s[i] == e
}
