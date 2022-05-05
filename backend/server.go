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

		// 输入端口号
		// fmt.Println("Enter the port to Listen: ")
    // var port string
		// fmt.Scanln(&port)

		listener, err := net.Listen("tcp", "0.0.0.0:"+ port)

		if err != nil {
				log.Fatalf("unable to Listen :%s", err.Error())
				fmt.Println(err)
				return
		}
		defer listener.Close()

		fmt.Println("开始监听")

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


