package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("ex) hexstrsender.exe 127.0.0.1:3002")
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("connected input hexastring to send")

	go func() {
		for {
			recv := make([]byte, 1024)
			n, err := conn.Read(recv)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(hex.EncodeToString(recv[:n]))
		}
	}()

	var packet string
	if len(os.Args) > 2 {
		packet = os.Args[2]
	}

	for {

		if len(packet) == 0 {
			fmt.Scanln(&packet)
		}

		b, err := hex.DecodeString(packet)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if _, err := conn.Write(b); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		packet = ""
	}
}
