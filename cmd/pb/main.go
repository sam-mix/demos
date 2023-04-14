package main

import (
	"demos/pb"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	protoApi()
	fileName := "writeMsg.out"
	writeMsg(fileName)
	readMsg(fileName)
}

func readMsg(fname string) {
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("%#v\n", book)
	fmt.Println("===========================================================================")
}

func writeMsg(fname string) {
	book := &pb.AddressBook{
		People: []*pb.Person{
			{
				Id:    1234,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.Person_HOME},
				},
			},
			{
				Id:    1234,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.Person_HOME},
				},
			},
			{
				Id:    1234,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.Person_HOME},
				},
			},
			{
				Id:    1234,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.Person_HOME},
				},
			},
		},
	}
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func protoApi() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	fmt.Printf("%#v\n", &p)
	fmt.Println("===========================================================================")
}
