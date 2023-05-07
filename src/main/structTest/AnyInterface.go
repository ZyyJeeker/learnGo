package structTest

import (
	"fmt"
	"sort"
	"unsafe"
)

type Element interface {
}

type ArrayList struct {
	elements []Element
}

func (s *ArrayList) Len() int {
	return len(s.elements)
}

func (s *ArrayList) Less(i, j int) bool {
	eli := s.elements[i]
	elj := s.elements[j]

	return unsafe.Sizeof(eli)-unsafe.Sizeof(elj) > 0
}

func (s *ArrayList) Swap(i, j int) {
	s.elements[i], s.elements[j] = s.elements[j], s.elements[i]
}

func (s *ArrayList) SetElements(elements []Element) {
	s.elements = elements
}

func (s *ArrayList) At(idx int) Element {
	return s.elements[idx]
}

func (s *ArrayList) Set(index int, el Element) {
	s.elements[index] = el
}

func AnyInterfaceTest() {
	elements := make([]Element, 5)
	list := new(ArrayList)
	list.SetElements(elements)

	list.Set(0, Data{[]Node{{1, 2}}})

	list.Set(1, Person{
		FirstName: "张",
		LastName:  "宇",
	})

	list.Set(2, Phone{
		cpu:     "骁龙",
		ram:     "12g",
		price:   3999,
		version: 2,
	})

	list.Set(3, Camera{"尼康k22"})
	list.Set(4, Car{
		Gopher: Gopher{2, "golang"},
	})

	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)

	for _, el := range list.elements {
		switch el.(type) {
		case Car:
			fmt.Printf("这是一个car: %v\n", el)
		case Camera:
			fmt.Printf("camera: %v\n", el)
		case Phone:
			fmt.Printf("这是一个phone: %v\n", el)
		case Person:
			fmt.Printf("这是一个person: %v\n", el)
		case Data:
			fmt.Printf("这是一个data: %v\n", el)
		default:
			fmt.Printf("这啥也不是: %v\n", el)
		}
	}

	//复制数据切片至空接口切片
	var dataSlice []Data = []Data{{}, {}}
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	// interfaceSlice = dataSlice // 无法编译，它们俩在内存中的布局是不一样的
	// 只能通过for-range 语句显式赋值
	for i, data := range dataSlice {
		interfaceSlice[i] = data
	}

	// 只要实现某接口的方法，就可以将其赋值给某接口变量
	var car = new(Car)
	var en Engine
	en = car
	var empty any
	empty = en
	fmt.Println(empty)
}

// TreeNode 使用空接口定义二叉树数据结构
type TreeNode struct {
	l    *TreeNode
	data any
	r    *TreeNode
}

func (t *TreeNode) SetData(data any) {
	t.data = data
}

func NewTreeNode(l *TreeNode, data any, r *TreeNode) *TreeNode {
	return &TreeNode{l: l, data: data, r: r}
}
