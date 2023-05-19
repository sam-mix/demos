package main

import (
	"demos/core/gamelog"
)

func main() {
	l := gamelog.Get()
	// hello := struct {
	// 	N string
	// 	A int
	// }{
	// 	N: "xx",
	// 	A: 10,
	// }
	// l.Error().Any("hello", hello)

	// time.Sleep(10 * time.Second)
	// l.Err().Str("key", "val")
	l.Error().Str("k", "v")
}
