package spark

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"spark/internal/service"

	"spark/api/spark/v3"
)

func (c *ControllerV3) Spark(ctx context.Context, req *v3.SparkReq) (res *v3.SparkRes, err error) {
	g.RequestFromCtx(ctx)
	type request struct {
		Name   string `v:"name@required|max-length:20"`
		OpenId string `v:"openid@required"`
	}
	name := req.Name
	openid := req.OpenId
	rr := request{Name: name, OpenId: openid}
	err = g.Validator().Data(rr).Run(gctx.New())
	res = &v3.SparkRes{}
	if err != nil {
		res.Code = 100
		res.Message = err.Error()
		res.Value = ""
		return
		//r.Response.WriteJsonExit(map[string]interface{}{"code": 100, "msg": err.Error(), "value": ""})
	}
	er, val, qt, pt, ct, tt := service.Gen(name, "v3")
	if er != nil {
		res.Code = 100
		res.Message = er.Error()
		res.Value = ""
		return
		//r.Response.WriteJsonExit(map[string]interface{}{"code": 100, "msg": er.Error(), "value": ""})
	}
	service.SaveSpark(name, val, openid, qt, pt, ct, tt)
	res.Code = 100
	res.Message = "success"
	res.Value = val
	return res, nil
}
