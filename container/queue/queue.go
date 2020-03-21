package queue

import "fmt"

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Empty queue, failed to pop ")
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
