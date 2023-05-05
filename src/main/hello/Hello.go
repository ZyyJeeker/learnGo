package hello

func init() {
	println("Hello.go 的 init 函数执行")
}

func PrintHello() {
	println("Hello World")
}

type Say interface {
	SayHello()
	sayHello()
}

// Hello 一个简单的结构体
type Hello struct {
	Number int "可以导出到外部包"
	number int "不可以导出到外部包"
}

// SayHello 可以导出到外部包
func (h *Hello) SayHello() {
	println("调用了Hello的SayHello方法")
}

// 不可以导出到外部包
func (h *Hello) sayHello() {
	println("调用了Hello的sayHello方法")
}
