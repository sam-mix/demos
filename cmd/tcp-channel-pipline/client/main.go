// 对于管道模型的实现，客户端需要从TCP连接中读取数据，并将数据发送到输入通道。然后客户端从输出通道中读取处理结果，并将结果写入TCP连接。

// 以下是一个示例客户端代码，用于与您提供的服务器代码进行通信：

// ```go
// package main

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"log"
// 	"net"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8899")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	for i := 0; i < 10; i++ {
// 		// 将数据发送到TCP连接
// 		buf := make([]byte, 4)
// 		binary.BigEndian.PutUint32(buf, uint32(i))
// 		conn.Write(buf)

// 		// 从TCP连接读取处理结果
// 		buf = make([]byte, 4)
// 		_, err := conn.Read(buf)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		result := int(binary.BigEndian.Uint32(buf))

// 		fmt.Println(result)
// 	}
// }
// ```

// 此客户端将数字0到9发送到服务器，并打印加1后的结果。您可以根据您的需求修改该代码。

package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		// 将数据发送到TCP连接
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(i))
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

		fmt.Println(result)
	}
}
