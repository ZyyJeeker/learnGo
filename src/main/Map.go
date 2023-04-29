package main

import (
	"fmt"
	"log"
	"sort"
)

func mapTest() {
	// map被声明
	var m1 map[string]int
	// map被初始化分配
	m1 = make(map[string]int)             // map的创建方式1
	m2 := make(map[int]string)            // map的创建方式2
	m3 := map[string]string{"three": "叁"} // map的创建方式3

	m1["one"] = 1
	m1["two"] = 2
	m2[1] = "one"
	m2[2] = "two"
	m3["one"] = "壹"
	m3["two"] = "贰"

	for key, value := range m1 {
		fmt.Printf("key:%v,value:%v\n", key, value)
	}

	for key, value := range m2 {
		fmt.Printf("key:%v,value:%v\n", key, value)
	}

	for key, value := range m3 {
		fmt.Printf("key:%v,value:%v\n", key, value)
	}

	handler := map[string]func(int, int) int{
		"+": add1,
		"-": sub1,
		"*": mul1,
		"/": div1,
	}

	res := handler["*"](12, 12)
	println(res)
	fmt.Println(handler)

	mp5 := make(map[string][]string)
	mp5["蔬菜"] = []string{"番茄", "窝瓜", "花菜"}
	mp5["水果"] = []string{"苹果", "香蕉", "橘子"}
	mp5["动物"] = []string{"猫"}
	// map返回值，第一个为value，第二个为bool，判断是否存在
	if vag, isPresent := mp5["水果"]; isPresent {
		af := append(vag, "榴莲")
		fmt.Println(af)
	}
	// 修改切片的数据
	mp5["水果"][1] = "西瓜"
	fmt.Println(mp5)

	// 删除map的元素
	delete(mp5, "动物")
	fmt.Println(mp5)

	// 只取key
	for key := range mp5 {
		fmt.Println(key)
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

func sliceMap() {
	items := make([]map[int]int, 3)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 3
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}

func add1(a, b int) int {
	return a + b
}

func sub1(a, b int) int {
	return a - b
}

func mul1(a, b int) int {
	return a * b
}

func div1(a, b int) int {
	return a / b
}

var (
	noteFrequency = map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func sortMap() {
	log.Println("unsorted:")
	fmt.Println(barVal)
	keys := make([]string, len(barVal))
	values := make([]int, len(barVal))
	idx := 0
	for k, v := range barVal {
		keys[idx] = k
		values[idx] = v
		idx++
	}
	sort.Strings(keys)
	sort.Ints(values)
	log.Println("After sorted:")
	fmt.Println(keys)
	fmt.Println(values)
}

func invert() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	fmt.Println(invMap)
}
