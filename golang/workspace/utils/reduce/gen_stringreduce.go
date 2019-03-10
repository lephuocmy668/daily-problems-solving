// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package reduce

// Iteratee is a function invoked per iteration.
type StringIteratee func(accumulator interface{}, item string, index int, arr []string) interface{}

// Reduce
func StringReduce(collection []string, iteratee StringIteratee, initAccum interface{}) interface{} {
	index := 0
	length := len(collection)

	accumulator := initAccum
	for index < length {
		accumulator = iteratee(accumulator, collection[index], index, collection)
		index++
	}

	return accumulator
}