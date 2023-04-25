package main

import (
	"github.com/goplus/igop"
	_ "github.com/goplus/igop/gopbuild"
)

var gopSrc string = `
fields := [
	"engineering",
	"STEM education", 
	"and data science",
]

println "The Go+ language for", fields.join(", ")
`

func main() {
	_, err := igop.RunFile("main.gop", gopSrc, nil, 0)
	if err != nil {
		panic(err)
	}
}
