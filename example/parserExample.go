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
