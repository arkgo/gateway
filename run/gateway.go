package main

import (
	"github.com/arkgo/ark"
	_ "github.com/arkgo/driver"
)

func main() {
	ark.Ready()

	ark.Debug("什么鬼")

	ark.Go()
}
