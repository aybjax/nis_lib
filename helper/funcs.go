package helper

import (
	"runtime"

	"github.com/aybjax/nis_lib/pbdto"
)

// Performs first - second as in sets
// repeated values stay intact
// output slice order is not guaranteed as it is filtered by map
func SetDiff[T comparable](first []T, second []T) (result []T) {
	filter := map[T]bool{}

	for ind, forRemoval := range second {
		if ind%100 == 0 {
			runtime.Gosched()
		}

		filter[forRemoval] = true
	}

	for ind, forAddition := range first {
		if ind%100 == 0 {
			runtime.Gosched()
		}

		if _, presentInSecond := filter[forAddition]; !presentInSecond {
			result = append(result, forAddition)
		}
	}

	return result
}

func GenerateUpdateMessage(modelId string, toUpdate []string, _type pbdto.UpdateType) <-chan *pbdto.UpdateEmbedded {
	ch := make(chan *pbdto.UpdateEmbedded)

	go func() {
		for _, tu := range toUpdate {
			if tu == "" || modelId == "" {
				continue
			}

			msg := &pbdto.UpdateEmbedded{
				Id:        tu,
				PayloadId: modelId,
				Type:      _type,
			}

			ch <- msg
		}

		close(ch)
	}()

	return ch
}
