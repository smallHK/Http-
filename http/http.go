package http

import "fmt"

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


