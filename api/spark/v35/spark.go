package v35

import "github.com/gogf/gf/v2/frame/g"

type SparkReq struct {
	g.Meta
	Name   string `json:"name"`
	OpenId string `json:"openid"`
}
type SparkRes struct {
	g.Meta  `mine:"application/json" example:"string"`
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Value   string `json:"value"`
}
