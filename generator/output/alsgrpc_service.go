package output

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type alsGrpcMethod struct {
	*data.Method
	ServiceName string
}

func (m *alsGrpcMethod) Path() string {
	return "/" + m.ServiceName + "." + m.Name
}

func (s *alsGrpcService) Path() string {
	return s.Methods[0].Package + "." + s.Name
}

func (m *alsGrpcMethod) ServiceType() string {
	if servType, ok := m.Options[data.MethodOptions[data.ServiceTypeMethodOption].Name]; ok {
		return servType
	}

	return "POST"
}

type alsGrpcService struct {
	*data.ServiceData
	Package string
	Methods []*alsGrpcMethod
}

func newAlsGrpcService(msg *data.ServiceData, packageName string) *alsGrpcService {
	o := &alsGrpcService{
		msg,
		packageName,
		nil,
	}
	o.init()
	return o
}

func (s *alsGrpcService) init() {
	s.Methods = make([]*alsGrpcMethod, len(s.ServiceData.Methods))
	for i, f := range s.ServiceData.Methods {
		mtd := f
		s.Methods[i] = &alsGrpcMethod{mtd, s.Name}
	}
}

func (s *alsGrpcService) AlsImports() (result string) {
	var imports []string

	imports = append(imports, "import"+" "+"com.als.grpc."+s.ServiceData.Methods[0].Package+".Definition"+";")

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

		input = "import" + " " + "com.als.grpc." + m.Package + "." + s.Name + "OuterClass." + m.InputType + ";"

		// if !IsContainsPoint(m.OutputType) {

		// 	in := strings.Split(m.OutputType, ".")
		// 	i := len(in)

		// 	in = append(in, "")
		// 	copy(in[i:], in[i-1:])
		// 	in[i-1] = "message"
		// 	newOutputType := strings.Join(in, ".")

		// 	output = "import" + " " + "com.als.grpc." + newOutputType + ";"
		// } else {
		// 	output = "import" + " " + "com.als.grpc." + m.Package + "." + m.OutputType + ";"
		// }

		output = "import" + " " + "com.als.grpc." + m.Package + "." + s.Name + "OuterClass." + m.OutputType + ";"

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

func (s *alsGrpcService) AlsPackage() (result string) {
	return "com.als.gateway.routing.zuul.grpc.service." + s.Methods[0].Package
}
