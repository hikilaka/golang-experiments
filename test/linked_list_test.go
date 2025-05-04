package test

import (
	"testing"

	"github.com/hikilaka/golang-experiments/types"
	"github.com/stretchr/testify/assert"
)

func TestAddAndAsSlice(t *testing.T) {
	// arrange & act
	list := CreateBaseList()

	expected := []int{0, 1, 2, 3, 4}
	actual := list.AsSlice()

	// assert
	assert.Equal(t, expected, actual)
}

func TestRemove(t *testing.T) {
	// arrange
	list := CreateBaseList()

	// act
	result := list.Remove(2)

	expected := []int{0, 1, 3, 4}
	actual := list.AsSlice()

	// assert
	assert.Empty(t, result)
	assert.Equal(t, expected, actual)
}

func TestRemoveWithNoElements(t *testing.T) {
	// arrange
	list := types.LinkedList[int]{}

	// act
	result := list.Remove(0)

	// assert
	assert.EqualError(t, result, "List does not contain any elements to delete")
}

func TestRemoveWithNonexistantElement(t *testing.T) {
	// arrange
	list := CreateBaseList()

	// act
	result := list.Remove(9001)

	// assert
	assert.EqualError(t, result, "List does not contain such element to delete")
}

// creates a list with 5 elements, 0-4
func CreateBaseList() types.LinkedList[int] {
	list := types.LinkedList[int]{}

	for i := range 5 {
		list.Add(i)
	}

	return list
}
