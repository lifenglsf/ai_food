package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"spark/internal/cmd"
	_ "spark/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
