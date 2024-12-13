package async

import "sync"

func NewTask[T any](f func() (T, error)) *Task[T] {
	r := make(chan T)
	e := make(chan error)
	go func() {
		res, err := f()
		if err != nil {
			e <- err
		}
		r <- res
	}()
	return &Task[T]{
		res: r,
		err: e,
	}
}

type Task[T any] struct {
	sync.Mutex
	res chan T
	err chan error
}
