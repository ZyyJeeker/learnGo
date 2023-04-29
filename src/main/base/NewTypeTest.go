package base

import "fmt"

func testNewStudent() {
	tom := Student{"Tom", 12}
	jack := new(Student)
	jack.Age = 14
	jack.Name = "Jack"
	println(jack)
	fmt.Println(tom)
	fmt.Println(jack)
	fmt.Println(*jack)

	println(jack.Name)
	println((*jack).Name)
	println(tom.Name)

	fmt.Println(getStudent("Coco", 34))
	fmt.Println(getStudent0("Coco", 34))
	fmt.Println(getStudent1("Coco", 34))
	fmt.Println(getStudent2("Coco", 34))
}

func getStudent(name string, age int) Student {
	res := new(Student)
	res.Age = age
	res.Name = name
	return *res
}

func getStudent0(name string, age int) *Student {
	res := new(Student)
	res.Age = age
	res.Name = name
	return res
}

func getStudent1(name string, age int) Student {
	return Student{name, age}
}

func getStudent2(name string, age int) *Student {
	return &Student{name, age}
}
