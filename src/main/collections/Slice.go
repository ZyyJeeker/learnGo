package collections

import "fmt"

func SliceTest() {
	// 切片和数组在内存的结构不同，使用时注意区分
	// 数组的声明一定是要指定长度，切片则不需要
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := []string{"a", "b", "c", "d"} // 切片
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	var c [4]string
	for i := 0; i <= 3; i++ {
		c[i] = string(rune('a' + i))
	}
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	// 声明了一个数组，此处的...编译时直接转换为5，所以等同于声明的数组
	var arr = [...]int{0, 1, 2, 3, 4}
	// 声明了一个数组
	var arr1 = [5]int{0, 1, 2, 3, 4}
	// 声明了一个切片
	var arr2 = []int{0, 1, 2, 3, 4}
	// 声明一个切片
	var arr3 = make([]int, 5)

	// 直接使用切片
	Sum(arr3)
	Sum(arr2)
	// 数组转换为切片
	Sum(arr1[:])
	Sum(arr[:])

	// 使用切片的指针
	Sum0(&arr2)
	Sum0(&arr3)
	its := arr[:]
	Sum0(&its)

	// 使用数组
	Sum1(arr)
	Sum1(arr1)

	// 使用数组指针
	Sum2(&arr)
	Sum2(&arr1)
}

// Sum 接受的参数是一个切片
func Sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// Sum0 接受的参数是一个切片指针
func Sum0(a *[]int) int {
	// 切片指针不支持索引，需要先解指针
	b := *a
	s := 0
	for i := 0; i < len(b); i++ {
		s += b[i]
	}
	return s
}

// Sum1 接受的参数是一个数组
func Sum1(a [5]int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// Sum2 接受的参数是一个数组指针
func Sum2(a *[5]int) int {
	// 数组指针支持索引
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
