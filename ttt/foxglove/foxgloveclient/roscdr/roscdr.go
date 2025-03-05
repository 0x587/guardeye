package roscdr

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	v8 "rogchap.com/v8go"

	"github.com/0x587/guardeye/common/pool"
	"github.com/0x587/guardeye/ttt/foxglove/foxgloveclient/roscdr/v8gojnject/textdecoder"
	"github.com/0x587/guardeye/ttt/foxglove/foxgloveclient/roscdr/v8gojnject/textencoder"
)

//go:embed out.js
var js string

type readerInfo struct {
	once                sync.Once
	pool                pool.IF[*v8.Context]
	foxgloveScriptCache *v8.CompilerCachedData
}

// map[string]*readerInfo
var readers sync.Map

func makeCache(source string) *v8.CompilerCachedData {
	iso := v8.NewIsolate()
	script := lo.Must(iso.CompileUnboundScript(source, "foxglove.js", v8.CompileOptions{}))
	return script.CreateCodeCache()
}

func jsWithSchema(schema string) string {
	return js + fmt.Sprintf("\nconst def = parse(`%s`);", schema)
}

func makeFoxgloveSchemaCache(schema string) {
	r, exist := readers.Load(schema)
	if !exist {
		logx.Must(errors.New("foxglove: Failed to load reader"))
	}
	reader := r.(*readerInfo)
	reader.foxgloveScriptCache = makeCache(jsWithSchema(schema))
}

func ctxMaker(schema string) *v8.Context {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)
	lo.Must0(textdecoder.InjectWith(iso, global))
	lo.Must0(textencoder.InjectTo(iso, global))
	ctx := v8.NewContext(iso, global)
	r, exist := readers.Load(schema)
	if !exist {
		logx.Must(errors.New("foxglove: Failed to load reader"))
	}
	reader := r.(*readerInfo)
	s := lo.Must(ctx.Isolate().CompileUnboundScript(jsWithSchema(schema), "foxglove.js", v8.CompileOptions{CachedData: reader.foxgloveScriptCache}))
	lo.Must(s.Run(ctx))
	return ctx
}

func ctxCleaner(ctx *v8.Context) {
	ctx.Close()
	ctx.Isolate().Dispose()
}

func Parse(schema string, buf []byte) (string, error) {
	r, exist := readers.Load(schema)
	var reader *readerInfo
	if exist {
		reader = r.(*readerInfo)
	} else {
		reader = &readerInfo{
			once:                sync.Once{},
			pool:                nil,
			foxgloveScriptCache: nil,
		}
		readers.Store(schema, reader)
	}
	reader.once.Do(func() {
		reader.pool = pool.New(10, 10, func() *v8.Context { return ctxMaker(schema) }, ctxCleaner)
		makeFoxgloveSchemaCache(schema)
	})
	var res string
	reader.pool.Use(func(ctx *v8.Context) *pool.UseOption {
		data := strings.Join(lo.Map(buf, func(i byte, _ int) string {
			return fmt.Sprintf("0x%02x", i)
		}), ",")
		source := fmt.Sprintf("JSON.stringify(read(def, [%s]));", data)
		v, err := ctx.RunScript(source, "")
		if err != nil {
			logx.Errorf("foxglove: Failed to read data: %v", err)
			return nil
		}
		res = v.String()
		return nil
	})
	return res, nil
}
