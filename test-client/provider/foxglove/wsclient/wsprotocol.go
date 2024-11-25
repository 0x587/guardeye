package wsclient

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/0x587/guardeye/test-client/provider/foxglove/v8gojnject/textdecoder"
	"github.com/0x587/guardeye/test-client/provider/foxglove/v8gojnject/textencoder"

	"github.com/samber/lo"
	v8 "rogchap.com/v8go"
)

//go:embed out.js
var js string

func newJsVm() (*v8.Context, error) {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)
	if err := textdecoder.InjectWith(iso, global); err != nil {
		return nil, err
	}
	if err := textencoder.InjectTo(iso, global); err != nil {
		return nil, err
	}
	ctx := v8.NewContext(iso, global)
	return ctx, nil
}

type foxgloveReader struct {
	ctx *v8.Context
}

func newFoxgloveReader(ctx *v8.Context, schema string) (*foxgloveReader, error) {
	if _, err := ctx.RunScript(js, ""); err != nil {
		panic(err)
	}
	if _, err := ctx.RunScript(fmt.Sprintf("const def_str =`%s`;", schema), ""); err != nil {
		panic(err)
	}
	if _, err := ctx.RunScript(`const def = parse(def_str)`, ""); err != nil {
		panic(err)
	}
	return &foxgloveReader{ctx: ctx}, nil
}

func (g *foxgloveReader) read(buf []byte) (string, error) {
	data := strings.Join(lo.Map(buf, func(i byte, _ int) string {
		return fmt.Sprintf("0x%02x", i)
	}), ",")
	if _, err := g.ctx.RunScript(fmt.Sprintf("dataSrc = [%s];", data), ""); err != nil {
		return "", err
	}
	res, err := g.ctx.RunScript("JSON.stringify(read(def, dataSrc))", "")
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
