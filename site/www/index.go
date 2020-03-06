package www

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
	. "github.com/arkgo/gateway/base"
)

func init() {
	Www.Router("index", Map{
		"uri":  "/",
		"name": "首页", "text": "首页",
		"action": func(ctx *ark.Http) {

			vvv := ark.Cache("asdf")
			ark.Debug("cache", vvv)
			ark.Cache("asdf", "ahahahahaha")

			ctx.Text("hello ark")
		},
	})

	Www.Router("about", Map{
		"uri":  "/about",
		"name": "关于我们", "text": "关于我们",
		"action": func(ctx *ark.Http) {
			ctx.Text("about ark")
		},
	})
}
