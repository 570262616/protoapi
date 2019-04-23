package output

import (
	"bytes"
	"strings"
	"text/template"
	"unicode"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/util"
)

var javaTypes = map[string]string{
	// https://developers.google.com/protocol-buffers/docs/proto#scalar
	"double":   "double",
	"float":    "float",
	"int32":    "int",
	"int64":    "long",
	"uint32":   "int",
	"uint64":   "long",
	"sint32":   "int",
	"sint64":   "long",
	"fixed32":  "int",
	"fixed64":  "long",
	"sfixed32": "int",
	"sfixed64": "long",
	"bool":     "boolean",
	"string":   "String",
	"bytes":    "ByteString",
}

var wrapperTypes = map[string]string{
	"double":   "Double",
	"float":    "Float",
	"int32":    "Integer",
	"int64":    "Long",
	"uint32":   "Integer",
	"uint64":   "Long",
	"sint32":   "Integer",
	"sint64":   "Long",
	"fixed32":  "Integer",
	"fixed64":  "Long",
	"sfixed32": "Integer",
	"sfixed64": "Long",
	"bool":     "Boolean",
	"string":   "String",
	"bytes":    "Byte",
	"int":      "Integer",
}

func toJavaType(dataType string, label string) string {
	// check if the field is repeated
	if label == data.FieldRepeatedLabel {
		// check if wrapper type
		if wrapperType, ok := wrapperTypes[dataType]; ok {
			return "List<" + wrapperType + ">"
		}
		return "List<" + dataType + ">"
	}
	// if not repeated filed
	// check if primary type
	if primaryType, ok := javaTypes[dataType]; ok {
		return primaryType
	}
	// if not primary type return data type and ignore the . in the data type
	return dataType
}

type springGen struct {
	ApplicationName string
	PackageName     string
	structTpl       *template.Template
	serviceTpl      *template.Template
	enumTpl         *template.Template
}

func (g *springGen) getTpl(path string) *template.Template {

	var funcs = template.FuncMap{
		"toLower":         Lcfirst,
		"removePkgName":   RemovePkgName,
		"isContainsPoint": IsContainsPoint,
	}
	var err error
	tpl := template.New("tpl").Funcs(funcs)
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		panic(err)
	}
	return result
}

func (g *springGen) init(applicationName, packageName string) {
	g.ApplicationName = applicationName
	g.PackageName = packageName
	g.structTpl = g.getTpl("/generator/template/ezSpring/spring_struct.gojava")
	g.serviceTpl = g.getTpl("/generator/template/ezSpring/spring_service.gojava")
	g.enumTpl = g.getTpl("/generator/template/ezSpring/spring_enum.gojava")
}

func (g *springGen) getStructFilename(packageName string, msg *data.MessageData, service *data.ServiceData) string {
	data.FlattenLocalPackage(msg)

	var name = strings.Replace(packageName, ".", "/", -1) + "/message/" + msg.Name + ".java"

	return name
}

func (g *springGen) genEnum(enum *data.EnumData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringEnum(enum, g.PackageName)
	err := g.enumTpl.Execute(buf, obj)
	if err != nil {
		util.Die(err)
	}

	return buf.String()
}

func (g *springGen) genStruct(msg *data.MessageData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringStruct(msg, g.PackageName)
	err := g.structTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *springGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func genSpringPackageName(packageName string, options data.OptionMap) string {
	if javaPckName, ok := options[data.JavaPackageOption]; ok {
		return javaPckName
	}

	return packageName
}

func (g *springGen) genServiceFileName(packageName string, service *data.ServiceData) string {
	return strings.Replace(packageName, ".", "/", -1) + "/api/" + service.Name + "ServiceAPI.java"
}

func (g *springGen) getEnumFilename(packageName string, enum *data.EnumData) string {
	// return packageName + "/" + enum.Name + ".java"
	var name = strings.Replace(packageName, ".", "/", -1) + "/message/" + enum.Name + ".java"

	return name
}

func (g *springGen) Init(request *plugin.CodeGeneratorRequest) {
}

func (g *springGen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	var service *data.ServiceData
	if len(services) > 1 {
		// util.Die(fmt.Errorf("found %d services; only 1 service is supported now", len(services)))
		service = services[len(services)-1]
	} else if len(services) == 1 {
		service = services[0]
	}

	// get java package name from options
	packageName = genSpringPackageName(packageName, options)
	g.init(applicationName, packageName)
	result = make(map[string]string)

	for _, msg := range messages {
		filename := g.getStructFilename(packageName, msg, service)
		content := g.genStruct(msg)

		newFlieName := strings.Replace(filename, ".java", "", -1)

		if !strings.Contains(newFlieName, ".") {
			result[filename] = content
		}
	}

	for _, enum := range enums {
		filename := g.getEnumFilename(g.PackageName, enum)
		content := g.genEnum(enum)

		if strings.Contains(g.PackageName, enum.Package) {
			result[filename] = content
		}
	}

	// make file name same as java class name
	filename := g.genServiceFileName(packageName, service)
	content := g.genServie(service)
	result[filename] = content

	return
}

func init() {
	data.OutputMap["spring"] = &springGen{}
}

func Lcfirst(str string) string {

	index := strings.Index(str, ".")
	l := strings.Count(str, "") - 1

	n := substring(str, index+1, l)

	for i, v := range n {
		return string(unicode.ToLower(v)) + n[i+1:]
	}
	return ""
}

func RemovePkgName(str string) string {

	if IsContainsPoint(str) {
		return str
	}

	pointIndex := strings.Index(str, ".")
	l := strings.Count(str, "") - 1

	index := strings.Index(str, "<")

	n1 := substring(str, 0, index+1)
	n2 := substring(str, pointIndex+1, l)

	return n1 + n2
}

func IsContainsPoint(str string) bool {
	return !strings.Contains(str, ".")
}

func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
