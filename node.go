package xml

import (
	"bytes"
	// "fmt"
	"sync"
)

type Node struct {
	lock   *sync.RWMutex
	Name   string            //节点名称
	Attr   map[string]string //节点属性
	Text   string            //节点值
	Parent *Node             //父节点
	Child  []*Node           //节点的子节点
	Type   int               //标签类型  1.开始标签，2.结束标签，3.自关闭标签，4.xml头部声明
}

//得到节点的名称
func (this *Node) GetName() string {
	return this.Name
}

//设置节点名称
func (this *Node) SetName(name string) {
	this.Name = name
}

//得到标签中的文本
func (this *Node) GetText() string {
	return this.Text
}

//设置标签中的文本
func (this *Node) SetText(text string) {
	this.Text = text
}

//得到一个属性
func (this *Node) GetAttr(name string) string {
	return this.Attr[name]
}

//设置一个属性
func (this *Node) SetAttr(name, value string) {
	this.Attr[name] = value
}

//删除一个属性
func (this *Node) DelAttr(name string) {
	delete(this.Attr, name)
}

//得到父节点
func (this *Node) GetParent() *Node {
	return this.Parent
}

//设置父节点
func (this *Node) SetParent(node *Node) {
	this.Parent = node
}

//得到根节点
func (this *Node) GetRoot() *Node {
	if this.Parent == nil {
		return this
	} else {
		root := this.Parent
		for {
			if root.Parent == nil {
				return root
			} else {
				root = root.Parent
			}
		}
	}
}

//添加一个子节点
func (this *Node) AddChild(node *Node) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.Child == nil {
		this.Child = make([]*Node, 0)
		this.Child = append(this.Child, node)
		node.SetParent(this)
		return
	}
	node.SetParent(this)
	this.Child = append(this.Child, node)
	// fmt.Println("node +++++  ", node)
}

//删除一个子节点
func (this *Node) DelChild(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	for i, nodeOne := range this.Child {
		if nodeOne.GetName() == name {
			tempChild := make([]*Node, 0)
			tempChild = append(tempChild, this.Child[:i]...)
			tempChild = append(tempChild, this.Child[i+1:]...)
			this.Child = tempChild
			return true
		}
	}
	return false
}

//得到所有子节点
func (this *Node) GetChild() []*Node {
	return this.Child
}

//得到一个子节点
func (this *Node) GetChildOne(name string) (*Node, bool) {
	for _, nodeOne := range this.Child {
		if nodeOne.GetName() == name {
			return nodeOne, true
		}
	}
	return &Node{}, false
}

//生成xml字符串
func (this *Node) BuildXML() string {
	return "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n" + this.buildXml(0)
}
func (n *Node) buildXml(level int) string {
	tabStr := ""
	for i := 0; i < level*4; i++ {
		tabStr = tabStr + " "
	}
	buf := bytes.NewBufferString(tabStr + "<")
	buf.WriteString(n.Name)
	for key, value := range n.Attr {
		buf.WriteString(" ")
		buf.WriteString(key + "=" + value)
	}
	if n.Type == 3 && len(n.Child) == 0 {
		buf.WriteString(" />" + n.Text + "\r\n")
		return buf.String()
	} else {
		buf.WriteString(">" + n.Text)
	}

	for _, node := range n.Child {
		buf.WriteString(node.buildXml(level + 1))
	}
	buf.WriteString(tabStr + "</" + n.Name + ">\r\n")
	return buf.String()
}

//对比两个node内容是否一样
func (this *Node) Equals(node *Node) bool {
	if this.Name != node.Name {
		return false
	}
	if this.Text != node.Text {
		return false
	}
	if len(this.Attr) != len(node.Attr) {
		return false
	}
	for keyOne, valueOne := range this.Attr {
		if value, ok := node.Attr[keyOne]; ok {
			if value != valueOne {
				return false
			}
		} else {
			return false
		}
	}
	if len(this.Child) != len(node.Child) {
		return false
	}
	for _, childOne := range this.Child {
		find := false
		for _, nodeChildOne := range node.Child {
			if childOne.Equals(nodeChildOne) {
				find = true
				break
			}
		}
		if !find {
			return false
		}
	}
	return true
}

//创建一个新的Node
func NewNode() *Node {
	return &Node{
		lock:  new(sync.RWMutex),
		Attr:  make(map[string]string),
		Child: make([]*Node, 0),
	}
}
