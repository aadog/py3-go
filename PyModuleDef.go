package py3

type PyModuleDef struct {
	Name       string
	Doc        string
	MethodDefs []*PyMethodDef
}

func NewModuleDef(name string, doc string) *PyModuleDef {
	def := &PyModuleDef{Name: name, Doc: doc}
	return def
}
func (m *PyModuleDef) AddMethodDef(methodDef *PyMethodDef) {
	m.MethodDefs = append(m.MethodDefs, methodDef)
}
