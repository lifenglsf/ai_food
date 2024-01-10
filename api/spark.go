package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"spark/internal/service"
)

func Food(r *ghttp.Request) {
	type req struct {
		Name   string `v:"name@required|max-length:20"`
		OpenId string `v:"openid@required"`
	}
	name := r.PostFormValue("name")
	openid := r.PostFormValue("openid")
	rr := req{Name: name, OpenId: openid}
	err := g.Validator().Data(rr).Run(gctx.New())
	if err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 100, "msg": err.Error(), "value": ""})
	}
	er, val, qt, pt, ct, tt := service.Gen(name)
	if er != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 100, "msg": er.Error(), "value": ""})
	}
	service.SaveSpark(name, val, openid, qt, pt, ct, tt)
	r.Response.WriteJsonExit(map[string]interface{}{"code": 0, "msg": "success", "value": val})
}
