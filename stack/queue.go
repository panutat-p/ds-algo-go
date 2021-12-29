package stack

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
)

type Queue struct { // linked list based queue
	li *list.List
}

func (q Queue) Print() {
	var forward strings.Builder
	current := q.li.Front()
	for {
		if current == nil {
			forward.WriteString("nil")
			break
		}
		tmp := fmt.Sprintf("%v -> ", current.Value)
		forward.WriteString(tmp)
		current = current.Next()
	}
	fmt.Println(forward.String())
}

func (q Queue) Size() int {
	return q.li.Len()
}

func (q Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q Queue) Peek() (interface{}, error) {
	if q.Size() == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.li.Front().Value, nil
}

func (q *Queue) Dequeue() interface{} {
	ele := q.li.Front()
	q.li.Remove(ele)
	return ele.Value
}

func (q *Queue) Enqueue(v interface{}) {
	q.li.PushBack(v)
}
