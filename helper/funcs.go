package helper

import (
	"runtime"
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
