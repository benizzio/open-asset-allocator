package util

import "reflect"

func ToPointerSlice[S any](slice []S) []*S {
	result := make([]*S, len(slice))
	for index, value := range slice {
		result[index] = &value
	}
	return result
}

type Iterator[T any] struct {
	index int
	slice []T
}

func (iterator *Iterator[T]) HasNext() bool {
	return iterator.index < len(iterator.slice)-1
}

func (iterator *Iterator[T]) NextPointer() (*T, int) {
	iterator.index++
	var result = &iterator.slice[iterator.index]
	var resultIndex = iterator.index
	return result, resultIndex
}

func (iterator *Iterator[T]) Next() (T, int) {
	var pointer, index = iterator.NextPointer()
	return *pointer, index
}

func (iterator *Iterator[T]) CurrentPointer() (*T, int) {
	return &iterator.slice[iterator.index], iterator.index
}

func (iterator *Iterator[T]) Current() (T, int) {
	var pointer, index = iterator.CurrentPointer()
	return *pointer, index
}

func (iterator *Iterator[T]) Size() int {
	return len(iterator.slice)
}

func NewIterator[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{index: -1, slice: slice}
}

func DeepCompleteStruct[T interface{}](target *T, source *T) {

	if target == nil || source == nil {
		return
	}

	var structType = reflect.TypeOf(*target)
	var targetStructValue = reflect.ValueOf(target).Elem()
	var sourceStructValue = reflect.ValueOf(source).Elem()

	if structType.Kind() == reflect.Struct {
		deepCompleteReflective(structType, targetStructValue, sourceStructValue)
	}
}

func deepCompleteReflective(structType reflect.Type, targetStructValue reflect.Value, sourceStructValue reflect.Value) {

	if !targetStructValue.IsZero() {
		return
	}

	for i := 0; i < structType.NumField(); i++ {

		if structType.Field(i).Type.Kind() == reflect.Struct {
			deepCompleteReflective(
				structType.Field(i).Type,
				targetStructValue.Field(i).Elem(),
				sourceStructValue.Field(i).Elem(),
			)
		}

		var sourceStructFieldValue = sourceStructValue.Field(i)
		var targetStructFieldValue = targetStructValue.Field(i)

		if targetStructFieldValue.IsZero() {
			targetStructFieldValue.Set(sourceStructFieldValue)
		}
	}
}
