package implcli

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
)

var pbTypeMap = map[string]string{
	"float32": "float",
	"float64": "double",
	"int8":    "int32",
	"uint8":   "uint32",
	"int16":   "int32",
	"uint16":  "uint32",
	"wstring": "string",
}

type PbEnv interface {
	AddMsg(msg *ros.Message)
	AddSrv(srv *ros.Service)
	Output() map[string]string
}
type pbEnvImpl struct {
	packages       map[string]map[string]*ros.MsgDefine
	packageDepends map[string]map[string]bool
	needOutputMsg  []*ros.Message
	needOutputSrv  []*ros.Service
}

func NewPbEnv() PbEnv {
	return &pbEnvImpl{
		packages:       make(map[string]map[string]*ros.MsgDefine),
		packageDepends: make(map[string]map[string]bool),
	}
}

func (i *pbEnvImpl) AddMsg(msg *ros.Message) {
	i.addMsg(msg.Msg)
	i.needOutputMsg = append(i.needOutputMsg, msg)
}

func (i *pbEnvImpl) addMsg(m *ros.MsgDefine) {
	var f func(msg *ros.MsgDefine)
	f = func(msg *ros.MsgDefine) {
		if _, ok := i.packages[msg.Package]; !ok {
			i.packages[msg.Package] = make(map[string]*ros.MsgDefine)
		}
		i.packages[msg.Package][msg.Name] = msg
		for _, field := range msg.Block.Fields {
			if field.IsComplex {
				f(field.ComplexType)
				if msg.Package == field.ComplexType.Package {
					continue
				}
				if i.packageDepends[msg.Package] == nil {
					i.packageDepends[msg.Package] = make(map[string]bool)
				}
				i.packageDepends[msg.Package][field.ComplexType.Package] = true
			}
		}
	}
	f(m)
}

func (i *pbEnvImpl) AddSrv(srv *ros.Service) {
	i.addMsg(&ros.MsgDefine{
		Package: srv.Srv.Package,
		Name:    srv.Srv.Name + "Req",
		Block:   srv.Srv.Req,
	})
	i.addMsg(&ros.MsgDefine{
		Package: srv.Srv.Package,
		Name:    srv.Srv.Name + "Rsp",
		Block:   srv.Srv.Rsp,
	})
	i.needOutputSrv = append(i.needOutputSrv, srv)
}

func (i *pbEnvImpl) Output() map[string]string {
	res := make(map[string]string, len(i.packages))
	for pName, msgs := range i.packages {
		res[pName] += "syntax = \"proto3\";\n"
		res[pName] += fmt.Sprintf("package %s;\n", pName)
		res[pName] += "option go_package = \"./pb\";\n"
		for k := range i.packageDepends[pName] {
			res[pName] += fmt.Sprintf("import \"%s.proto\";\n", k)
		}
		res[pName] += "\n"
		for name, m := range msgs {
			res[pName] += fmt.Sprintf("message %s {\n", name)
			for i, field := range m.Block.Fields {
				t := field.Type
				if pbType, ok := pbTypeMap[field.Type]; ok {
					t = pbType
				}
				if field.IsComplex && field.ComplexType.Package != pName {
					t = field.ComplexType.Package + "." + t
				}
				fieldStr := fmt.Sprintf("%s %s = %d;\n", t, field.Name, i+1)
				if field.IsArray {
					fieldStr = "repeated " + fieldStr
				}
				res[pName] += "\t" + fieldStr
			}
			res[pName] += fmt.Sprintf("}\n\n")
		}
	}

	mainContent := "syntax = \"proto3\";\noption go_package = \"./pb\";\n"
	for pName := range i.packages {
		mainContent += fmt.Sprintf("import \"%s.proto\";\n", pName)
	}
	mainContent += "\n"

	mainContent += "service Api {\n"
	for _, msg := range i.needOutputMsg {
		mainContent += fmt.Sprintf("\t// topic: %s\n", msg.Topic)
		mainContent += fmt.Sprintf("\trpc PublishTopic%s(%s) returns (%s);\n", Topic2Name(msg.Topic), msg.Msg.Package+"."+msg.Msg.Name, msg.Msg.Package+"."+msg.Msg.Name)
	}

	for _, srv := range i.needOutputSrv {
		mainContent += fmt.Sprintf("\t// service: %s\n", srv.Topic)
		mainContent += fmt.Sprintf("\trpc CallService%s(%sReq) returns (%sRsp);\n", Topic2Name(srv.Topic), srv.Srv.Package+"."+srv.Srv.Name, srv.Srv.Package+"."+srv.Srv.Name)
	}
	mainContent += "}\n"

	res["main"] = mainContent
	return res
}

func Topic2Name(topic string) string {
	var result strings.Builder
	uppercaseNext := false

	for _, r := range topic { // 去掉开头的 `/`
		if r == '/' {
			uppercaseNext = true
		} else {
			if uppercaseNext {
				result.WriteRune(unicode.ToUpper(r))
				uppercaseNext = false
			} else {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

func Name2Topic(name string) string {
	var result strings.Builder
	result.WriteString("/")
	for _, r := range name {
		if unicode.IsUpper(r) {
			if result.Len() > 1 { // 不是第一个字母
				result.WriteRune('/')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
