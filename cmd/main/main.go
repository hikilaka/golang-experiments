package main

import (
	"fmt"

	"github.com/hikilaka/golang-experiments/types"
)

func main() {
	list := types.LinkedList[int]{}

	for i := range 5 {
		list.Add(i)
	}

	fmt.Println(list.AsSlice())
}
