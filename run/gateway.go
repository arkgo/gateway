package main

import (
	"github.com/arkgo/ark"
	_ "github.com/arkgo/driver"

	_ "github.com/arkgo/gateway/site"
	_ "github.com/arkgo/gateway/site/www"
)

func main() {
	ark.Go()
}
