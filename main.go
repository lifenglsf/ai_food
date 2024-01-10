package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"spark/api"
	_ "spark/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	router(s)
	s.Run()
	//cmd.Main.Run(gctx.GetInitCtx())
}
func router(s *ghttp.Server) {
	gr := s.Group("/spark")
	gr.POST("/food", api.Food)
}
