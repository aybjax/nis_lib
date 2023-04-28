package cmntypes

type AppQueue interface {
	Notify(data []byte, queue string, topic string) error
	NotifyListener(queue string, topic string, cb func(data []byte) error)
}
