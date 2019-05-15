package reduce

import "github.com/cheekybits/genny/generic"

// Item the type of item
type Item generic.Type

// GenericIteratee is a function invoked per iteration.
type GenericIteratee func(accumulator interface{}, item Item, index int, arr []Item) interface{}

// GenericReduce
func GenericReduce(collection []Item, iteratee GenericIteratee, initAccum interface{}) interface{} {
	index := 0
	length := len(collection)

	accumulator := initAccum
	for index < length {
		accumulator = iteratee(accumulator, collection[index], index, collection)
		index++
	}

	return accumulator
}
