package types

import "errors"

type List[T comparable] interface {
	Add(T)
	Remove(T) error
	AsSlice() []T
}

var (
	ErrNoElementsToDelete      = errors.New("List does not contain any elements to delete")
	ErrElementToDeleteNotFound = errors.New("List does not contain such element to delete")
)
