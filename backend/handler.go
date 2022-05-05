package backend

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
	// "os"
)

type handler struct {
	conn net.Conn
	// reply string 
		//指定仓库目录
	repoPath string
}

func (h *handler) run() {

	for {
		msg, err := bufio.NewReader(h.conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("> " + msg)
		h.parseCmd(msg)
	}
}


func (h *handler)parseCmd (msg string) {

	msg = strings.Trim(msg, "\r\n")
	args := strings.Split(msg, " ")
	cmd := strings.TrimSpace(args[0])

	if len(args) == 1  {

			switch cmd {
			case "time":
				h.timeCmd()
			case "exit":
				h.exitCmd()
			case "ls":
				h.lsCmd()
			default:
				h.noCmd(msg)
			}
	} else if len(args) == 2 {

			switch cmd {
			case "req":
				h.reqCmd(args[1])
			default:
				h.noCmd(msg)
			}

	} else{
			h.noCmd(msg)
	}
}


func (h *handler)noCmd(msg string) {
		h.replyClient(msg)
		fmt.Println("无效命令...")
}


func  (h *handler)timeCmd() {
		// h.conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + "\n" ));
		h.replyClient(time.Now().Format("2006-01-02 15:04:05"))
		// h.reply = time.Now().Format("2006-01-02 15:04:05")
}


func  (h *handler)exitCmd() {
		// h.conn.Write([]byte("Bye" +"\r\n"));
		h.replyClient("Bye")
		h.conn.Close()
}

/*
处理ls命令
*/
func  (h *handler)lsCmd() {

	files, err := h.listFiles()
	if err != nil {
		// h.conn.Write([]byte("Ls Error:" + err.Error() + "\n"));
		fmt.Println("Error:", err)
		h.replyClient("服务器错误!")
		return
	}	
	s := strings.Join(files, ",")
	// fmt.Println(s)
	h.replyClient(s)
	// h.reply = s
}

func  (h *handler)reqCmd(filename string) {

	files, _ := h.listFiles()
	var prompt string
	for _, f := range files {
			if filename == f {
					h.replyClient( f + "?")
					buf := make([]byte, 1024)
					n, _ := h.conn.Read(buf)
					fmt.Println("等待ACK")
					//如果客户端回复为ACK则发送文件
					if string(buf[:n]) == "ACK" {
						 h.sendFile(filename)
						 fmt.Println("传输成功")
						 return
					}
			}
			if strings.HasPrefix(f, filename) {
					prompt = f
			}
	}
	h.replyClient( filename + "不存在，你或许是想输入——" + prompt)
	
}


func (h *handler)replyClient(msg string) {
	// 同时打印回复给客户端的内容
	// fmt.Println("< " + h.reply)
	// h.conn.Write([]byte(h.reply + "\r\n" ))
		 fmt.Println("< " + msg)
	   h.conn.Write([]byte(msg + "\r\n" ))
}


func (h *handler)sendFile(filename string) {

	filePath := h.repoPath + string(os.PathSeparator) + filename
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.open error" , err)
	}
	buf := make([]byte, 4096)

	for {
		n, err := file.Read(buf)

		if err == io.EOF {
				fmt.Println("读取完毕")
				// h.conn.Write([]byte("EOF"))
				return
		}
		if err != nil {
				fmt.Println("file read error" , err)
				return
		}
		fmt.Println("n =" , n)
		h.conn.Write(buf[:n])
	}
} 


func  (h *handler)listFiles() (files []string, err error) {

	dirPath := h.repoPath
	//为slice分配空间
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
			 return nil, err
	}
	// pathSep := string(os.PathSeparator)
	// suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, f := range dir {
	 if f.IsDir() { // 忽略目录
				continue
	 }
	//  if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
		// files = append(files, dirPath + pathSep + fi.Name())
	//  }
		files = append(files, f.Name())
	}
	return files, nil
}