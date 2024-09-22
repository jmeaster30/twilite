package twilite

import "fmt"

type twiUnit struct{}

type twiResult[T any] struct {
	data *T
	err  error
}

func Ok[T any](value T) twiResult[T] {
	return twiResult[T]{
		data: &value,
		err:  nil,
	}
}

func Error[T any](err error) twiResult[T] {
	return twiResult[T]{
		data: nil,
		err:  err,
	}
}

func Errorf[T any](message string, values ...any) twiResult[T] {
	return twiResult[T]{
		data: nil,
		err:  fmt.Errorf(message, values...),
	}
}

func ErrorOnMissing[T any](value T, ok bool, err error) twiResult[T] {
	if ok {
		return twiResult[T]{
			data: &value,
			err:  nil,
		}
	}
	return twiResult[T]{
		data: nil,
		err:  err,
	}
}

func ToResult[T any](value T, err error) twiResult[T] {
	return twiResult[T]{
		data: &value,
		err:  err,
	}
}

func (r twiResult[T]) IsOk() bool {
	return r.data != nil
}

func (r twiResult[T]) IsError() bool {
	return r.err != nil
}

func (r twiResult[T]) Unwrap() (T, error) {
	return *r.data, r.err
}

func (r twiResult[T]) Value() T {
	if r.err != nil {
		panic(r.err)
	}
	return *r.data
}

func (r twiResult[T]) Error() error {
	return r.err
}

func (r twiResult[T]) ValueOrDefault(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return *r.data
}

func OnOk[T any, U any](result twiResult[T], onOk func(T) twiResult[U]) twiResult[U] {
	if result.IsOk() {
		return onOk(*result.data)
	}
	return Error[U](result.err)
}

func OnError[T any](result twiResult[T], onError func(error) twiResult[T]) twiResult[T] {
	if result.IsError() {
		return onError(result.err)
	}
	return result
}

func (result twiResult[T]) OnError(onError func(error) twiResult[T]) twiResult[T] {
	if result.IsError() {
		return onError(result.err)
	}
	return result
}
