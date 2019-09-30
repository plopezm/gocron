package gocron

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	type args struct {
		ctx       context.Context
		startTime time.Time
		delay     time.Duration
		onTick    OnTick
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Can assign and increment values in counter",
			args: args{
				ctx:       context.WithValue(context.Background(), "counter", 0),
				startTime: time.Now(),
				delay:     2 * time.Second,
				onTick: func(ctx context.Context, tick time.Time, cancel context.CancelFunc) (context.Context, error) {
					t := ctx.Value("t").(*testing.T)
					wg := ctx.Value("wg").(*sync.WaitGroup)
					counter := ctx.Value("counter").(int)
					if counter == 2 {
						assert.Equal(t, 2, counter)
						cancel()
						wg.Done()
						return ctx, nil
					}
					counter++
					return context.WithValue(ctx, "counter", counter), nil
				},
			},
		},
		{
			name: "Cancel never called, error ends the cron task",
			args: args{
				ctx:       context.WithValue(context.Background(), "counter", 0),
				startTime: time.Now(),
				delay:     2 * time.Second,
				onTick: func(ctx context.Context, tick time.Time, cancel context.CancelFunc) (context.Context, error) {
					t := ctx.Value("t").(*testing.T)
					wg := ctx.Value("wg").(*sync.WaitGroup)
					counter := ctx.Value("counter").(int)
					if counter == 1 {
						assert.Equal(t, 1, counter)
						defer wg.Done()
						return nil, errors.New("testing")
					}
					counter++
					return context.WithValue(ctx, "counter", counter), nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			_ = Do(context.WithValue(context.WithValue(tt.args.ctx, "t", t), "wg", wg), tt.args.startTime, tt.args.delay, tt.args.onTick)
			wg.Wait()
			time.Sleep(1)
		})
	}
}
