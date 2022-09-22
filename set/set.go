package set

import "sync"

type Set[T comparable] struct {
	m    map[T]struct{}
	lock sync.RWMutex
}

// Add a new key
func (s *Set[T]) Add(key ...T) {
	if len(key) == 0 {
		return
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	for i := 0; i < len(key); i++ {
		s.m[key[i]] = struct{}{}
	}
}

// Del key
func (s *Set[T]) Del(key T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.m[key]; ok {
		delete(s.m, key)
	}
}

// Values returns unique keys
func (s *Set[T]) Values() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	results := make([]T, 0, len(s.m))
	for key, _ := range s.m {
		results = append(results, key)
	}
	return results
}

// Has returns has key
func (s *Set[T]) Has(key T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, has := s.m[key]
	return has
}

// Count returns element count
func (s *Set[T]) Count() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.m)
}

// NewSet init Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}
