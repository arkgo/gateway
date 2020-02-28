package model

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {
	ark.Register("test", Map{
		"name": "asdf",
	})

}
