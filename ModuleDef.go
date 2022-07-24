package py3_go

type PyModuleDef struct {
	Name       string
	Doc        string
	MethodDefs []*PyMethodDef
}

func NewModuleDef(name string, doc string) *PyModuleDef {
	return &PyModuleDef{Name: name, Doc: doc}
}
func (m *PyModuleDef) AddMethodDef(methodDef *PyMethodDef) {
	m.MethodDefs = append(m.MethodDefs, methodDef)
}
