package cmntypes

//go:generate mockgen -source=./queue.go -destination=./mock_cmntypes/mock_queue.go
type AppQueue interface {
	Publish(data []byte, queue string, topic string) error
	Subscribe(queue string, topic string, cb func(data []byte) error)
}
