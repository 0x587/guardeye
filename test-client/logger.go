package main

import (
	"github.com/kardianos/service"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogWriter struct {
	service.Logger
}

func (l *LogWriter) Alert(v any) {
	_ = l.Logger.Warning(v)
}

func (l *LogWriter) Close() error {
	return nil
}

func (l *LogWriter) Debug(v any, fields ...logx.LogField) {
	_ = l.Logger.Info(append([]any{v}, lo.ToAnySlice(fields)...)...)
}

func (l *LogWriter) Error(v any, fields ...logx.LogField) {
	_ = l.Logger.Error(append([]any{v}, lo.ToAnySlice(fields)...)...)
}

func (l *LogWriter) Info(v any, fields ...logx.LogField) {
	_ = l.Logger.Info(append([]any{v}, lo.ToAnySlice(fields)...)...)
}

func (l *LogWriter) Severe(v any) {
	_ = l.Logger.Error(v)
}

func (l *LogWriter) Slow(v any, fields ...logx.LogField) {
	_ = l.Logger.Info(append([]any{v}, lo.ToAnySlice(fields)...)...)
}

func (l *LogWriter) Stack(v any) {
	_ = l.Logger.Info(v)
}

func (l *LogWriter) Stat(v any, fields ...logx.LogField) {
	_ = l.Logger.Info(append([]any{v}, lo.ToAnySlice(fields)...)...)
}

var logger *LogWriter
