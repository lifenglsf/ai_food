package v1

import "github.com/gogf/gf/v2/frame/g"

type SparkReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api1"`
}
type SparkRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
