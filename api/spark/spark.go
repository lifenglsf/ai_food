// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package spark

import (
	"context"

	"spark/api/spark/v1"
	"spark/api/spark/v2"
	"spark/api/spark/v3"
	"spark/api/spark/v35"
	"spark/api/spark/vswxg"
)

type ISparkV1 interface {
	Spark(ctx context.Context, req *v1.SparkReq) (res *v1.SparkRes, err error)
}

type ISparkV2 interface {
	Spark(ctx context.Context, req *v2.SparkReq) (res *v2.SparkRes, err error)
}

type ISparkV3 interface {
	Spark(ctx context.Context, req *v3.SparkReq) (res *v3.SparkRes, err error)
}

type ISparkV35 interface {
	Spark(ctx context.Context, req *v35.SparkReq) (res *v35.SparkRes, err error)
}

type ISparkVswxg interface {
	Spark(ctx context.Context, req *vswxg.SparkReq) (res *vswxg.SparkRes, err error)
}
