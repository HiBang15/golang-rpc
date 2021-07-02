package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}

func (service *HelloService)Hello(request string, reply *string) error {
	*reply = "Hello " + request

	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("ListenTCP error: ", err.Error())
	}

	// vòng lặp để phục vụ nhiều client
	for {
		// accept một connection đến
		conn, err := listener.Accept()
		// in ra lỗi nếu có
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// phục vụ client trên một goroutine khác
		// để giải phóng main thread tiếp tục vòng lặp
		go rpc.ServeConn(conn)
	}
}