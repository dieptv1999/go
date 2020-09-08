package main

import (
	"context"
)

type CalculateHandler struct {
}

func (c CalculateHandler) Add(ctx context.Context, a int32, b int32) (r int32, err error) {
	return a + b, nil
}

func NewCalculateHandler() *CalculateHandler {
	return &CalculateHandler{}
}
