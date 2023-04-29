package main

import "fmt"

func forRange() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for i, season := range seasons {
		fmt.Printf("第%d个季节为：%s", i, season)
	}
	items := []int{1, 2, 3}
	for _, item := range items {
		//第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；
		//他们都是仅在 for 循环内部可见的局部变量。value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。
		item *= 2
	}
	fmt.Println(items)
}

func sliceReCon() {
	items := make([]int, 2, 5)
	for i := range items {
		items[i] = (i + 1) * 2
	}
	fmt.Println(items) //当前只有两个值
	// 切片重组，给末尾加两个元素
	var idx = len(items)
	items = items[0 : len(items)+2]
	items[idx] = 8
	items[idx+1] = 8
	fmt.Println(items)
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println(a)
	a = a[0:4]
	fmt.Println(a)
}

func sliceCopy() {
	sl1 := []int{1, 2, 3}
	sl2 := make([]int, 10)
	n := copy(sl2, sl1)
	fmt.Println(sl2)
	fmt.Println(n)

	sl3 := []int{3, 4, 5}
	sl3 = append(sl3, 5, 6)
	fmt.Println(sl3)
	sl4 := []int{8, 9}
	sl3 = append(sl3, sl4...)
	fmt.Println(sl3)
}
