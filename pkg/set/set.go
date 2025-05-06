// set представляет реализацию структуры данных множества
package set

import (
	"iter"
	"maps"
)

type nothing struct{}

type Set[T comparable] struct {
	items map[T]nothing
}

func New[T comparable](items ...T) *Set[T] {
	set := &Set[T]{items: make(map[T]nothing, len(items))}
	for _, item := range items {
		set.items[item] = nothing{}
	}

	return set
}

func FromSlice[T comparable](items []T) *Set[T] {
	set := &Set[T]{items: make(map[T]nothing, cap(items))}
	for _, item := range items {
		set.items[item] = nothing{}
	}

	return set
}

func (s *Set[T]) Items() iter.Seq[T] {
	return maps.Keys(s.items)
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) Has(item T) bool {
	_, has := s.items[item]
	return has
}

func (s *Set[T]) Add(item T) {
	s.items[item] = nothing{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Intersection(set *Set[T]) *Set[T] {
	intersection := New[T]()
	for item := range s.items {
		if _, has := set.items[item]; has {
			intersection.Add(item)
		}
	}

	return intersection
}

func (s *Set[T]) Difference(set *Set[T]) *Set[T] {
	diff := New[T]()
	for item := range s.items {
		if _, has := set.items[item]; !has {
			diff.Add(item)
		}
	}

	return diff
}

func (s *Set[T]) Union(set *Set[T]) *Set[T] {
	union := s.items
	for item := range set.items {
		union[item] = nothing{}
	}

	return &Set[T]{items: union}
}
