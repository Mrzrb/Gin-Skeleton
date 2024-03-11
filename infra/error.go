package infra

import (
	"fmt"
	"strings"
)

type BizError struct {
	err  error
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
}

func (b *BizError) Error() string {
	var builder strings.Builder
	builder.WriteString(b.Msg)
	if b.err != nil {
		builder.WriteString("   | internal err:")
		builder.WriteString(b.err.Error())
	}
	return builder.String()
}

func (b *BizError) Wrap(err error) *BizError {
	b.err = err
	return b
}

func (b *BizError) Append(format string, objs ...any) {
	b.Msg = b.Msg + "   | " + fmt.Sprintf(format, objs...)
}
