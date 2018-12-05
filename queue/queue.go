package queue

type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() int {
	tail := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return tail
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
