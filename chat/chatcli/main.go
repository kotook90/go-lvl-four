package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Соединение с сервером установлено\n")
	}
	defer conn.Close()

	/*
		who := conn.LocalAddr().String()
			fmt.Println("Введите свой никнейм: ")
			var name string
			fmt.Scan(&name)
	*/

	go func() {
		io.Copy(os.Stdout, conn)
	}()

	io.Copy(conn, os.Stdin) // until you send ^Z

	fmt.Printf("%s: exit", conn.LocalAddr())
}
