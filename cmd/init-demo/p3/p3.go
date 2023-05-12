package p3

import "fmt"

var X int = 3

func init() {
	defer func() {
		fmt.Println("defer p3")
	}()
}
