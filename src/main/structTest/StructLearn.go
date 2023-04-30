package structTest

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type Phone struct {
	// 可以给结构体的字段添加tag
	cpu     string  "中央处理器"
	ram     string  "内存"
	price   float64 "价格"
	version int     "版本"
}

func NewPhone(cpu, ram string, price float64, version int) *Phone {
	phone := new(Phone)
	phone.cpu = cpu
	phone.ram = ram
	phone.price = price
	phone.version = version
	fmt.Println(phone)
	return phone
}

func PhoneTest() {
	phone := NewPhone("8Gen2", "12g", 3999.0, 1)
	fmt.Printf("手机为：%v", phone)

	// new(Type) 和 &Type{} 是等价的
	squ1 := &Squ{3, 4}
	fmt.Println(squ1)
	squ2 := new(Squ)
	squ2.width = 4
	squ2.length = 3

	// 该方式等价于 squ3 := Squ{}
	var squ3 Squ
	squ3.width = 3
	squ3.length = 4
	fmt.Printf("Squ3为%v\n", squ3)

	squ4 := Squ{
		length: 3,
		width:  4,
	}

	// 指针类型和结构体的值类型大小不同
	fmt.Printf("\nSqu1的大小为:%v,Squ2的大小为:%v,Squ3的大小为:%v,Squ4的大小为:%v\n",
		unsafe.Sizeof(squ1), unsafe.Sizeof(squ2), unsafe.Sizeof(squ3), unsafe.Sizeof(squ4))

	// 只初始化部分字段
	fmt.Println(squ2)
	phone1 := &Phone{
		cpu: "Kire9000",
		ram: "12g",
	}
	fmt.Println(phone1)

	structTagTest()
	noNameTest()
	ambiguous()
}

type Squ struct {
	length int
	width  int
}

type Person struct {
	firstName string
	lastName  string
}

func PersonTest() {
	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

type number struct {
	f float32
}

type nr number

func StructTransfer() {
	a := number{5.0}
	b := nr{5.0}
	c := number(b)
	fmt.Println(a, b, c)
}

func structTagTest() {
	phone := Phone{}
	for i := 0; i < 4; i++ {
		refTag(phone, i)
	}
}

func refTag(p Phone, idx int) {
	ref := reflect.TypeOf(p)
	field := ref.Field(idx)
	fmt.Printf("字段名为%v,字段类型为%v,字段的位置为%v,字段的tag为%v\n", field.Name, field.Type, field.Index, field.Tag)
}

type NoName struct {
	F   string "具名字段"
	Squ "匿名字段"
}

func noNameTest() {
	name := &NoName{
		F:   "HK",
		Squ: Squ{1, 2},
	}
	fmt.Println(name)
}

type A struct{ a int }
type B struct{ a, b int }
type C struct {
	A
	B
}

func ambiguous() {
	var c C = C{
		A{1},
		B{
			2, 3,
		},
	}
	println(c.A.a)
	println(c.B.a)
	// println(c.a)	// 编译错误
}
