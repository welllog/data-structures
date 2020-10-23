package binaryheap

type Iterator struct {
    heap *Heap
    index int
}

func (h *Heap) Iterator() Iterator {
    return Iterator{heap: h, index: -1}
}

func (iterator *Iterator) Next() bool {
    if iterator.index < iterator.heap.Size() {
        iterator.index++
    }
    return iterator.heap.withinRange(iterator.index)
}

func (iterator *Iterator) Prev() bool {
    if iterator.index >= 0 {
        iterator.index--
    }
    return iterator.heap.withinRange(iterator.index)
}

func (iterator *Iterator) Value() interface{} {
    value, _ := iterator.heap.list.Get(iterator.index)
    return value
}

func (iterator *Iterator) Index() int {
    return iterator.index
}

func (iterator *Iterator) Begin() {
    iterator.index = -1
}

func (iterator *Iterator) End() {
    iterator.index = iterator.heap.Size()
}

func (iterator *Iterator) First() bool {
    iterator.Begin()
    return iterator.Next()
}

func (iterator *Iterator) Last() bool {
    iterator.End()
    return iterator.Prev()
}