package structTest

import (
	"fmt"
	"math"
)

type Python struct {
	Code
	DeBug
}

func (p *Python) code() {
	println("Use python code.")
}

func (p *Python) deBug() {
	println("Use python debug.")
}

type Developer struct {
	code  Code
	debug DeBug
}

func (d *Developer) SetCode(code Code) {
	d.code = code
}

func (d *Developer) SetDebug(debug DeBug) {
	d.debug = debug
}

// InterfaceTest 使用接口实现多态
func InterfaceTest() {
	python := new(Python)
	java := &Java{1, "java"}

	developer := new(Developer)
	developer.SetCode(java)
	developer.SetDebug(java)

	println("#######开发人员使用java开发########")
	developer.code.code()
	developer.debug.deBug()

	developer.SetDebug(python)
	developer.SetCode(python)

	println("#######开发人员使用Python开发########")
	developer.code.code()
	developer.debug.deBug()

	println("######################")
	sq1 := new(Square)
	sq1.side = 5.2

	var areaInf Shaper
	areaInf = sq1
	fmt.Printf("The square has area: %f\n", areaInf.Area())

	r := &Rectangle{1.2, 3.4}
	q := &Square{3.4}

	// 接口类型断言：如果areaInf是*Square类型则将areaInf赋值给v
	if v, ok := areaInf.(*Square); ok {
		fmt.Println(v)
	}

	shapers := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for _, shaper := range shapers {
		fmt.Println("Shape details: ", shaper)
		fmt.Println("Area of this shape is: ", shaper.Area())
	}

	var areaType Shaper = &Cycle{1.2}

	if v, ok := areaType.(*Cycle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", v)
	}

	if v, ok := areaType.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", v)
	}

	// type-switch
	switch t := areaType.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Cycle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	// 不需要赋值
	switch areaType.(type) {
	case *Square:
		println("这是一个Square")
	case *Cycle:
		println("这是一个Cycle")
	default:
		println("啥也不是")
	}

	interfaceSetTest()

	r1 := Rectangle{5, 3} // Area() of Rectangle needs a value
	q1 := &Square{5}      // Area() of Square needs a pointer

	shapers = []Shaper{&r1, q1}
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

type Cycle struct {
	radius float32
}

func (c *Cycle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

// Programmer 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样
type Programmer interface {
	DeBug
	Code
	Square
	Engine
}

type List []int

type Liner interface {
	Lin() int
}

type Appender interface {
	Append(int)
}

func (l List) Lin() int {
	return len(l)
}

func (l *List) Append(i int) {
	*l = append(*l, i)
}

func CountInto(a Appender, start, end int) {
	println(a)
	fmt.Println(a)
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Liner) bool {
	println(l)
	fmt.Println(l)
	return l.Lin()*10 > 42
}

func interfaceSetTest() {
	var lst List // 相当于 lst := []int{}
	CountInto(&lst, 1, 10)
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
