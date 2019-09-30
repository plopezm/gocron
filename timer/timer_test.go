package timer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx       context.Context
		startTime time.Duration
		delay     time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "From now start 2 seconds delay",
			args: args{
				startTime: 0,
				delay:     2 * time.Second,
			},
		},
		{
			name: "2 seconds start 2 seconds delay",
			args: args{
				startTime: 2 * time.Second,
				delay:     2 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctxWithCancel, cancel := context.WithCancel(context.Background())
			got := New(ctxWithCancel, time.Now().Add(tt.args.startTime), tt.args.delay)
			now := time.Now().Add(tt.args.startTime)
			i := 0
			for tick := range got {
				assert.Equal(t, now.Format(time.RFC1123), tick.Format(time.RFC1123))
				if i == 2 {
					cancel()
				}
				i++
				now = time.Now().Add(tt.args.delay)
			}
		})
	}
}
