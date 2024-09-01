package wire

import "sync"

type Singleton[T any] struct {
	instance *T
	once     sync.Once
	initFunc func() *T
}

func (w *Singleton[T]) Get() *T {
	w.once.Do(func() {
		w.instance = w.initFunc()
	})
	return w.instance
}

func NewSingleton[T any](initFunc func() *T) *Singleton[T] {
	return &Singleton[T]{initFunc: initFunc}
}
