package singlylinkedlist

import (
    "github.com/welllog/data-structures/containers"
)

func assertIteratorImplementation() {
    var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

func (l *List) Iterator() *Iterator {
    return &Iterator{
        list:    l,
        index: -1,
    }
}

type Iterator struct {
    list *List
    index int
    element *element
}

func (i *Iterator) Next() bool {
    if i.index < i.list.size {
        i.index++
    }
    if !i.list.withinRange(i.index) {
        i.element = nil
        return false
    }
    if i.index == 0 {
        i.element = i.list.first
    } else {
        i.element = i.element.next
    }
    return true
}

func (i *Iterator) Value() interface{} {
    if i.element != nil {
        return i.element.value
    }
    return nil
}

func (i *Iterator) Index() int {
    return i.index
}

func (i *Iterator) Begin() {
    i.element = nil
    i.index = -1
}

func (i *Iterator) First() bool {
    i.Begin()
    return i.Next()
}
