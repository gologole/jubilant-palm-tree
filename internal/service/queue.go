package service

import "sync"

// Queue представляет собой очередь с мьютексом.
type Queue struct {
	items []interface{}
	mu    sync.Mutex
}

// NewQueue создает новую очередь.
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Dequeue удаляет элемент из начала очереди и возвращает его.
func (q *Queue) Dequeue() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Get возвращает первый элемент из очереди, не удаляя его.
func (q *Queue) Get() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, false
	}

	return q.items[0], true
}

// Size возвращает количество элементов в очереди.
func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
