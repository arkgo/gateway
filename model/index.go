package model

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {
	ark.Register("test", Map{
		"name": "asdf",
	})

	// ark.Register("int", ark.Type{
	// 	Name: "整型", Desc: "整型",
	// 	Valid: func() bool {
	// 		return false
	// 	},
	// 	Value: func(val Any) Any {
	// 		return val
	// 	},
	// })

}
