package langext

// TODO add an interface to this
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
