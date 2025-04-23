package foxglovetopb

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"

	"github.com/0x587/guardeye/common/foxglovetopb/parser"
	"github.com/0x587/guardeye/common/foxglovetopb/walk"
)

var pbTypeMap = map[string]string{
	"float32": "float",
	"float64": "double",
	"byte":    "uint32",
	"int8":    "int32",
	"uint8":   "uint32",
	"int16":   "int32",
	"uint16":  "uint32",
	"wstring": "string",
}

type PbEnv interface {
	ParseMsg(topic, typename, schema string) error
	ParseSrv(topic, typename, reqSchema, rspSchema string) error
	Output() map[string]string
	OutputForType(t string) (string, string)
}

func New() PbEnv {
	return &impl{
		packageMsgs:    make(map[string]map[string]*MsgDefine),
		packageDepends: make(map[string]map[string]bool),
		msgDepends:     make(map[string]map[string]bool),
		messageApi:     make(map[string]*MsgDefine),
		serviceApi:     make(map[string]*SrvDefine),
	}
}

type impl struct {
	packageMsgs    map[string]map[string]*MsgDefine // pkg name -> type name -> MsgDefine
	packageDepends map[string]map[string]bool       // pkg name -> pkg name
	msgDepends     map[string]map[string]bool       // pkg.msg name -> []pkg.msg name
	messageApi     map[string]*MsgDefine            // topic -> MsgDefine
	serviceApi     map[string]*SrvDefine            // topic -> SrvDefine
}

func (i *impl) ParseMsg(topic, typename, schema string) error {
	pkg, name := i.nameParse(typename)
	msg, deps, err := i.parse(schema)
	if err != nil {
		return err
	}
	i.addMsg(&MsgDefine{
		Package: pkg,
		Name:    name,
		Block:   msg,
	})
	for name, define := range deps {
		p, n := i.nameParse(name)
		i.addMsg(&MsgDefine{
			Package: p,
			Name:    n,
			Block:   define,
		})
	}
	i.messageApi[topic] = &MsgDefine{
		Package: pkg,
		Name:    name,
		Block:   msg,
	}
	return nil
}

func (i *impl) ParseSrv(topic, typename, reqS, rspS string) error {
	pkg, name := i.nameParse(typename)
	req, deps, err := i.parse(reqS)
	if err != nil {
		return err
	}
	i.addMsg(&MsgDefine{
		Package: pkg,
		Name:    name + "Req",
		Block:   req,
	})
	for name, define := range deps {
		p, n := i.nameParse(name)
		i.addMsg(&MsgDefine{
			Package: p,
			Name:    n,
			Block:   define,
		})
	}
	rsp, deps, err := i.parse(rspS)
	if err != nil {
		return err
	}
	i.addMsg(&MsgDefine{
		Package: pkg,
		Name:    name + "Rsp",
		Block:   rsp,
	})
	for name, define := range deps {
		p, n := i.nameParse(name)
		i.addMsg(&MsgDefine{
			Package: p,
			Name:    n,
			Block:   define,
		})
	}
	i.serviceApi[topic] = &SrvDefine{
		Package: pkg,
		Name:    name,
		Req:     req,
		Rsp:     rsp,
	}
	return nil
}

func (i *impl) nameParse(name string) (string, string) {
	parts := strings.Split(name, "/")
	if len(parts) == 1 {
		return "", parts[0]
	}
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], parts[2]
}

func (i *impl) addMsg(m *MsgDefine) {
	var f func(msg *MsgDefine)
	f = func(msg *MsgDefine) {
		if _, ok := i.packageMsgs[msg.Package]; !ok {
			i.packageMsgs[msg.Package] = make(map[string]*MsgDefine)
		}
		i.packageMsgs[msg.Package][msg.Name] = msg
		for _, field := range msg.Block.Fields {
			if field.Type.IsCustom {
				p, name := i.nameParse(field.Type.Name)
				if p == "" {
					p = msg.Package
				}
				if i.msgDepends[msg.Package+"."+msg.Name] == nil {
					i.msgDepends[msg.Package+"."+msg.Name] = make(map[string]bool)
				}
				i.msgDepends[msg.Package+"."+msg.Name][p+"."+name] = true

				if msg.Package == p {
					continue
				}
				if i.packageDepends[msg.Package] == nil {
					i.packageDepends[msg.Package] = make(map[string]bool)
				}
				i.packageDepends[msg.Package][p] = true
			}
		}
	}
	f(m)
}

func (i *impl) parse(s string) (*walk.StructDefine, map[string]*walk.StructDefine, error) {
	input := antlr.NewInputStream(s)
	lexer := parser.NewFoxgloveSchemaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewFoxgloveSchemaParser(stream)
	el := newErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
	}))
	res := p.Parse()
	if el.GetErrs() != nil {
		return nil, nil, errors.Join(el.GetErrs()...)
	}
	main, sub := walk.Walk(res)
	return main, sub, nil
}

