package spark

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
	"spark/internal/service"
	"strings"
	"time"

	"spark/api/spark/vswxg"
)

func (c *ControllerVswxg) Spark(ctx context.Context, req *vswxg.SparkReq) (res *vswxg.SparkRes, err error) {
	g.RequestFromCtx(ctx)
	type request struct {
		Name   string `v:"name@required|max-length:20"`
		OpenId string `v:"openid@required"`
	}
	name := req.Name
	openid := req.OpenId
	rr := request{Name: name, OpenId: openid}
	err = g.Validator().Data(rr).Run(gctx.New())
	res = &vswxg.SparkRes{}
	now := time.Now()
	if err != nil {
		res.Code = 100
		res.Message = err.Error()
		res.Value = ""
		return
		//r.Response.WriteJsonExit(map[string]interface{}{"code": 100, "msg": err.Error(), "value": ""})
	}
	er, val, qt, pt, ct, tt := service.Gen(name, "vswxg")
	if er != nil {
		res.Code = 100
		res.Message = er.Error()
		res.Value = ""
		return
		//r.Response.WriteJsonExit(map[string]interface{}{"code": 100, "msg": er.Error(), "value": ""})
	}
	val = strings.ReplaceAll(val, "```json\n", "")
	val = strings.ReplaceAll(val, "```", "")
	val = strings.ReplaceAll(val, "\n", "")
	log.Printf("%#v", val)
	service.SaveSpark(name, val, openid, qt, pt, ct, tt)
	res.Code = 100
	res.Message = "success"
	res.Value = val
	var rrs interface{}

	err = json.Unmarshal([]byte(val), &rrs)
	log.Println(time.Since(now))
	return res, nil
}
