package main

import (
	"fmt"
	"github.com/prestonTao/xml"
	"io/ioutil"
	"os"
)

func main() {
	example2()
}

func example2() {
	data, err := ioutil.ReadFile("strings_91.xml")
	if err != nil {
		fmt.Println(err.Error())
	}
	node := xml.Unmarshal(data)
	// for _, nodeOne := range node.GetChild() {
	// 	fmt.Println(nodeOne)
	// }
	fmt.Println(node.BuildXML())

	newfile, _ := os.Create("temp.xml")
	newfile.WriteString(node.BuildXML())
	newfile.Close()
}

func example1() {
	data, err := ioutil.ReadFile("AndroidManifest.xml")
	if err != nil {
		fmt.Println(err.Error())
	}
	node := xml.Unmarshal(data)
	fmt.Println(node, "\n")

	manifest := new(Manifest)
	node.Mapping(manifest)
	fmt.Println(manifest)

}

type Manifest struct {
	Name        string    `name:"manifest"`
	Theme       string    `attr:"android:theme"`
	versionCode string    `attr:"android:versionCode"`
	Supports    *Supports `child:"supports-screens"`
}

type Supports struct {
	Name string `name:"supports-screens"`
}
