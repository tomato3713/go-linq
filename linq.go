package linq

import (
	"iter"
	"slices"
)

type query[T any] iter.Seq[T]

func New[T any](iterator iter.Seq[T]) query[T] {
	return query[T](iterator)
}

func (f query[T]) Where(cond func(x T) bool) query[T] {
	return func(yield func(x T) bool) {
		for v := range f {
			if !cond(v)  {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func (f query[T]) OrderBy(cmp func(a, b T) int) query[T] {
	values := make([]T, 0)
	for v := range f {
		values = append(values, v)
	}

	slices.SortFunc(values, cmp)

	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}

func (f query[T]) Select(selector func(item T) T) query[T] {
	return func(yield func(T) bool) {
		for v := range f {
			if !(yield(selector(v))) {
				return
			}
		}
	}
}
