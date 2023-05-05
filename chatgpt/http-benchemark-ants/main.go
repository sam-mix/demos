package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	ants "github.com/panjf2000/ants/v2"
)

func main() {
	var wg sync.WaitGroup

	// 创建一个 Ants 并发池，控制最大 Goroutine 数量为 1000
	p, err := ants.NewPool(1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer p.Release()

	// 启动 1000 个并发任务
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		p.Submit(func() {
			defer wg.Done()

			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			request := "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"
			_, err = conn.Write([]byte(request))
			if err != nil {
				fmt.Println(err)
				return
			}

			buf := make([]byte, 1024)
			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
		})
	}

	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed.Seconds(), "seconds")
	fmt.Println("Requests per second:", float64(1000)/elapsed.Seconds())
}
