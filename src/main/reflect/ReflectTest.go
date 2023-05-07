package reflect

import (
	"fmt"
	. "learnGo/structTest"
	"reflect"
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
}
