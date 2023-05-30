// package main

// 在Go语言中，可以通过TCP连接和channel实现pipeline模式。TCP连接用于实现不同计算机之间的数据传输，而pipeline模式用于将一个大任务拆分成多个小任务交由不同协程处理。

// 以下是一个简单的TCP channel pipeline示例，其中包含两个阶段：

// ```go
// func pipeline(conn net.Conn) {
//     defer conn.Close()

//     in := make(chan int)
//     out := make(chan int)

//     go func() {
//         defer close(in)
//         for i := 0; i < 10; i++ {
//             in <- i
//         }
//     }()

//     go func() {
//         defer close(out)
//         for num := range out {
//             // 第一步：将输入加上1
//             plusOne := num + 1

//             // 第二步：将结果发送到输出通道
//             out <- plusOne
//         }
//     }()

//     for num := range in {
//         // 将输入发送到TCP连接
//         buf := make([]byte, 4)
//         binary.BigEndian.PutUint32(buf, uint32(num))
//         conn.Write(buf)

//         // 从TCP连接读取处理结果
//         buf = make([]byte, 4)
//         _, err := conn.Read(buf)
//         if err == io.EOF {
//             break
//         }
//         if err != nil {
//             log.Fatal(err)
//         }
//         result := int(binary.BigEndian.Uint32(buf))

//         // 将结果发送到输出通道
//         out <- result
//     }
// }

// func main() {
//     listener, err := net.Listen("tcp", ":8888")
//     if err != nil {
//         log.Fatal(err)
//     }
//     for {
//         conn, err := listener.Accept()
//         if err != nil {
//             log.Fatal(err)
//         }
//         go pipeline(conn)
//     }
// }
// ```

// 上述示例中，pipeline函数接受一个TCP连接作为参数。函数内部创建输入通道in和输出通道out，通过for循环将0到9的整数发送到输入通道in中。在第二个go协程中，通过for...range遍历输出通道out，对每个元素进行处理，然后发送到TCP连接中。在主循环中，通过listener.Accept()函数接受新的连接，并将TCP连接传入pipeline函数中。

// 值得注意的是，本例中使用了`binary.BigEndian.PutUint32`函数将整数转换成字节序列，然后通过conn.Write函数将字节序列发送到TCP连接中。在读取结果时，需要先从TCP连接中读取字节序列，然后使用`binary.BigEndian.Uint32`函数将字节序列转换成整数。

// 总的来说，TCP channel pipeline是一种高效的并发编程方式，能够提高程序运行效率，降低资源消耗。在Go语言中，通过TCP连接和channel实现pipeline非常容易，只需借助TCP连接实现计算机间的数据传输，借助channel实现协程间的数据传输即可。
package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

func pipeline(conn net.Conn) {
	defer conn.Close()

	in := make(chan int)
	out := make(chan int)

	go func() {
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	go func() {
		defer close(out)
		for num := range out {
			// 第一步：将输入加上1
			plusOne := num + 1

			// 第二步：将结果发送到输出通道
			out <- plusOne
		}
	}()

	for num := range in {
		// 将输入发送到TCP连接
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(num))
		conn.Write(buf)

		// 从TCP连接读取处理结果
		buf = make([]byte, 4)
		_, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result := int(binary.BigEndian.Uint32(buf))

		// 将结果发送到输出通道
		out <- result
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go pipeline(conn)
	}
}
