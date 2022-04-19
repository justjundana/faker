package faker

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	mu                   = &sync.Mutex{}
	generateUniqueValues = false
	uniqueValues         = map[string][]interface{}{}
)

var (
	ErrUniqueFailure = "Failed to generate a unique value for field \"%s\""
	ErrNoUniqueTag   = "Failed to generate fake data without unique values for tag \"%s\""
)

const (
	maxRetry = 10000
)

type TaggedFunction func(reflect.Value) interface{}

func isUnique(tag string, value interface{}) bool {
	if _, ok := uniqueValues[tag]; !ok {
		uniqueValues[tag] = []interface{}{value}
		return true
	}

	for _, v := range uniqueValues[tag] {
		if v == value {
			return false
		}
	}

	uniqueValues[tag] = append(uniqueValues[tag], value)
	return true
}

func uniqueTag(tag string, fn TaggedFunction) (interface{}, error) {
	if _, ok := uniqueValues[tag]; !ok {
		return fn(reflect.Value{}), nil
	}

	for i := 0; i < maxRetry; i++ {
		res := fn(reflect.Value{})
		if isUnique(tag, res) {
			return res, nil
		}
	}

	return nil, fmt.Errorf(ErrUniqueFailure, tag)
}

func fakeData(tag string, fn TaggedFunction) interface{} {
	if generateUniqueValues {
		res, err := uniqueTag(tag, fn)
		if err != nil {
			panic(fmt.Errorf(ErrNoUniqueTag, tag))
		}

		return res
	}

	return fn(reflect.Value{})
}
