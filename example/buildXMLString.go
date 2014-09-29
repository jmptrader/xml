package main

import (
	"fmt"
	"github.com/prestonTao/xml"
)

func main() {
	user := xml.Node{Name: "user", Attr: map[string]string{"name": "tao"}, Value: "你好"}
	person := xml.Node{Name: "person", Child: []xml.Node{user}}
	fmt.Println(person.BuildXML())
}
