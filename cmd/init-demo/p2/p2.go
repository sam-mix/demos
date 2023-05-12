package p2

import (
	"flag"
	"fmt"
)

var X int = 2

func init() {
	port := flag.String("port", "", "环境配置文件路径")
	flag.Parse()
	fmt.Println("*port: ", *port)
	// defer func() {
	// 	fmt.Println("defer p2")
	// }()
}
