package backend

import (
	"fmt"
	"log"
	"net"
	// "os"
)

type server struct {

}


func NewServer() *server {
	return &server{
	}
}

func (s *server) Startup(port string) {

  	fmt.Println("I'm Server")

		listener, err := net.Listen("tcp", "0.0.0.0:"+ port)

		if err != nil {
				log.Fatalf("unable to listen :%s", err.Error())
				fmt.Println(err)
				return
		}
		defer listener.Close()

		fmt.Println("start listening")
		//每接收到一个新的TCP连接，就用new一个Handler来进行处理
		for {
				conn, err := listener.Accept()
				if err != nil {
						fmt.Println(err)
						return
				}
				// s.conn = conn 
				go s.newHandler(conn)
		}

}

func (s *server) newHandler(conn net.Conn) {

		h := &handler{
			conn: conn,
			repoPath: "D://repo",
		}

		h.run()
}


