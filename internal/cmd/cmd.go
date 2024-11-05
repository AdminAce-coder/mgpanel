package cmd

import (
	"context"
	"mgpanel/internal/controller/aws"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					ghttp.MiddlewareCORS)
				group.Bind(
					aws.NewV1(),
				)
			})

			s.SetAddr("0.0.0.0:1213")
			g.Log().Info(ctx, "start http server success")
			s.Run()

			return nil
		},
	}
)
