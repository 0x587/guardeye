package async

import (
	"context"

	"github.com/samber/lo/parallel"
)

func GoAndWait(ctx context.Context, fs ...func() error) error {
	ch := make(chan error)
	go func() {
		errs := parallel.Map(fs, func(f func() error, _ int) error {
			return f()
		})
		for _, err := range errs {
			if err != nil {
				ch <- err
			}
		}
		ch <- nil
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
	}
}
