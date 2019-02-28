package output

import (
	"bytes"
	"log"
	"text/template"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
)

type blazerGrpcGen struct {
	ApplicationName string
	PackageName     string
	serviceTpl      *template.Template
}

func (g *blazerGrpcGen) getTpl(path string) *template.Template {

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

func (g *blazerGrpcGen) init(applicationName, packageName string) {
	g.ApplicationName = applicationName
	g.PackageName = packageName
	g.serviceTpl = g.getTpl("/generator/template/ezJavaGrpc/blazer_service.gojava")
}

func (g *blazerGrpcGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newBlazerGrpcService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func genBlazerGrpcPackageName(packageName string, options data.OptionMap) string {
	if javaPckName, ok := options[data.JavaPackageOption]; ok {
		return javaPckName
	}

	return packageName
}

func (g *blazerGrpcGen) genServiceFileName(packageName string, service *data.ServiceData) string {
	return "com/blazer/gateway/grpc/service/" + service.Methods[0].Package + "/" + service.Name + "Service.java"
}

func (g *blazerGrpcGen) Init(request *plugin.CodeGeneratorRequest) {
}

func (g *blazerGrpcGen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
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
	data.OutputMap["blazer"] = &blazerGrpcGen{}
}
