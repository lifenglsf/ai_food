package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"spark/internal/controller/spark"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			gr := s.Group("/food")
			gr.Middleware(ghttp.MiddlewareHandlerResponse)
			gr.Group("/", func(group *ghttp.RouterGroup) {
				gr.POST("v1", spark.NewV1())
				gr.POST("v2", spark.NewV2())
				gr.POST("v3", spark.NewV3())
			})

			s.Run()
			return nil
		},
	}
)
