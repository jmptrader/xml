package xml

import (
	"fmt"
	"reflect"
)

//把dom文件映射到对象中
func (this *Node) Mapping(obj interface{}) {
	buildNode(obj, this)
}

func Mapping(obj interface{}) *Node {
	return nil
}

func buildNode(obj interface{}, node *Node) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		fmt.Println("传入的值必须是指针")
		return
	}
	objType = objType.Elem()
	if objType.Kind() != reflect.Struct {
		fmt.Println("传入的值必须是结构体")
		return
	}

	objValue := reflect.ValueOf(obj).Elem()
	buildNodeLoop(&objType, &objValue, node)

}

func buildNodeLoop(objType *reflect.Type, objValue *reflect.Value, node *Node) {

	for i := 0; i < (*objType).NumField(); i++ {
		if !objValue.Field(i).CanSet() {
			continue
		}
		tags := (*objType).Field(i).Tag
		if name := tags.Get("name"); name != "" {
			objValue.Field(i).SetString(node.GetName())
		}
		if attr := tags.Get("attr"); attr != "" {
			objValue.Field(i).SetString(node.GetAttr(attr))
		}
		if text := tags.Get("text"); text != "" {
			objValue.Field(i).SetString(node.GetText())
		}
		if child := tags.Get("child"); child != "" {
			// childs := node.GetChild()
			// for _, childOne := range childs {
			// 	if childOne.GetName() == child {
			// 		objTypeOne := reflect.TypeOf((*objType).Field(i))
			// 		fmt.Println(objTypeOne)
			// 		objValueOne := reflect.ValueOf((*objValue).Field(i))
			// 		buildNodeLoop(&objTypeOne, &objValueOne, childOne)
			// 	}
			// }
		}
	}
}
