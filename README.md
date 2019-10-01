# gocron

gocron is a time-based job scheduler for golang golang applicacions. gocron can be used to schedule jobs to run periodically at fixed times, dates, or intervals. It typically automates system maintenance or administration—though its general-purpose nature makes it useful for things like downloading files from the Internet and downloading email at regular intervals. gocron supports contexts.

# Examples

gocron allows us to create tasks to be repeated:

```
cancel = gocron.Do(context.Background(), time.Now(), 2*time.Second, tt.func(ctx context.Context, tick time.Time, cancel context.CancelFunc) (context.Context, error) {					
		counter := ctx.Value("counter").(int)
		if counter == 2 {
			cancel()
			return ctx, nil
		}
		counter++
		return context.WithValue(ctx, "counter", counter), nil
	})
```
