package structTest

type Python struct {
	Code
	DeBug
}

func (p *Python) code() {
	println("Use python code.")
}

func (p *Python) deBug() {
	println("Use python debug.")
}

type Developer struct {
	code  Code
	debug DeBug
}

func (d *Developer) SetCode(code Code) {
	d.code = code
}

func (d *Developer) SetDebug(debug DeBug) {
	d.debug = debug
}

// InterfaceTest 使用接口实现多态
func InterfaceTest() {
	python := new(Python)
	java := &Java{1, "java"}

	developer := new(Developer)
	developer.SetCode(java)
	developer.SetDebug(java)

	println("#######开发人员使用java开发########")
	developer.code.code()
	developer.debug.deBug()

	developer.SetDebug(python)
	developer.SetCode(python)

	println("#######开发人员使用Python开发########")
	developer.code.code()
	developer.debug.deBug()
}
