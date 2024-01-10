package logic

import "context"

type ISpark interface {
	Gen(ctx context.Context, i string) (error, string, float64, float64, float64, float64)
}
