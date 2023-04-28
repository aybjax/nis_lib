package cmntypes

type AppQueue interface {
	Publish(data []byte, queue string, topic string) error
	Subscribe(queue string, topic string, cb func(data []byte) error)
}
