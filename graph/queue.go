package graph

import "container/list"

// Queue a wrapper queue for std list.List
type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Enqueue(item interface{}) *list.Element {
	return q.list.PushBack(item)
}

func (q *Queue) Pop() *list.Element {
	back := q.list.Back()
	q.list.Remove(back)
	return back
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() <= 0
}
