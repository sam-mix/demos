package main

import (
	"flag"
	"fmt"
)

func init() {
	env := flag.String("env", "", "环境配置文件路径")
	flag.Parse()
	fmt.Println("*env: ", *env)
}

func main() {

}
