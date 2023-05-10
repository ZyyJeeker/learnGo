package IO

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func IOTest() {
	// Scan 和以 Scan 开头的函数是从控制台读取
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName) // 读取一行，用空格隔开
	fmt.Printf("Your name is %s %s\n", firstName, lastName)
	fmt.Scanf("My name is %s %s", &firstName, &lastName) // 参数化格式读取
	fmt.Printf("Your name is %s %s\n", firstName, lastName)

	// Sscan 和以 Sscan 开头的函数则是从字符串读取
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
