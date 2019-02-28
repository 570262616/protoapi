package output

import (
	"bytes"
	"log"
	"text/template"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
)

type alsGrpcGen struct {
	ApplicationName string
	PackageName     string
	serviceTpl      *template.Template
}

func (g *alsGrpcGen) getTpl(path string) *template.Template {

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

func (g *alsGrpcGen) init(applicationName, packageName string) {
	g.ApplicationName = applicationName
	g.PackageName = packageName
	g.serviceTpl = g.getTpl("/generator/template/ezJavaGrpc/als_service.gojava")
}

// func (g *javaGrpcGen) getStructFilename(packageName string, msg *data.MessageData, service *data.ServiceData) string {
// 	data.FlattenLocalPackage(msg)

// 	var name = strings.Replace(packageName, ".", "/", -1) + "/message/" + msg.Name + ".java"

// 	return name
// }

func (g *alsGrpcGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newAlsGrpcService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func genAlsGrpcPackageName(packageName string, options data.OptionMap) string {
	if javaPckName, ok := options[data.JavaPackageOption]; ok {
		return javaPckName
	}

	return packageName
}

func (g *alsGrpcGen) genServiceFileName(packageName string, service *data.ServiceData) string {
	return "com/als/gateway/routing/zuul/grpc/service/" + service.Methods[0].Package + "/" + service.Name + "Service.java"
}

func (g *alsGrpcGen) Init(request *plugin.CodeGeneratorRequest) {
}

func (g *alsGrpcGen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	var service *data.ServiceData
	if len(services) > 1 {
		// util.Die(fmt.Errorf("found %d services; only 1 service is supported now", len(services)))
		service = services[len(services)-1]
	} else if len(services) == 1 {
		service = services[0]
	}

	// get java package name from options
	packageName = genAlsGrpcPackageName(packageName, options)
	g.init(applicationName, packageName)
	result = make(map[string]string)

	// make file name same as java class name
	filename := g.genServiceFileName(packageName, service)
	content := g.genServie(service)
	result[filename] = content

	log.Printf(service.Name)

	return
}

func init() {
	data.OutputMap["als"] = &alsGrpcGen{}
}
