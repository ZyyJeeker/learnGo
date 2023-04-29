package structTest

import (
	"fmt"
	"strings"
)

type Phone struct {
	cpu     string
	ram     string
	price   float64
	version int
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

	// 只初始化部分字段
	fmt.Println(squ2)
	phone1 := &Phone{
		cpu: "Kire9000",
		ram: "12g",
	}
	fmt.Println(phone1)
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
