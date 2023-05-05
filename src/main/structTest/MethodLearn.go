package structTest

import (
	"fmt"
	. "learnGo/hello" // 前面加了点号.后引用该包内类型或者方法时不需要包名
	"math"
	"runtime"
)

type Code interface {
	code()
}

type DeBug interface {
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
	fmt.Print()

	java := Java{8, "java"}
	java.code()
	java.deBug()

	aliasMe()

	car := new(Car)
	car.GoToWork()

	n := &NamedPoint{Point: Point{3, 4}, name: "Pythagoras"}
	fmt.Println(n.Abs()) // 打印5
	n.SayHello()

	cp := &CameraPhone{
		Camera:      Camera{"尼康"},
		MobilePhone: MobilePhone{"小灵通"},
	}

	println(cp.Call())
	println(cp.TakeAPicture())

	fmt.Printf("two1 is: %v\n", cp)
	fmt.Println("two1 is:", cp)
	fmt.Printf("two1 is: %T\n", cp)
	fmt.Printf("two1 is: %#v\n", cp)
	println(cp.String())

	// 获取运行时内存状态
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
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

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	Gopher
}

func (car *Car) GoToWork() {
	//car.Start()		// 未实现的方法不能直接调用
	//car.Stop()
	println("########################")
	car.deBug()
	car.code()
}

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) sub() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point // 匿名类型一般作为被继承的对象，模拟经典面向对象语言中的子类和继承相关的效果
	Hello
	name string
}

// Abs 和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。
func (p *NamedPoint) Abs() float64 {
	return p.Point.Abs() * 100
}

type Camera struct {
	name string
}

func (camera *Camera) String() string {
	return "Camera name is " + camera.name
}

func (camera *Camera) TakeAPicture() string {
	return "Click"
}

type MobilePhone struct {
	name string
}

func (m *MobilePhone) String() string {
	return "MobilePhone name is " + m.name
}

func (m *MobilePhone) Call() string {
	return "Ring Ring"
}

// CameraPhone 通过匿名内嵌实现多继承
type CameraPhone struct {
	Camera
	MobilePhone
	DeBug
}

func (cp *CameraPhone) String() string {
	return cp.MobilePhone.String() + "\n" + cp.Camera.String()
}
