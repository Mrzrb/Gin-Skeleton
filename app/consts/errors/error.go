package errors

import "app/infra"

var ServerError = &infra.BizError{
	Msg:  "服务器错误",
	Code: 1000,
}
