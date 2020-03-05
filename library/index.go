package library

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

var (
	test = ark.Library("test")
)

func init() {

	test.Method("test", Map{
		"name": "测试", "text": "测试",
		"action": func(ctx *ark.Logic) (*Res) {
			ark.Debug("测试方法被调用")
			return ark.OK
		},
	})

	// test.Method("asdf", Map{
	// 	"name": "什么", "text": "什么",
	// 	"action": func() (*Res) {
	// 		return ark.OK
	// 	},
	// })

}