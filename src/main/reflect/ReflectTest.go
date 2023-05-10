package reflect

import (
	"fmt"
	. "learnGo/structTest"
	"os"
	"reflect"
	"strconv"
)

func ReflectTest() {
	person := Person{
		FirstName: "张",
		LastName:  "宇",
	}

	// TypeOf饭回类型，ValueOf饭回值，Kind返回底层类型

	// 结构体的类型和值
	println("结构体的类型和值######################")
	personType := reflect.TypeOf(person)
	fmt.Println(personType)
	fmt.Println(personType.Kind())
	personValue := reflect.ValueOf(person)
	fmt.Println(personValue)
	fmt.Println(personValue.Kind())

	p := new(Person)
	p.FirstName = "胡"
	p.LastName = "图图"
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(p).Kind())
	fmt.Println(reflect.ValueOf(p))
	fmt.Println(reflect.ValueOf(p).Kind())

	println("基本类型的类型和值######################")

	// 基本类型的类型和值
	fmt.Println(reflect.TypeOf(1.23))
	fmt.Println(reflect.TypeOf(1.23).Kind())
	fmt.Println(reflect.ValueOf(1.23))
	fmt.Println(reflect.ValueOf(1.23).Kind())

	fmt.Println(reflect.TypeOf("good"))
	fmt.Println(reflect.TypeOf("good").Kind())
	fmt.Println(reflect.ValueOf("good"))
	fmt.Println(reflect.ValueOf("good").Kind())

	println("######################")

	// 函数的类型和值
	fmt.Println(reflect.TypeOf(ReflectTest))
	fmt.Println(reflect.TypeOf(ReflectTest).Kind())
	fmt.Println(reflect.ValueOf(ReflectTest))
	fmt.Println(reflect.ValueOf(ReflectTest).Kind())

	println("######################")

	// 接口,方法的类型和值
	var code Code
	gopher := Gopher{}
	code = gopher
	fmt.Println(reflect.TypeOf(code))
	fmt.Println(reflect.TypeOf(code).Kind())
	fmt.Println(reflect.ValueOf(code))
	fmt.Println(reflect.ValueOf(code).Interface())
	fmt.Println(reflect.ValueOf(code).Kind())

	// 使用反射修改值
	p1 := new(Person)
	p1.FirstName = "胡"
	p1.LastName = "歌"
	v := reflect.ValueOf(p1).Elem() // 获取结构体可修改的值
	field1 := v.Field(1)
	fmt.Println(field1.CanSet())
	field1.SetString("图图")
	fmt.Println(field1)
	fmt.Println(field1.Interface())
	fmt.Println(p1)

	camera := &Camera{"徕卡"}
	elem := reflect.ValueOf(camera).Elem()
	for i := 0; i < elem.NumField(); i++ {
		fmt.Println(elem.Field(i).Interface())
	}

	for i := 0; i < elem.NumMethod(); i++ {
		elem.Method(i).Call(nil)
	}

	printLn(12)
	printLn("12")

	println()

	duck := Bird{"麻雀"}
	fmt.Println(duck)
	TestInterfaceParam(duck, &duck)
	fmt.Println(duck)

	println("################")
	// 方法的receiver本质就是函数的入参，也分值传递和指针传递
	b := Bird{}
	b.Quack()
	fmt.Println(b)
	b.Walk()
	fmt.Println(b)

	println("################")

	c := &Bird{}
	c.Quack()
	fmt.Println(c)
	c.Walk()
	fmt.Println(c)

	// 方法的receive是值的话，不管对象是值还是指针，传递的都是值的复制

	// 方法的receive是指针的话，不管对象是值还是指针，传递的都是对象指针的复制

	motorcars := Motorcars{
		&Motorcar{Manufacturer: "BMW", BuildYear: 1988},
		&Motorcar{Manufacturer: "Audi", BuildYear: 1952},
		&Motorcar{Manufacturer: "BYD", BuildYear: 2011},
	}

	res := motorcars.Filter(func(t *Motorcar) bool {
		return t.BuildYear > 1980
	}).Process(func(t *Motorcar) {
		t.Model = "自动挡"
	}).Map(func(t *Motorcar) any {
		return t.Manufacturer
	})

	fmt.Println(res)
	fmt.Println(motorcars)

	foo, resMap := MakeSortedAppender([]string{"BMW", "Audi"})
	motorcars.Process(foo)
	fmt.Println(resMap)

	motors := []Motorcar{{"Q6", "Audi", 2002}}   // 切片元素是值的话，不实现String也可以打印出值
	motorPs := []*Motorcar{{"Q6", "Audi", 2002}} // 切片元素是指针的话，必须实现String才可以打印出值,不然只会打印指针
	fmt.Println(motors)
	fmt.Println(motorPs)

	mo := Motorcar{
		Model:        "A6",
		Manufacturer: "Audi",
		BuildYear:    2006,
	}
	fmt.Printf("\n这就是:%+v\n", &mo) // 打印属性结构格式
	fmt.Println(&mo)               // 不实现String的话,打印地址格式,实现String则按照String格式打印
}

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func printLn(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) { // type switch
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		// more types
		default:
			os.Stdout.WriteString("???")
		}
	}
}

type Quack interface {
	Quack()
}

type Walk interface {
	Walk()
}
type Bird struct {
	Name string
}

func (b Bird) Quack() {
	b.Name = "孔雀"
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	b.Name = "啄木鸟"
	fmt.Println("I am walking.")
}

func TestInterfaceParam(quack Quack, walk Walk) {
	if bird, ok := quack.(Bird); ok {
		bird.Name = "百灵鸟" // 传递值 修改的是传参复制的值
		fmt.Println(bird.Name)
	}
	fmt.Println(walk)
	if bird, ok := walk.(*Bird); ok {
		bird.Name = "喜鹊" // 传递指针 修改的是传参来的指针解指针后的值
		fmt.Println(bird.Name)
	}
}

type Motorcar struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

func (m *Motorcar) String() string {
	return "{ " + m.Model + " " + m.Manufacturer + " " + strconv.Itoa(m.BuildYear) + " }"
}

type Motorcars []*Motorcar

func (ts Motorcars) Process(f func(t *Motorcar)) Motorcars {
	for _, t := range ts {
		f(t)
	}
	return ts
}

func (ts Motorcars) Filter(f func(t *Motorcar) bool) Motorcars {
	res := make(Motorcars, 0)

	for _, t := range ts {
		if f(t) {
			res = append(res, t)
		}
	}

	return res
}

func (ts Motorcars) Map(f func(t *Motorcar) any) []any {
	res := make([]any, len(ts))
	idx := 0
	ts.Process(func(tp *Motorcar) {
		res[idx] = f(tp)
		idx++
	})
	return res
}

func MakeSortedAppender(manufacturers []string) (func(c *Motorcar), map[string]Motorcars) {
	resMap := make(map[string]Motorcars)
	for _, manufacturer := range manufacturers {
		resMap[manufacturer] = make(Motorcars, 0)
	}

	resMap["default"] = make(Motorcars, 0)

	foo := func(c *Motorcar) {

		if _, ok := resMap[c.Manufacturer]; ok {
			resMap[c.Manufacturer] = append(resMap[c.Manufacturer], c)
		} else {
			resMap["default"] = append(resMap["default"], c)
		}
	}

	return foo, resMap
}
