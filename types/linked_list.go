package types

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	baseNode *Node[T]
}

func (list *LinkedList[T]) Add(data T) {
	// create the node which we'll add to this list
	newNode := &Node[T]{data, nil}

	// check if the list has any elements at all
	if list.baseNode == nil {
		// we don't have elements, add the first
		list.baseNode = newNode
	} else {
		// we do have elements already, so iterate over them to find the last one
		currentNode := list.baseNode

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		// we have the last node in the list, so append the new one to it
		currentNode.next = newNode
	}
}

func (list *LinkedList[T]) Remove(data T) error {
	// check that the list has any elements
	if list.baseNode == nil {
		return ErrNoElementsToDelete
	}

	// attempt to find the element to delete
	// and keep track of the previous node so that we can
	// attach it to the currentNode's next element
	currentNode := list.baseNode
	var previousNode *Node[T] = nil

	for currentNode != nil {
		if currentNode.data == data {
			// we found the first node which contains the data
			// now point the previous node to currentNode's child node
			// to prevent gaps, this effectively deletes currentNode from
			// the list. we immediately return nothing, as no errors occurred.
			previousNode.next = currentNode.next
			return nil
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}

	return ErrElementToDeleteNotFound
}

func (list LinkedList[T]) AsSlice() []T {
	var slice []T

	// only if the list contains elements will we append to the empty slice
	if list.baseNode != nil {
		// iterate through nodes and add their values to the slice
		currentNode := list.baseNode

		for {
			slice = append(slice, currentNode.data)

			// if the current node has a child then set it as the next current node
			// otherwise we break from the loop and return the resulting slice
			if currentNode.next != nil {
				currentNode = currentNode.next
			} else {
				break
			}
		}
	}

	return slice
}
