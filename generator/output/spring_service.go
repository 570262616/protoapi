package output

import (
	"reflect"
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type springMethod struct {
	*data.Method
	ServiceName string
}

func (m *springMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

func (m *springMethod) ServiceType() string {
	if servType, ok := m.Options[data.MethodOptions[data.ServiceTypeMethodOption].Name]; ok {
		return servType
	}

	return "POST"
}

type springService struct {
	*data.ServiceData
	Package string
	Methods []*springMethod
}

func newSpringService(msg *data.ServiceData, packageName string) *springService {
	o := &springService{
		msg,
		packageName,
		nil,
	}
	o.init()
	return o
}

func (s *springService) init() {
	s.Methods = make([]*springMethod, len(s.ServiceData.Methods))
	for i, f := range s.ServiceData.Methods {
		mtd := f
		s.Methods[i] = &springMethod{mtd, s.Name}
	}
}

func (s *springService) Imports() (result string) {
	var imports []string

	for _, m := range s.Methods {

		var input string
		var output string
		if !IsContainsPoint(m.InputType) {
			input = "import" + " " + m.InputType
		} else {
			input = "import" + " " + m.Package + ".message." + m.InputType
		}

		if !IsContainsPoint(m.OutputType) {
			output = "import" + " " + m.OutputType
		} else {
			output = "import" + " " + m.Package + ".message." + m.OutputType
		}

		if !in_array(input, imports) {
			imports = append(imports, input)
		}
		if !in_array(output, imports) {
			imports = append(imports, output)
		}
	}

	result = strings.Join(imports, "\n")

	return
}

func in_array(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}

// func IsContainsPoint(str string) bool {
// 	return !strings.Contains(str, ".")
// }
