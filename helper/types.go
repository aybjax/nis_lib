package helper

import "encoding/json"

type mapError map[string]interface{}

func (m *mapError) Error() string {
	val, err := json.Marshal(m)

	if err != nil {
		return err.Error()
	}

	return string(val)
}

// Error as map for response object
// implements error interface
func NewMapError(err error) *mapError {
	return &mapError{
		"error": err.Error(),
	}
}
