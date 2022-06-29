package cmd

import (
	"context"
	"os/signal"
	"syscall"
)

func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigs := make(chan os.signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signs:
			cancel()
		case <-ctx.Done():
			return
		}
	}()
	return ctx
}
