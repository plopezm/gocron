package gocron

import (
	"context"
	"time"

	"github.com/plopezm/gocron/timer"
)

type OnTick func(context.Context, time.Time, context.CancelFunc) (context.Context, error)

func Do(ctx context.Context, startTime time.Time, delay time.Duration, onTick OnTick) context.CancelFunc {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	timerChannel := timer.New(ctxWithCancel, startTime, delay)
	go func(ctx context.Context, timerChannel <-chan time.Time, onTick OnTick, cancel context.CancelFunc) {
		var err error
		timerCtx := ctx
		for tick := range timerChannel {
			timerCtx, err = onTick(timerCtx, tick, cancel)
			if err != nil {
				cancel()
				return
			}
		}
	}(ctxWithCancel, timerChannel, onTick, cancel)
	return cancel
}
