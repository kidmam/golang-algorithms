package queue

type QueueItem struct {
	item interface{}
	prev *QueueItem
}

// Base data structure for Queue
type Queue struct {
	current *QueueItem
	last *QueueItem
	depth uint64
}

// Initializes new Queue and return it
func New() *Queue {
	var queue *Queue = new(Queue)

	queue.depth = 0

	return queue
}

// Puts a given item into Queue
func (queue *Queue) Enqueue(item interface{}) {
	if (queue.depth == 0) {
		queue.current = &QueueItem{item: item, prev: nil}
		queue.last = queue.current
		queue.depth++
		return
	}
		
	q := &QueueItem{item: item, prev: nil}
	queue.last.prev = q
	queue.last = q
	queue.depth++
}

// Extracts first item from the Queue
func (queue *Queue) Dequeue() interface{} {
	if (queue.depth > 0) {
		item := queue.current.item
		queue.current = queue.current.prev
		queue.depth--
		
		return item
	}

	return nil
}
