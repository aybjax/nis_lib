package cmntypes

//go:generate mockgen -source=./cache.go -destination=./mock_cmntypes/mock_cache.go
type AppCache interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	Delete(key string) error
}
