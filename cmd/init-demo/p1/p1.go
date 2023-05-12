package p1

import (
	"flag"
	"fmt"
)

var X int = 1

func init() {
	env := flag.String("env", "", "环境配置文件路径")
	flag.Parse()
	fmt.Println("*env: ", *env)
	// defer func() {
	// 	fmt.Println("defer p1")
	// }()
}