func (i *impl) Output() map[string]string {
	res := make(map[string]string, len(i.packageMsgs))
	for pName, msgs := range i.packageMsgs {
		res[pName] += "syntax = \"proto3\";\n"
		res[pName] += fmt.Sprintf("package %s;\n", pName)
		res[pName] += "option go_package = \"./pb\";\n"
		for k := range i.packageDepends[pName] {
			res[pName] += fmt.Sprintf("import \"%s.proto\";\n", k)
		}
		res[pName] += "\n"
		for name, m := range msgs {
			res[pName] += i.msgToPbForMulti(pName, name, m)
		}
	}

	mainContent := "syntax = \"proto3\";\noption go_package = \"./pb\";\n"
	for pName := range i.packageMsgs {
		mainContent += fmt.Sprintf("import \"%s.proto\";\n", pName)
	}
	mainContent += "\n"

	mainContent += "service Api {\n"
	for topic, msg := range i.messageApi {
		mainContent += fmt.Sprintf("\t// topic: %s\n", topic)
		mainContent += fmt.Sprintf("\trpc PublishTopic%s(%s) returns (%s);\n", Topic2Name(topic), msg.Package+"."+msg.Name, msg.Package+"."+msg.Name)
		mainContent += fmt.Sprintf("\trpc SubscribeTopic%s(%s) returns (stream %s);\n", Topic2Name(topic), msg.Package+"."+msg.Name, msg.Package+"."+msg.Name)
	}

	for topic, srv := range i.serviceApi {
		mainContent += fmt.Sprintf("\t// service: %s\n", topic)
		mainContent += fmt.Sprintf("\trpc CallService%s(%sReq) returns (%sRsp);\n", Topic2Name(topic), srv.Package+"."+srv.Name, srv.Package+"."+srv.Name)
	}
	mainContent += "}\n"

	res["main"] = mainContent

	return res
}

func (i *impl) msgToPbForMulti(pName, name string, m *MsgDefine) string {
	res := ""
	res += fmt.Sprintf("message %s {\n", name)
	for index, field := range m.Block.Fields {
		t := field.Type.Name
		if t == "byte" && field.Type.IsArray {
			t = "bytes"
		} else {
			if pbType, ok := pbTypeMap[field.Type.Name]; ok {
				t = pbType
			}
			if p, _ := i.nameParse(field.Type.Name); field.Type.IsCustom && p != pName {
				t = strings.Replace(t, "/", ".", -1)
			}
			if field.Type.IsArray {
				t = "repeated " + t
			}
		}
		fieldStr := fmt.Sprintf("\t%s %s = %d;\n", t, field.Name, index+1)
		res += fieldStr
	}
	res += fmt.Sprintf("}\n\n")
	return res
}

func (i *impl) OutputForType(t string) (string, string) {
	pName, name := i.nameParse(t)
	m := i.getMsg(pName, name)
	if m == nil {
		return "", ""
	}
	res := ""
	res += "syntax = \"proto3\";\n"
	res += fmt.Sprintf("package %s;\n", pName)
	res += "option go_package = \"./pb\";\n"
	res += "\n"
	res += i.msgToPbForSignal(pName, name, m)
	msgDup := map[string]bool{}
	var solve func(string, string)
	solve = func(pName, name string) {
		for subMsgName := range i.msgDepends[pName+"."+name] {
			parts := strings.SplitN(subMsgName, ".", 2)
			pName, name := parts[0], parts[1]
			if msgDup[pName+"."+name] {
				continue
			}
			msgDup[pName+"."+name] = true
			res += i.msgToPbForSignal(pName, name, i.getMsg(pName, name))
			solve(pName, name)
		}
	}
	solve(pName, name)
	return pName + "." + strings.Replace(t, "/", "_", -1), res
}

func (i *impl) msgToPbForSignal(pName, name string, m *MsgDefine) string {
	res := ""
	res += fmt.Sprintf("message %s {\n", pName+"_"+name)
	for index, field := range m.Block.Fields {
		t := field.Type.Name
		if t == "byte" && field.Type.IsArray {
			t = "bytes"
		} else {
			if pbType, ok := pbTypeMap[field.Type.Name]; ok {
				t = pbType
			}
			p, _ := i.nameParse(field.Type.Name)
			if p == "" {
				p = pName
				if field.Type.IsCustom {
					t = pName + "_" + t
				}
			}
			if field.Type.IsCustom && p != pName {
				t = strings.Replace(t, "/", "_", -1)
			}
			if field.Type.IsArray {
				t = "repeated " + t
			}
		}
		fieldStr := fmt.Sprintf("\t%s %s = %d;\n", t, field.Name, index+1)
		res += fieldStr
	}
	res += fmt.Sprintf("}\n\n")
	return res
}

func (i *impl) getMsg(pName, name string) *MsgDefine {
	p, ok := i.packageMsgs[pName]
	if !ok {
		return nil
	}
	m, ok := p[name]
	if !ok {
		return nil
	}
	return m
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
