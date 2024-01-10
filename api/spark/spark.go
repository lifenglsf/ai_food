// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package spark

import (
	"context"

	"spark/api/spark/v1"
)

type ISparkV1 interface {
	Spark(ctx context.Context, req *v1.SparkReq) (res *v1.SparkRes, err error)
}
