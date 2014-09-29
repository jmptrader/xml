package xml

import (
	"bytes"
	"sync"
)

type Node struct {
	lock  *sync.RWMutex
	Name  string            //节点名称
	Attr  map[string]string //节点属性
	Value string            //节点值
	Child []Node            //节点的子节点
}

//得到节点的名称
func (this *Node) GetName() string {
	return this.Name
}

//设置节点名称
func (this *Node) SetName(name string) {
	this.Name = name
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

//添加一个子节点
func (this *Node) AddChild(node Node) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.Child == nil {
		this.Child = make([]Node, 0)
		this.Child = append(this.Child, node)
		return
	}
	this.Child = append(this.Child, node)
}

//删除一个子节点
func (this *Node) DelChild(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	for i, nodeOne := range this.Child {
		if nodeOne.GetName() == name {
			tempChild := make([]Node, 0)
			tempChild = append(tempChild, this.Child[:i]...)
			tempChild = append(tempChild, this.Child[i+1:]...)
			this.Child = tempChild
			return true
		}
	}
	return false
}

//得到所有子节点
func (this *Node) GetChild() []Node {
	return this.Child
}

//得到一个子节点
func (this *Node) GetChildOne(name string) (Node, bool) {
	for _, nodeOne := range this.Child {
		if nodeOne.GetName() == name {
			return nodeOne, true
		}
	}
	return Node{}, false
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
	buf.WriteString(">" + n.Value + "\r\n")
	for _, node := range n.Child {
		buf.WriteString(node.buildXml(level + 1))
	}
	buf.WriteString(tabStr + "</" + n.Name + ">\r\n")
	return buf.String()
}

//创建一个新的Node
func NewNode() *Node {
	return &Node{
		lock:  new(sync.RWMutex),
		Attr:  make(map[string]string),
		Child: make([]Node, 0),
	}
}
