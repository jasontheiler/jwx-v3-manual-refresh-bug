package main

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

const (
	// Any number of workers is okay. The more workers, the longer it will take
	// to get all of them stuck.
	numWorkers = 3
	// Delay between manual refresh attempts to allow the automatic refresh to
	// make attempts, too.
	delayBetweenManualRefreshes = 5 * time.Second
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.DebugContext(ctx, "received JWK set request")
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	errorSink := &jwkCacheErrorSink{
		logger: logger,
	}

	// Create a new JWK cache with the specified number of workers.
	cache, err := jwk.NewCache(ctx, httprc.NewClient(httprc.WithErrorSink(errorSink), httprc.WithWorkers(numWorkers)))
	if err != nil {
		panic(err)
	}

	// Don't wait for the automatic cache refresh to warm the cache and do it
	// manually instead further below.
	err = cache.Register(ctx, server.URL, jwk.WithWaitReady(false), jwk.WithConstantInterval(time.Second))
	if err != nil {
		panic(err)
	}

	// Do as many manual cache refresh attempts as there are workers.
	for range numWorkers {
		time.Sleep(delayBetweenManualRefreshes)

		if _, err := cache.Refresh(ctx, server.URL); err != nil {
			logger.ErrorContext(ctx, "failed to refresh JWK cache manually")
		}
	}

	// After the manual cache refresh attempts all cache workers are stuck and
	// no further cache refreshes (automatic or manual) can be done.

	_, err = cache.Refresh(ctx, server.URL)
	// This will never run, since all cache workers are stuck.
	panic(err)
}

type jwkCacheErrorSink struct {
	logger *slog.Logger
}

func (es *jwkCacheErrorSink) Put(ctx context.Context, err error) {
	es.logger.WarnContext(ctx, "failed to refresh JWK cache automatically", "error", err)
}
