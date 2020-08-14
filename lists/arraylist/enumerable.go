package arraylist

import "github.com/welllog/data-structures/containers"

func assertEnumerableImplementation() {
    var _ containers.EnumerableWithIndex = (*List)(nil)
}

func (l *List) Each(f func(index int, value interface{})) {
    iterator := l.Iterator()
    for iterator.Next() {
        f(iterator.Index(), iterator.Value())
    }
}

func (l *List) Map(f func(index int, value interface{}) interface{}) *List {
    newList := &List{}
    iterator := l.Iterator()
    for iterator.Next() {
        newList.Add(f(iterator.Index(), iterator.Value()))
    }
    return newList
}

func (l *List) Select(f func(index int, value interface{}) bool) *List {
    newList := &List{}
    iterator := l.Iterator()
    for iterator.Next() {
        if f(iterator.Index(), iterator.Value()) {
            newList.Add(iterator.Value())
        }
    }
    return newList
}

func (l *List) Any(f func(index int, value interface{}) bool) bool {
    iterator := l.Iterator()
    for iterator.Next() {
        if f(iterator.Index(), iterator.Value()) {
            return true
        }
    }
    return false
}

func (l *List) All(f func(index int, value interface{}) bool) bool {
    iterator := l.Iterator()
    for iterator.Next() {
        if !f(iterator.Index(), iterator.Value()) {
            return false
        }
    }
    return true
}

func (l *List) Find(f func(index int, value interface{}) bool) (int, interface{}) {
    iterator := l.Iterator()
    for iterator.Next() {
        if f(iterator.Index(), iterator.Value()) {
            return iterator.Index(), iterator.Value()
        }
    }
    return -1, nil
}