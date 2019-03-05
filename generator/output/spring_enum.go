package output

import (
	"fmt"
	"strings"

	"github.com/yoozoo/protoapi/generator/data"
)

type springEnumField struct {
	data.EnumField
}

func newSpringEnum(msg *data.EnumData, packageName string) *springEnum {
	ss := strings.Split(packageName, "/")
	s := ss[len(ss)-1]
	o := &springEnum{
		msg,
		s,
		nil,
	}
	o.init()
	return o
}

type springEnum struct {
	*data.EnumData
	Package string
	Fields  []*springEnumField
}

func (s *springEnum) init() {
	s.Fields = make([]*springEnumField, len(s.EnumData.Fields))
	for i, f := range s.EnumData.Fields {
		s.Fields[i] = &springEnumField{f}
	}
}

func (s *springEnum) EnumFields() string {
	var fields []string

	for _, f := range s.Fields {
		field := f.Name + "(" + fmt.Sprintf("%d", f.Value) + ", " + "\"" + f.Comment + "\"" + ")"
		fields = append(fields, field)
	}
	return strings.Join(fields, ",\n\t")
}
