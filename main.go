package main

import (
	_ "mgpanel/internal/logic"

	"mgpanel/internal/cmd"
	_ "mgpanel/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
