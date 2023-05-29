package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	<-timer.C
	fmt.Println("Timer expired!")
	timer.Reset(5 * time.Second) // 重新设置定时器
	// default:
	// 	fmt.Println("Doing some other work.")
	// 	time.Sleep(1 * time.Second)
}
