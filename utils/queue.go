package utils

type Queue[T any] struct {
	items  []T
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items:  make([]T, 0),
		length: 0,
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
	q.length++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.length > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		q.length--
		return item, true
	}
	var zero T
	return zero, false
}

func (q *Queue[T]) Items() []T {
	return q.items
}
