
package main

import (
	"fmt"
	"net"
)




type HttpRequest struct {
	StartLine
	Header
	Body
}

type Body struct {
	bytes []byte
}

type Header struct {
	count int
	items []HeaderItem
}

func NewEmptyHeader() Header {
	return Header{
		count:0,
		items:make([]HeaderItem, 10),
	}
}

type HeaderItem struct {
	key string
	value string
}

func (p *HeaderItem) printStr() {
	fmt.Println(p.key+"="+p.value)
}

func (p *HeaderItem) parseByte(bytes []byte) error {

	return nil
}

type StartLine struct {
	method string
	requestTarget string
	httpVersion string
}

/**
解析起始行
 */
func (p *StartLine) parseByte(bytes []byte) error {
	cur := 0
	for !(bytes[cur] == 32) {
		cur++
	}
	p.method = string(bytes[0:cur])
	cur++//此时cur越过空格

	prev := cur
	for !(bytes[cur] == 32) {
		cur++
	}
	p.requestTarget = string(bytes[prev:cur])
	cur++

	p.httpVersion = string(bytes[cur:len(bytes)-2])

	return nil
}

/**
起始行打印控制台
 */
func (p *StartLine) printStr() {
	fmt.Println(p.method + " " + p.requestTarget + " " + p.httpVersion)
}

func handleConn(conn net.Conn) {

	if conn == nil {
		return
	}

	buf := make([]byte, 4096)

	//读取所有数据
	bytes := make([]byte, 3000)
	for {
		cnt, err := conn.Read(buf)
		if err != nil || cnt == 0 {
			err := conn.Close()
			if err != nil {
				fmt.Println("Connection close error!")
			}
			fmt.Println("Connection has closed!")
			break
		}
		bytes = append(bytes, buf...)
	}

	cur := 0

	startLineBytes := make([]byte, 30)
	//读取起始行
	for !(bytes[cur] == 13 && bytes[cur+1] == 10){
		cur++
	}
	cur += 2


	var startLine StartLine
	err := startLine.parseByte(append(startLineBytes, bytes[0:cur]...))
	if err != nil {
		fmt.Println("start line parse exception!")
		return
	}
	startLine.printStr()

	//读取首部
	headerFlag := 0
	header := NewEmptyHeader()
	for headerFlag != 1 {
		prev := cur
		for !(bytes[cur] == 13 && bytes[cur+1] == 10) {
			cur++
		}

		//如果为空行
		if prev == cur {
			headerFlag = 1
			continue
		}

		//读取header
		var item HeaderItem
		err := item.parseByte(bytes[prev:cur+2])
		if err != nil {
			fmt.Println("Header parse error!")
		}
		header.items[header.count] = item
		header.count++

		cur += 2
		prev = cur
	}


	//读取正文
	//body := make([]byte, 20)


	//产生响应

}


func main()  {

	fmt.Println("Welcome HK Http Server!")

	server, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}


}