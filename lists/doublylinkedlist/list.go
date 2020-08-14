package doublylinkedlist

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
    root element
    size int
}

type element struct {
    value interface{}
    prev *element
    next *element
}

func New(values ...interface{}) *List {
    list := new(List).Init()
    if len(values) > 0 {
        list.Add(values...)
    }
    return list
}

func (l *List) Init() *List {
    l.root.next = &l.root
    l.root.prev = &l.root
    l.size = 0
    return l
}

func (l *List) Add(values ...interface{}) {
    for _, value := range values {
        l.insert(&element{value: value}, l.root.prev)
    }
}

func (l *List) Append(values ...interface{}) {
    l.Add(values...)
}

func (l *List) Prepend(values ...interface{}) {
    for v := len(values) - 1; v >= 0; v-- {
        l.insert(&element{value: values[v]}, &l.root)
    }
}

func (l *List) Get(index int) (interface{}, bool) {
    if !l.withinRange(index) {
        return nil, false
    }
    elm := l.root.next
    for e := 0; e < index; e++ {
        elm = elm.next
    }
    return elm.value, true
}

func (l *List) Remove(index int) {
    if !l.withinRange(index) {
        return
    }
    elm := l.root.next
    for e := 0; e < index; e++ {
        elm = elm.next
    }
    l.remove(elm)
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
        for elm := l.root.next; elm != nil && elm != &l.root; elm = elm.next {
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
    for e, elm := 0, l.root.next; elm != nil && elm != &l.root; e, elm = e+1, elm.next {
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
    l.Init()
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
        for e, currentElement := 0, l.root.next; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
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
    if index == 0 {
        l.Prepend(values...)
        return
    }
    if !l.withinRange(index) {
        // Append
        if index == l.size {
            l.Add(values...)
        }
        return
    }
    elm := l.root.next
    for e := 0; e < index - 1; e++ {
        elm = elm.next
    }
    for _, value := range values {
        newElm := &element{value: value}
        l.insert(newElm, elm)
        elm = newElm
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
    elm := l.root.next
    for e := 0; e < index; e++ {
        elm = elm.next
    }
    elm.value = value
}

func (l *List) String() string {
    str := "DoublyLinkedList\n"
    values := []string{}
    for element := l.root.next; element != nil && element != &l.root; element = element.next {
        values = append(values, fmt.Sprintf("%v", element.value))
    }
    str += strings.Join(values, ", ")
    return str
}

func (l *List) insert(e, at *element) {
    n := at.next
    at.next = e
    e.prev = at
    e.next = n
    n.prev = e
    l.size++
}

func (l *List) remove(e *element) {
    e.prev.next = e.next
    e.next.prev = e.prev
    e.next = nil // avoid memory leaks
    e.prev = nil // avoid memory leaks
    l.size--
}

func (l *List) withinRange(index int) bool {
    return index >= 0 && index < l.size
}
