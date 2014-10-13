package main

import (
	"fmt"
	"github.com/prestonTao/xml"
	"io/ioutil"
)

func main() {
	example1()
}

func example1() {
	data, err := ioutil.ReadFile("styles.xml")
	if err != nil {
		fmt.Println(err.Error())
	}
	node := xml.Unmarshal(data)
	// for _, nodeOne := range node.GetChild() {
	// 	fmt.Println(nodeOne)
	// }
	fmt.Println(node.BuildXML())

}
