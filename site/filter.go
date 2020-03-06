package site

import (
	"time"

	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {
	ark.Filter("www.logger", Map{
		"name": "请求日志拦截器", "text": "请求日志拦截器",
		"request": func(ctx *ark.Http) {
			begin := time.Now()

			ctx.Next()

			logger := true
			if vv, ok := ctx.Config["logger"].(bool); ok {
				logger = vv
			}

			if logger {
				ark.Trace(ctx.Ip(), ctx.Id, ctx.Method, ctx.Site, ctx.Uri, ctx.Code, time.Now().Sub(begin))
			}
		},
	})
}
