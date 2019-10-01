# gocron

gocron is a time-based job scheduler for golang golang applicacions. gocron can be used to schedule jobs to run periodically at fixed times, dates, or intervals. It typically automates system maintenance or administration—though its general-purpose nature makes it useful for things like downloading files from the Internet and downloading email at regular intervals. gocron supports contexts.

# How to use it

gocron allows us to create tasks to be repeated. By using the context, variables can be passed to use them inside the runner and modify inside if needed. The task can be canceled inside by using funcion cancel() or outside.

```
cancel = gocron.Do(context.Background(), time.Now(), 2*time.Second, func(ctx context.Context, tick time.Time, cancel context.CancelFunc) (context.Context, error) {					
		counter := ctx.Value("counter").(int)
		if counter == 2 {
			cancel()
			return ctx, nil
		}
		counter++
		return context.WithValue(ctx, "counter", counter), nil
	})
```

The function Do(...) will return a cancel function to cancel the timer when needed.
