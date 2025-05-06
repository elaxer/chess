// Description: Определение структуры данных "очередь".
package queue

import "container/list"

// Queue представляет собой структуру данных "очередь".
type Queue[T any] struct {
	list *list.List
}

// New создает новую очередь.
func New[T any]() *Queue[T] {
	return &Queue[T]{list.New()}
}

// Push добавляет элемент в конец очереди.
func (q *Queue[T]) Push(value T) {
	q.list.PushBack(value)
}

// Pop удаляет и возвращает элемент из начала очереди.
func (q *Queue[T]) Pop() T {
	element := q.list.Front()
	q.list.Remove(element)

	return element.Value.(T)
}

// Len возвращает количество элементов в очереди.
func (q *Queue[T]) Len() int {
	return q.list.Len()
}

// IsEmpty возвращает true, если очередь пуста.
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
