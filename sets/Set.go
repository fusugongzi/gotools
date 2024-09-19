package sets

type set[T comparable, R any] struct {
	items   []R
	m       map[T]int
	keyFunc func(R) T
}

func NewSet[T comparable, R any](keyFunc func(R) T) *set[T, R] {
	s := &set[T, R]{
		items:   make([]R, 0),
		m:       make(map[T]int),
		keyFunc: keyFunc,
	}
	return s
}

// The value entered later will overwrite the previous value
func (s *set[T, R]) Add(value R) {
	key := s.keyFunc(value)
	if idx, ok := s.m[key]; ok {
		s.items[idx] = value
	} else {
		s.items = append(s.items, value)
		s.m[s.keyFunc(value)] = len(s.items) - 1
	}
}

func (s *set[T, R]) Values() []R {
	return s.items
}
