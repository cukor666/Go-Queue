package queue

type Queue struct {
	front, rear int
	elems       []any
}

func (q *Queue) Init(cap int) {
	q.front = 0
	q.rear = 0
	q.elems = make([]any, cap)
}

func New(cap int) *Queue {
	q := new(Queue)
	q.Init(cap)
	return q
}

func (q *Queue) Size() int {
	if q.rear < q.front {
		return q.rear + len(q.elems) - q.front
	}
	return q.rear - q.front
}

func (q *Queue) Empty() bool {
	return q.front == q.rear
}

func (q *Queue) Front() any {
	if q.front == q.rear {
		return nil
	}
	return q.elems[q.front]
}

func (q *Queue) Rear() any {
	if q.front == q.rear {
		return nil
	}
	if q.rear == 0 {
		return q.elems[len(q.elems)-1]
	}
	return q.elems[q.rear-1]
}

func (q *Queue) Push(val any, increase bool) bool {
	if q.Size() < len(q.elems)-1 {
		q.elems[q.rear] = val
		q.rear = (q.rear + 1) % len(q.elems)
		return true
	} else if increase {
		q.elems[q.rear] = val
		m := len(q.elems)
		n := m + (m >> 1)
		newElems := make([]any, n)
		for i := 0; i < m; i++ {
			newElems[i] = q.elems[q.front]
			q.front = (q.front + 1) % m
		}
		q.elems = newElems
		q.front = 0
		q.rear = m
		return true
	}
	return false
}

func (q *Queue) Pop() any {
	if q.Empty() {
		return nil
	}
	delVal := q.elems[q.front]
	q.front = (q.front + 1) % len(q.elems)
	return delVal
}
