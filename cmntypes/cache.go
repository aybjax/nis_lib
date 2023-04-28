package cmntypes

type AppCache interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
}
