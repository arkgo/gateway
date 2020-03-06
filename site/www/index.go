package www

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {
	ark.Router("www.index", Map{
		"uri":  "/",
		"name": "首页", "text": "首页",
		"action": func(ctx *ark.Http) {
			ctx.Text("hello ark")
		},
	})

	ark.Router("www.about", Map{
		"uri":  "/about",
		"name": "关于我们", "text": "关于我们",
		"action": func(ctx *ark.Http) {
			ctx.Text("hello ark")
		},
	})
}
