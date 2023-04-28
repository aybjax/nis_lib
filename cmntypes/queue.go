package cmntypes

type AppQueue interface {
	Enqueue(data []byte, queue string, topic string) error
	Dequeue(queue string, topic string, cb func(data []byte) error)
}
