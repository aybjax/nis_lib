package cmntypes

type AppQueue interface {
	Notify(data []byte, queue string, topic string) error
	NotifyListener(func(data []byte, queue string, topic string) error)
}
