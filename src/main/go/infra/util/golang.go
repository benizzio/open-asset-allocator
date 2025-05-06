package util

import (
	"reflect"
	"sync"
)

func ToPointerSlice[S any](slice []S) []*S {
	result := make([]*S, len(slice))
	for index, value := range slice {
		result[index] = &value
	}
	return result
}

type DeferRegistry struct {
	DeferFunctions []func()
	deferMutex     sync.Mutex
}

func (registry *DeferRegistry) RegisterDefer(f func()) {
	registry.deferMutex.Lock()
	defer registry.deferMutex.Unlock()
	registry.DeferFunctions = append(registry.DeferFunctions, f)
}

func (registry *DeferRegistry) Execute() {
	registry.deferMutex.Lock()
	defer registry.deferMutex.Unlock()
	for i := len(registry.DeferFunctions) - 1; i >= 0; i-- {
		registry.DeferFunctions[i]()
	}
	registry.DeferFunctions = make([]func(), 0)
}

func BuildDeferRegistry() *DeferRegistry {
	return &DeferRegistry{
		DeferFunctions: make([]func(), 0),
	}
}

func IsZeroValue[T any](value T) bool {
	return reflect.ValueOf(value).Equal(reflect.Zero(reflect.TypeOf(value)))
}
