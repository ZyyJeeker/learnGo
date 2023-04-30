package structTest

import "fmt"

type code interface {
	code()
}

type deBug interface {
	deBug()
}

type Gopher struct {
	num      int
	language string
}

// 方法的receiver如果是值类型，则传递的是结构体实例的拷贝，修改结构体实例的拷贝不会影响原来的结构体
func (p Gopher) code() {
	p.num++
	fmt.Printf("I am coding %s language, num is %d\n", p.language, p.num)
}

// 方法的receiver如果是指针类型，则传递的是结构体实例指针的拷贝，修改结构体实例的拷贝会影响原来的结构体
func (p *Gopher) deBug() {
	p.num++
	fmt.Printf("I am debuging %s language, num is %d\n", p.language, p.num)
}

func MethodTest() {
	gopher1 := Gopher{1, "golang"}
	fmt.Printf("原来的gopher1为:%v", gopher1)
	gopher1.code()
	fmt.Printf("原来的gopher1为:%v", gopher1)
	gopher1.deBug()
	fmt.Printf("原来的gopher1为:%v", gopher1)
	println("==================")
	gopher2 := &Gopher{2, "java"}
	fmt.Printf("原来的gopher2为:%v", gopher2)
	gopher2.code()
	fmt.Printf("原来的gopher2为:%v", gopher2)
	gopher2.deBug()
	fmt.Printf("原来的gopher2为:%v", gopher2)

	println("=================")

	java := Java{8, "java"}
	java.code()
	java.deBug()

	aliasMe()
}

type Java struct {
	num      int
	language string
}

// 不同结构体的同名方法可以重载
func (p Java) code() {
	p.num++
	fmt.Printf("I am coding %s language, num is %d\n", p.language, p.num)
}

func (p *Java) deBug() {
	p.num++
	fmt.Printf("I am debuging %s language, num is %d\n", p.language, p.num)
}

// 非结构体类型的别民也可以有方法

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

// 别名类型没有原始类型上已经定义过的方法

type Vector IntVector

func (v Vector) Sub() (s int) {
	for _, x := range v {
		s -= x
	}
	return
}

func aliasMe() {
	intVector := IntVector{23, 34, 5}
	println(intVector.Sum())
	vector := Vector{3, 4, 5}
	println(vector.Sub())
}
