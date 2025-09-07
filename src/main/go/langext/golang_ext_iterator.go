package langext

type Iterator[T any] interface {
	HasNext() bool
	NextPointer() (*T, int)
	Next() (T, int)
	CurrentPointer() (*T, int)
	Current() (T, int)
	Size() int
}

type SliceIterator[T any] struct {
	index int
	slice []T
}

func (iterator *SliceIterator[T]) HasNext() bool {
	return iterator.index < len(iterator.slice)-1
}

func (iterator *SliceIterator[T]) NextPointer() (*T, int) {
	iterator.index++
	var result = &iterator.slice[iterator.index]
	var resultIndex = iterator.index
	return result, resultIndex
}

func (iterator *SliceIterator[T]) Next() (T, int) {
	var pointer, index = iterator.NextPointer()
	return *pointer, index
}

func (iterator *SliceIterator[T]) CurrentPointer() (*T, int) {
	return &iterator.slice[iterator.index], iterator.index
}

func (iterator *SliceIterator[T]) Current() (T, int) {
	var pointer, index = iterator.CurrentPointer()
	return *pointer, index
}

func (iterator *SliceIterator[T]) Size() int {
	return len(iterator.slice)
}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{index: -1, slice: slice}
}
