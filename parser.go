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
	//记录左尖括号下标
	startIndex := 0
	//记录右尖括号下标
	endIndex := 0
	marker := false
	for i := 0; i < len(this.data); i++ {
		if this.data[i] == startDirective {
			if !marker {
				marker = true
				startIndex = i
				//拿到text的内容
				if endIndex != 0 && endIndex < i && endIndex+1 != i {
					text := string(this.data[endIndex+1 : i])
					// if i < 200 {
					// 	fmt.Println("+++++++++++", text)
					// }

					//去除空格
					// text = strings.Replace(text, "\\s", "", -1)
					text = strings.Replace(text, "\n", "", -1)
					text = strings.Replace(text, "\t", "", -1)
					text = strings.Replace(text, "\r", "", -1)
					// if i < 200 {
					// 	fmt.Println("===========", text)
					// }
					if text != "" && len(root.GetChild()) == 0 {
						// fmt.Println("++", string(this.data[endIndex+1:i]), startIndex, endIndex)
						root.SetText(string(this.data[endIndex+1 : i]))
					}
				}
				// if startIndex != 0 && startIndex+1 != i {
				// 	// fmt.Println(len(this.data), startIndex, i)
				// 	text := string(this.data[startIndex+1 : i])
				// 	//去除空格
				// 	// text = strings.Replace(text, "\\s", "", -1)
				// 	text = strings.Replace(text, "\n", "", -1)
				// 	text = strings.Replace(text, "\t", "", -1)
				// 	text = strings.Replace(text, "\r", "", -1)
				// 	if text != "" && len(root.GetChild()) == 0 {
				// 		fmt.Println(string(this.data[startIndex+1 : i]))
				// 		root.SetText(string(this.data[startIndex+1 : i]))
				// 	}
				// }

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
				endIndex = i
			}

		}
	}
	// fmt.Println(root)
	if len(root.GetRoot().GetChild()) != 0 {
		root.GetRoot().SetText("")
	}
	return root.GetRoot()
}

//分析一个标记
func (this *Parser) ParseLine(str string) *Node {
	// fmt.Println(str, "\n++++++++++++")
	//去除空格
	// str = strings.Replace(str, "\\s", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	// fmt.Println(str, "\n===============")
	node := NewNode()
	if strings.HasPrefix(str, "?") {
		node.Type = 4
		return node
	}
	if strings.HasPrefix(str, "/") {
		node.Type = 2
		return node
	}
	if strings.HasSuffix(str, "/") {
		node.Type = 3
		str = str[:len(str)-1]
	}

	strs := strings.Split(str, " ")
	// fmt.Println(strs, "\n-------------------------------------", len(strs))
	node.SetName(strs[0])
	for i := 1; i < len(strs); i++ {
		// fmt.Println(strs[i])
		if strs[i] == "" {
			continue
		}
		attrOne := strings.Split(strs[i], "=")
		if len(attrOne) == 0 {
			continue
		}
		attrValue := attrOne[1]
		node.SetAttr(attrOne[0], attrValue[1:len(attrValue)-1])
	}

	if node.Type == 3 {
		return node
	}
	node.Type = 1

	// fmt.Println(node)
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
