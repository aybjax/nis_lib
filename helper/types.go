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

func NewMapError(err error) *mapError {
	return &mapError{
		"error": err.Error(),
	}
}
