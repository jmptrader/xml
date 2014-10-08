package xml

import (
	// "encoding/xml"
	// "fmt"
	"strings"
	// "io/ioutil"
	// "bytes"
)

var (
	// endComment     = []byte("-->")
	// endProcInst    = []byte("?>")
	startDirective = byte('<')
	endDirective   = byte('>')
)

func Unmarshal(data []byte) *Node {
	p := Parser{data: data}
	return p.Parse()

	// fmt.Println(p.getMark())
}

type Parser struct {
	data  []byte
	index int
}

//开始分析
func (this *Parser) Parse() *Node {
	var root *Node
	startIndex := 0
	// endIndex := 0
	marker := false
	for i := 0; i < len(this.data); i++ {
		if this.data[i] == startDirective {
			if marker {
				panic("")
			} else {
				marker = true
				startIndex = i
			}
		} else if this.data[i] == endDirective {

			if marker {
				marker = false
				// startIndex = i
				node := this.ParseLine(string(this.data[startIndex+1 : i]))

				// fmt.Println(root, "\n")
				// if root != nil {
				// 	fmt.Println(root.GetParent())
				// }
				// fmt.Println("-----------------")
				// fmt.Println(string(this.data[startIndex+1:i]), "\n")

				switch node.Type {
				case 1:
					if root == nil {
						root = node
						continue
					}
					root.AddChild(node)
					// node.SetParent(root)
					root = node
				case 2:
					if node.GetName() == "" {
						if root.GetParent() != nil {
							root = root.GetParent()
						}
						continue
					}
					panic(node.GetName() + " 缺少结尾标签")
				case 3:
					if root == nil {
						root = node
						continue
					}
					root.AddChild(node)
					// root = node.GetParent()
				case 4:
				}
				// if node.GetName() == "" && node.isStart == false {
				// 	root = node
				// 	continue
				// }
				// if node.isStart {
				// 	root.AddChild(*node)
				// 	continue
				// }
				// if node.Name == root.GetName() {
				// 	root = root.GetParent()
				// }
			} else {
				panic("")
			}

		}
	}
	// fmt.Println(root)
	return root.GetRoot()
}

//分析一个标记
func (this *Parser) ParseLine(str string) *Node {
	node := NewNode()
	if strings.HasPrefix(str, "?") {
		node.Type = 4
		return node
	}
	if strings.HasPrefix(str, "/") {
		node.Type = 2
		return node
	}
	strs := strings.Split(str, " ")
	node.SetName(strs[0])
	for i := 1; i < len(strs)-2; i++ {
		attrOne := strings.Split(strs[i], "=")
		node.SetAttr(attrOne[0], attrOne[1])
	}
	if strings.HasSuffix(str, "/") {
		node.Type = 3
		return node
	}
	node.Type = 1
	return node
}

//得到一个标记
func (this *Parser) getMark() string {
	for i := this.index; i < len(this.data); i++ {
		if this.data[i] == endDirective {
			return string(this.data[:i+1])
		}
	}
	return ""
}
