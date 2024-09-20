package stream

type stream[T any] struct {
	values []T
}

func NewStream[T any](values []T) *stream[T] {
	return &stream[T]{
		values: values,
	}
}

// Returns a stream consisting of the elements of this stream that match the given predicate
func (s *stream[T]) Filter(predicate func(T) bool) *stream[T] {
	result := make([]T, 0)
	for _, v := range s.values {
		v := v
		if predicate(v) {
			result = append(result, v)
		}
	}
	return &stream[T]{
		values: result,
	}
}

// only support same type map
func (s *stream[T]) Map(mapper func(T) T) *stream[T] {
	result := make([]T, 0)
	for _, v := range s.values {
		v := v
		result = append(result, mapper(v))
	}
	return &stream[T]{
		values: result,
	}
}

func (s *stream[T]) Collect() []T {
	return s.values
}
