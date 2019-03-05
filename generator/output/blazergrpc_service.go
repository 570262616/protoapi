package output

import (
	"strings"
	"unicode"

	"github.com/yoozoo/protoapi/generator/data"
)

type blazerGrpcMethod struct {
	*data.Method
	ServiceName string
}

func (m *blazerGrpcMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

func (s *blazerGrpcService) Path() string {
	return s.Methods[0].Package + "." + s.Name
}

func (m *blazerGrpcMethod) ServiceType() string {
	if servType, ok := m.Options[data.MethodOptions[data.ServiceTypeMethodOption].Name]; ok {
		return servType
	}

	return "POST"
}

type blazerGrpcService struct {
	*data.ServiceData
	Package string
	Methods []*blazerGrpcMethod
}

func newBlazerGrpcService(msg *data.ServiceData, packageName string) *blazerGrpcService {
	o := &blazerGrpcService{
		msg,
		packageName,
		nil,
	}
	o.init()
	return o
}

func (s *blazerGrpcService) init() {
	s.Methods = make([]*blazerGrpcMethod, len(s.ServiceData.Methods))
	for i, f := range s.ServiceData.Methods {
		mtd := f
		s.Methods[i] = &blazerGrpcMethod{mtd, s.Name}
	}
}

func (s *blazerGrpcService) AlsImports() (result string) {
	var imports []string

	for _, m := range s.Methods {

		var input string
		var output string
		// if !IsContainsPoint(m.InputType) {
		// 	in := strings.Split(m.InputType, ".")
		// 	i := len(in)

		// 	in = append(in, "")
		// 	copy(in[i:], in[i-1:])
		// 	in[i-1] = "message"
		// 	newInputType := strings.Join(in, ".")

		// 	input = "import" + " " + "com.als.grpc." + newInputType + ";"
		// } else {
		// 	input = "import" + " " + "com.als.grpc." + m.Package + "." + m.InputType + ";"
		// }

		if !IsContainsPoint(m.InputType) {
			in := strings.Split(m.InputType, ".")
			i := len(in)

			in = append(in, "")
			copy(in[i:], in[i-1:])
			p := in[i-2]
			in[i-1] = Ucfirst(p) + "Public"
			newInputType := strings.Join(in, ".")

			input = "import" + " " + "com.ezbuy.blazer.api." + newInputType + ";"
		} else {
			input = "import" + " " + "com.ezbuy.blazer.api." + m.Package + "." + s.Name + "Public." + m.InputType + ";"
		}

		if !IsContainsPoint(m.OutputType) {

			in := strings.Split(m.OutputType, ".")
			i := len(in)

			in = append(in, "")
			copy(in[i:], in[i-1:])
			p := in[i-2]
			in[i-1] = Ucfirst(p) + "Public"
			newOutputType := strings.Join(in, ".")

			output = "import" + " " + "com.ezbuy.blazer.api." + newOutputType + ";"
		} else {
			output = "import" + " " + "com.ezbuy.blazer.api." + m.Package + "." + s.Name + "Public." + m.OutputType + ";"
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

func (s *blazerGrpcService) AlsPackage() (result string) {
	return "com.ezbuy.blazer.server.service." + s.Methods[0].Package
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
