package singlylinkedlist

import (
    "fmt"
    "github.com/welllog/data-structures/lists"
    "github.com/welllog/data-structures/utils"
    "strings"
)

func assertListImplementation() {
    var _ lists.List = (*List)(nil)
}

type List struct {
    first *element
    last *element
    size int
}

type element struct {
    value interface{}
    next *element
}

func New(values ...interface{}) *List {
    list := &List{}
    if len(values) > 0 {
        list.Add(values...)
    }
    return list
}

func (l *List) Add(values ...interface{}) {
    for _, value := range values {
        newElm := &element{value: value}
        if l.size == 0 {
            l.first = newElm
            l.last = newElm
        } else {
            l.last.next = newElm
            l.last = newElm
        }
        l.size++
    }
}

func (l *List) Append(values ...interface{}) {
    l.Add(values...)
}

func (l *List) Prepend(values ...interface{}) {
    for v := len(values) - 1; v >= 0; v-- {
        newElm := &element{value: values[v], next: l.first}
        l.first = newElm
        if l.size == 0 {
            l.last = newElm
        }
        l.size++
    }
}

func (l *List) Get(index int) (interface{}, bool) {
    if !l.withinRange(index) {
        return nil, false
    }
    elm := l.first
    for e := 0; e < index; e++ {
        elm = elm.next
    }
    return elm.value, true
}

func (l *List) Remove(index int) {
    if !l.withinRange(index) {
        return
    }
    if l.size == 0 {
        l.Clear()
        return
    }
    var beforeElm *element
    elm := l.first
    for e := 0; e < index; e++ {
        beforeElm = elm
        elm = elm.next
    }
    if elm == l.first {
        l.first = elm.next
    } else if elm == l.last {
        l.last = beforeElm
    }
    if beforeElm != nil {
        beforeElm.next = elm.next
    }
    elm.next = nil
    l.size--
}

func (l *List) Contains(values ...interface{}) bool {
    if len(values) == 0 {
        return true
    }
    if l.size == 0 {
        return false
    }
    for _, value := range values {
        found := false
        for elm := l.first; elm != nil; elm = elm.next {
            if elm.value == value {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }
    return true
}

func (l *List) Values() []interface{} {
    values := make([]interface{}, l.size, l.size)
    for e, elm := 0, l.first; elm != nil; e, elm = e+1, elm.next {
        values[e] = elm.value
    }
    return values
}

func (l *List) IndexOf(value interface{}) int {
    if l.size == 0 {
        return -1
    }
    for index, val := range l.Values() {
        if val == value {
            return index
        }
    }
    return -1
}

func (l *List) Empty() bool {
    return l.size == 0
}

func (l *List) Size() int {
    return l.size
}

func (l *List) Clear() {
    l.size = 0
    l.first = nil
    l.last = nil
}

func (l *List) Sort(comparator utils.Comparator) {
    
    if l.size < 2 {
        return
    }
    
    values := l.Values()
    utils.Sort(values, comparator)
    
    l.Clear()
    
    l.Add(values...)
    
}

func (l *List) Swap(i, j int) {
    if l.withinRange(i) && l.withinRange(j) && i != j {
        var element1, element2 *element
        for e, currentElement := 0, l.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
            switch e {
            case i:
                element1 = currentElement
            case j:
                element2 = currentElement
            }
        }
        element1.value, element2.value = element2.value, element1.value
    }
}

func (l *List) Insert(index int, values ...interface{}) {
    if !l.withinRange(index) {
        // Append
        if index == l.size {
            l.Add(values...)
        }
        return
    }
    
    l.size += len(values)
    
    var beforeElement *element
    foundElement := l.first
    for e := 0; e != index; e, foundElement = e+1, foundElement.next {
        beforeElement = foundElement
    }
    
    if foundElement == l.first {
        oldNextElement := l.first
        for i, value := range values {
            newElement := &element{value: value}
            if i == 0 {
                l.first = newElement
            } else {
                beforeElement.next = newElement
            }
            beforeElement = newElement
        }
        beforeElement.next = oldNextElement
    } else {
        oldNextElement := beforeElement.next
        for _, value := range values {
            newElement := &element{value: value}
            beforeElement.next = newElement
            beforeElement = newElement
        }
        beforeElement.next = oldNextElement
    }
}

func (l *List) Set(index int, value interface{}) {
    
    if !l.withinRange(index) {
        // Append
        if index == l.size {
            l.Add(value)
        }
        return
    }
    
    foundElement := l.first
    for e := 0; e != index; {
        e, foundElement = e+1, foundElement.next
    }
    foundElement.value = value
}

func (l *List) String() string {
    str := "SinglyLinkedList\n"
    values := []string{}
    for element := l.first; element != nil; element = element.next {
        values = append(values, fmt.Sprintf("%v", element.value))
    }
    str += strings.Join(values, ", ")
    return str
}

func (l *List) withinRange(index int) bool {
    return index >= 0 && index < l.size
}
