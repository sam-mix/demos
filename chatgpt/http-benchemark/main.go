package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
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
		}()
	}

	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed.Seconds(), "seconds")
	fmt.Println("Requests per second:", float64(1000)/elapsed.Seconds())
}
