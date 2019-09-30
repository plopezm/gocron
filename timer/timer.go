package timer

import (
	"context"
	"time"
)

func New(ctx context.Context, startTime time.Time, delay time.Duration) <-chan time.Time {
	stream := make(chan time.Time, 1)

	go func() {
		t := <-time.After(time.Until(startTime))
		stream <- t
		ticker := time.NewTicker(delay)
		for {
			select {
			case tick := <-ticker.C:
				stream <- tick
			case <-ctx.Done():
				close(stream)
				ticker.Stop()
				return
			}
		}
	}()
	return stream
}
