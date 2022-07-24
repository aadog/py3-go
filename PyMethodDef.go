package py3

type PyMethodDef struct {
	Name   string
	Method func(self *PyObject, args *PyObject) *PyObject
	flags  int
	Doc    string
}

func NewMethodDef(name string, meth func(self *PyObject, args *PyObject) *PyObject, doc string, flags int) *PyMethodDef {
	return &PyMethodDef{
		Name:   name,
		Method: meth,
		flags:  flags,
		Doc:    doc,
	}
}
