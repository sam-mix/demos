/*
 * SERVIDOR TCP CON COMUNICACION SEGURA A TRAVES DEL PROTOCOLO TLS
 * Envio de informacion en texto plano UTF8 de forma segura (uso de tls con certificado autofirmado)
 */
package main

import (
	"crypto/rand"
	"crypto/tls"
	"demos/msg"
	"demos/pb"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	StartServer()
}

// StartServer Inicia el servidor TCP - TLS
func StartServer() {
	//Cargar el certificado y la clave privada de la comunicacion
	cert, err := tls.LoadX509KeyPair("resources/certs/server.pem", "resources/certs/server.key")
	checkError(err) //Comprobar errores

	//Crear configuracion del socket seguro
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	now := time.Now()
	config.Time = func() time.Time { return now } //Establecer hora actual
	config.Rand = rand.Reader                     //Numero aleatorio utilizado por TLS para el Blinding de RSA

	//Iniciar el socket de servidor
	service := "0.0.0.0:2005"
	socket, err := tls.Listen("tcp", service, &config)
	checkError(err) // comprobar errores

	//Defer hace que el codigo se ejecute al final de la ejecucion del bloque de codigo
	defer socket.Close()

	//Escuchar nuevos clientes
	fmt.Println("Servidor TCP-TLS esperando clientes...")
	for {
		conn, err := socket.Accept()
		//Si ocurre un error continuar con el bucle de escucha
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("✅ Nuevo cliente conectado (" + conn.RemoteAddr().String() + ")")
		//Iniciar gorutina de comunicacion
		go handleClientCommunication(conn)
	}
}

//------------------------------------------------------------>

// Manejar la comunicacion con un cliente
func handleClientCommunication(conn net.Conn) {
	//Defer hace que el codigo se ejecute al final de la ejecucion del bloque de codigo
	defer conn.Close()

	inputBuffer := make([]byte, 512)
	for {
		// Leer mensajes
		n, err := conn.Read(inputBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		txtInputBuffer := string(inputBuffer[:n])

		time.Sleep(3 * time.Second)
		if n < msg.CLIMSGHEADLEN {
			txtInputBuffer = "消息长度比消息头的长度还小"
			fmt.Println("消息总长度: ", n)
		} else {
			headBuf := make([]byte, msg.CLIMSGHEADLEN)
			head := new(msg.ClientMsgHead)
			head.UnMsgPackBE(headBuf)
			fmt.Printf("传过来的消息头:%v\n", head)
			if uint32(n) != head.MsgLen+msg.CLIMSGHEADLEN {
				txtInputBuffer = fmt.Sprintf("消息体长度不对: %d", n-msg.CLIMSGHEADLEN)

			} else {

				if head.MsgID == uint16(pb.MSG_ID_HEART) {
					out, err := proto.Marshal(&pb.S2CHeart{
						Timestamp: time.Now().UnixMilli(),
					})
					if err != nil {
						log.Fatalln("Failed to encode address book:", err)
					}
					txtInputBuffer = string(out)
				} else {
					txtInputBuffer = string(inputBuffer[msg.CLIMSGHEADLEN:n])
				}
			}

		}
		// tcp server 服务端代码
		var cmh = &msg.ClientMsgHead{
			MsgLen: uint32(len([]byte(txtInputBuffer))),
			MsgID:  1,
			SvrID:  0,
			CCode:  0,
		}
		bodyBytes := []byte(txtInputBuffer)
		headBytes, _ := cmh.MsgPackBE()
		msgBin := append(headBytes, bodyBytes...)

		// Enviar mensaje
		_, err2 := conn.Write(msgBin)

		fmt.Print("\n---------------------------------------\n")
		fmt.Println("🟦 Recibido paquete del cliente " + conn.RemoteAddr().String() + ":")
		fmt.Println(txtInputBuffer)
		fmt.Println(inputBuffer)
		fmt.Print("---------------------------------------\n")

		if err2 != nil {
			fmt.Println(err2)
		}
	}

	println("🔌 Fin de la conexion del cliente " + conn.RemoteAddr().String())
}

//------------------------------------------------------------>

// Imprimir un error
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
