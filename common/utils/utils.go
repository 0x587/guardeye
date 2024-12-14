package utils

import (
	"fmt"
	"strings"

	"github.com/0x587/guardeye/report/report"
)

func ProviderToStr(p *report.Provider) string {
	return fmt.Sprintf("%s(%s)", p.GetType(), strings.Join(p.GetArgs(), ","))
}

func NewErrOr[T any](val T, err error) ErrOr[T] {
	return ErrOr[T]{
		Err: err,
		Val: &val,
	}
}

type ErrOr[T any] struct {
	Err error
	Val *T
}

func (e *ErrOr[T]) Get() (*T, error) {
	if e.Err != nil {
		return nil, e.Err
	}
	return e.Val, nil
}
