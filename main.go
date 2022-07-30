package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	Listener, Err := net.Listen("tcp", "0.0.0.0:8888")

	if Err != nil {

	}

	fmt.Printf("Echo Server Started : %v\n", Listener.Addr())

	Chan := AcceptClient(Listener)
	for {
		go HandleClient(<-Chan)
	}
}

func AcceptClient(Listener net.Listener) chan net.Conn {
	Chan := make(chan net.Conn)

	go func() {
		for {
			Client, Err := Listener.Accept()
			if Err != nil {

			}

			fmt.Printf("New Connection : %v \n", Client.LocalAddr())
			Chan <- Client
		}
	}()

	return Chan
}

func HandleClient(Client net.Conn) {
	defer Client.Close()

	Reader := bufio.NewReader(Client)
	for {
		Msg, Err := Reader.ReadBytes('\n')
		if Err != nil {
			break
		}

		fmt.Println(string(Msg))

		_, Err = Client.Write(Msg)
		if Err != nil {
			break
		}
	}
}
