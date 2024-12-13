package profile

import "time"

func Measure[T any](f func() (T, error)) (T, error, time.Duration) {
	start := time.Now()
	result, err := f()
	elapsed := time.Since(start)
	return result, err, elapsed
}
