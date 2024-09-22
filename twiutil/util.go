package twiutil

import "fmt"

type Unit struct{}

type Result[T any] struct {
	data *T
	err  error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		data: &value,
		err:  nil,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		data: nil,
		err:  err,
	}
}

func Errorf[T any](message string, values ...any) Result[T] {
	return Result[T]{
		data: nil,
		err:  fmt.Errorf(message, values...),
	}
}

func ErrorOnMissing[T any](value T, ok bool, err error) Result[T] {
	if ok {
		return Result[T]{
			data: &value,
			err:  nil,
		}
	}
	return Result[T]{
		data: nil,
		err:  err,
	}
}

func ToResult[T any](value T, err error) Result[T] {
	return Result[T]{
		data: &value,
		err:  err,
	}
}

func (r Result[T]) IsOk() bool {
	return r.data != nil
}

func (r Result[T]) IsError() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() (T, error) {
	return *r.data, r.err
}

func (r Result[T]) Value() T {
	if r.err != nil {
		panic(r.err)
	}
	return *r.data
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) ValueOrDefault(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return *r.data
}

func OnOk[T any, U any](result Result[T], onOk func(T) Result[U]) Result[U] {
	if result.IsOk() {
		return onOk(*result.data)
	}
	return Error[U](result.err)
}

func OnError[T any](result Result[T], onError func(error) Result[T]) Result[T] {
	if result.IsError() {
		return onError(result.err)
	}
	return result
}

func (result Result[T]) OnError(onError func(error) Result[T]) Result[T] {
	if result.IsError() {
		return onError(result.err)
	}
	return result
}
