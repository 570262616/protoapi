package output

import (
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type springField struct {
	*data.MessageField
}

func (s *springField) Title() string {
	return strings.Title(s.Name)
}

func (s *springField) JavaType() string {
	return toJavaType(s.MessageField.DataType, s.MessageField.Label)
}

func newSpringStruct(msg *data.MessageData, packageName string) *springStruct {
	o := &springStruct{
		msg,
		packageName,
		nil,
	}
	o.init()
	return o
}

type springStruct struct {
	*data.MessageData
	Package string
	Fields  []*springField
}

func (s *springStruct) init() {
	s.Fields = make([]*springField, len(s.MessageData.Fields))
	for i, f := range s.MessageData.Fields {
		s.Fields[i] = &springField{f}
	}
}

func (s *springStruct) ContructParam() string {
	params := make([]string, len(s.Fields))
	for i, f := range s.Fields {
		params[i] = "@JsonProperty(\"" + f.Name + "\") " + f.JavaType() + " " + f.Name
	}
	return strings.Join(params, ", ")
}

func (s *springStruct) ClassName() string {
	return s.Name
}

func (s *springStruct) Imports() string {

	var imports []string
	var typeName string

	for _, f := range s.Fields {

		dataType := customeType(f.DataType)

		if !IsContainsPoint(dataType) {

			in := strings.Split(dataType, ".")
			i := len(in)

			in = append(in, "")
			copy(in[i:], in[i-1:])
			in[i-1] = "message"
			newType := strings.Join(in, ".")

			typeName = "import" + " " + "com.ezbuy.blazer.api." + newType + ";"
		} else {
			typeName = "import" + " " + s.Package + ".message." + dataType + ";"
		}

		if dataType != "" && !in_array(typeName, imports) {
			imports = append(imports, typeName)
		}

	}

	return strings.Join(imports, "\n")

}

func customeType(dataType string) string {
	// if not repeated filed
	// check if primary type
	if _, ok := javaTypes[dataType]; ok {
		return ""
	}
	// if not primary type return data type and ignore the . in the data type
	return dataType
}
