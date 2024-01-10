package spark

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spark/api/spark/v1"
)

func (c *ControllerV1) Spark(ctx context.Context, req *v1.SparkReq) (res *v1.SparkRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
