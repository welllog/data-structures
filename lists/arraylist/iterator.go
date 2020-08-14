package arraylist

import "github.com/welllog/data-structures/containers"

func assertIteratorImplementation() {
    var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

type Iterator struct {
    list  *List
    index int
}

func (l *List) Iterator() Iterator {
    return Iterator{list: l, index: -1}
}

func (i *Iterator) Next() bool {
    if i.index < i.list.size {
        i.index++
    }
    return i.list.withinRange(i.index)
}

func (i *Iterator) Prev() bool {
    if i.index >= 0 {
        i.index--
    }
    return i.list.withinRange(i.index)
}

func (i *Iterator) Value() interface{} {
    return i.list.elements[i.index]
}

func (i *Iterator) Index() int {
    return i.index
}

func (i *Iterator) Begin() {
    i.index = -1
}

func (i *Iterator) End() {
    i.index = i.list.size
}

func (i *Iterator) First() bool {
    i.Begin()
    return i.Next()
}

func (i *Iterator) Last() bool {
    i.End()
    return i.Prev()
}
