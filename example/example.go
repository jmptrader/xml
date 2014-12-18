package main

import (
	"fmt"
	"github.com/prestonTao/xml"
	"io/ioutil"
)

func main() {
	TestEqualsExample()
	DelChildExample()
}

//删除节点测试
func DelChildExample() {
	root := xml.NewNode()
	root.SetName("root")

	ch1 := xml.NewNode()
	ch1.SetName("ch1")
	root.AddChild(ch1)

	ch2 := xml.NewNode()
	ch2.SetName("ch1")
	root.AddChild(ch2)

	fmt.Println(root.BuildXML())

	temRoot := *root
	temRoot.DelChild(*ch1)

	fmt.Println(temRoot.BuildXML())

	root.DelChildForName("ch1")

	fmt.Println(root.BuildXML())

}

//对比node测试
func TestEqualsExample() {
	root := xml.NewNode()
	root.SetName("root")

	ch1 := xml.NewNode()
	ch1.SetName("ch1")
	ch1.SetAttr("nimei", "haha")
	root.AddChild(ch1)

	ch2 := xml.NewNode()
	ch2.SetName("ch2")
	root.AddChild(ch2)

	root2 := xml.NewNode()
	root2.SetName("root")
	root2.AddChild(ch1)
	root2.AddChild(ch2)

	fmt.Println(root.Equals(root2))

}

//构建一个xml文件
func BuildXMLExample() {
	user := xml.Node{Name: "user", Attr: map[string]string{"name": "tao"}, Text: "你好"}
	person := xml.Node{Name: "person", Child: []*xml.Node{&user}}
	fmt.Println(person.BuildXML())
}

//测试SetText方法
func TestSetTextExample() {
	root := xml.NewNode()
	root.SetName("root")

	ch1 := xml.NewNode()
	ch1.SetName("ch1")
	root.AddChild(ch1)

	ch2 := xml.NewNode()
	ch2.SetName("ch2")
	root.AddChild(ch2)

	for _, one := range root.GetChild() {
		one.SetText("123")
	}

	fmt.Println(root.BuildXML())
}

//解析styles.xml文件
func ParseExample1() {
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
