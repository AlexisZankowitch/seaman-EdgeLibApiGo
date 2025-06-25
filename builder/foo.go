package builder

type MyFoo struct {
	namespace     string
	componentName string
	moduleName    string
	moduleVersion string
}

func (s *MyFoo) GetNameSpace() string {
	return s.namespace
}
