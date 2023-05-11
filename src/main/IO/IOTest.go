package IO

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

const PATH = "./resource/poem.txt"
const NewPath = "./resource/poem1.txt"

func AllIOTest() {
	ReadFileTest(PATH)
	ReadFileInByte(PATH)
	WriteFile("\n唐-李白", PATH)
	ReadBytes(PATH)
	_, err := CopyFile(PATH, NewPath)
	ErrHandle(err)
	fmt.Println(os.Args[1:]) // go的main程序参数和
}

func ReadFileTest(filePath string) {
	file, err := os.Open(filePath)
	ErrHandle(err)

	defer CloseFile(file)

	reader := bufio.NewReader(file)
	idx := 1
	for {
		str, err := reader.ReadString('\n')
		fmt.Printf("读取的第%d行为：%s\n", idx, str)
		if err == io.EOF {
			break
		}
		idx++
	}
}

func ErrHandle(err error) {
	if err != nil {
		fmt.Printf("文件操作错误:%s\n", err.Error())
	}
}

func ReadFileInByte(path string) {
	bytes, err := os.ReadFile(path)
	ErrHandle(err)

	fmt.Println(string(bytes))
}

func WriteFile(str, filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)

	defer CloseFile(file)

	ErrHandle(err)

	writer := bufio.NewWriter(file)

	_, r := writer.WriteString(str)
	ErrHandle(r)
	ErrHandle(writer.Flush())
}

func ReadBytes(filePath string) {
	file, err := os.Open(filePath)

	defer CloseFile(file)

	ErrHandle(err)

	reader := bufio.NewReader(file)

	bytes := make([]byte, 1024)

	for {
		_, err := reader.Read(bytes)
		fmt.Println(string(bytes))
		flushBytes(bytes)
		if err == io.EOF {
			break
		}
	}
}

func flushBytes(bytes []byte) {
	for idx, _ := range bytes {
		bytes[idx] = 0
	}
}

func CopyFile(src, dst string) (written int64, err error) {
	file, err := os.Open(src)
	ErrHandle(err)
	defer CloseFile(file)

	newFile, err1 := os.Create(dst)
	ErrHandle(err1)

	return io.Copy(newFile, file)
}

func CloseFile(file *os.File) {
	err := file.Close()
	ErrHandle(err)
}
