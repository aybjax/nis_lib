package cmntypes

type AppQueue interface {
	Notify([]byte) error
	NotifyListener(func([]byte) error)
}
