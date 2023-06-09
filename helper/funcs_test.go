package helper

import (
	"reflect"
	"testing"

	"github.com/aybjax/nis_lib/pbdto"
)

func TestSetDiff(t *testing.T) {
	t.Run("It should return empty if first slice is empty", func(t *testing.T) {
		data := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		result := SetDiff(nil, data[:])

		if len(result) != 0 {
			t.Error("Empty first arg returns non-empty")
		}
	})
	t.Run("It should return first if second slice is empty", func(t *testing.T) {
		data := [...]int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		result := SetDiff(data[:], nil)

		if !reflect.DeepEqual(data[:], result) {
			t.Error("First slice is not returned")
		}
	})

	t.Run("It should return slice from subtraction of second from first", func(t *testing.T) {
		first := [...]int{1, 1, 2, 3, 4, 4, 6, 7, 8, 9}
		second := [...]int{1, 1, 2, 3}
		expected := map[int]int{
			4: 2,
			6: 1,
			7: 1,
			8: 1,
			9: 1,
		}
		resultMap := map[int]int{}

		result := SetDiff(first[:], second[:])

		for _, v := range result {
			resultMap[v] += 1
		}

		if !reflect.DeepEqual(expected, resultMap) {
			t.Error("First slice is not returned")
		}
	})
}

func TestGenerateUpdateMessage(t *testing.T) {
	t.Run("Should not return anything if modelId is nil value", func(t *testing.T) {
		toUpdate := [...]string{"1", "2", "3"}

		for range GenerateUpdateMessage("", toUpdate[:], 0) {
			t.Error("Returned a value")
		}
	})
	t.Run("Should not return anything if updateId is nil value", func(t *testing.T) {
		toUpdate := [...]string{"", "", "3"}
		const modelId = "ModelId"
		expected := [...]*pbdto.UpdateEmbedded{{
			Id:        "3",
			PayloadId: modelId,
		}}
		var results []*pbdto.UpdateEmbedded

		for v := range GenerateUpdateMessage(modelId, toUpdate[:], 0) {
			results = append(results, v)
		}

		if !reflect.DeepEqual(expected[:], results) {
			t.Error("Ids are not filtered well")
		}
	})
	t.Run("Update type should be given", func(t *testing.T) {
		toUpdate := [...]string{"3"}
		const modelId = "ModelId"
		expected := 1
		var result int

		for v := range GenerateUpdateMessage(modelId, toUpdate[:], 1) {
			result = int(v.Type)
		}

		if result != expected {
			t.Error("Update type is not preserved")
		}
	})
}
