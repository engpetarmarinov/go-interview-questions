package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	client := http.Client{Timeout: 10 * time.Second}
	urls := []string{
		"https://google.com",
		"https://slavi.bg",
		"https://amazon.com",
		"https://bbc.com",
		"https://facebook.com",
		"https://sadkjhsakjdhkjh.com",
	}

	type resp struct {
		Took       time.Duration
		StatusCode int
		Url        string
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	responses := make(chan resp, len(urls))
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			r, err := client.Get(url)
			if err != nil {
				logger.Error("Error fetching url", slog.String("url", url), slog.String("error", err.Error()))
				return
			}

			select {
			case responses <- resp{
				Took:       time.Since(start),
				StatusCode: r.StatusCode,
				Url:        url,
			}:
				return
			case <-ctx.Done():
				logger.Info("Request timed out", slog.String("url", url))
				return
			}
		}()
	}

	go func() {
		wg.Wait()
		close(responses)
	}()

	for r := range responses {
		logger.Info("Response", slog.String("url", r.Url), slog.String("took", r.Took.String()), slog.Int("status_code", r.StatusCode))
	}
}
