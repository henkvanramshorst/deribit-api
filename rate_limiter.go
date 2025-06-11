package deribit

import (
	"context"
	"golang.org/x/time/rate"
)

type rateLimiter struct {
	ctx    context.Context
	global *rate.Limiter
	custom *rate.Limiter
}

func newRateLimiter(ctx context.Context) *rateLimiter {
	return &rateLimiter{
		ctx:    ctx,
		global: rate.NewLimiter(rate.Limit(5), 20),
		custom: rate.NewLimiter(rate.Limit(0.1), 5),
	}
}

func (r *rateLimiter) Wait(method string) error {
	if method == "public/get_instruments" {
		return r.custom.Wait(r.ctx)
	}
	return r.global.Wait(r.ctx)
}
