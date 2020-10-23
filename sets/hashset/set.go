package hashset

import (
    "fmt"
    "strings"
)

type Set struct {
    items map[interface{}]struct{}
}

var itemExists = struct{}{}

func New(values ...interface{}) *Set {
    set := &Set{items: make(map[interface{}]struct{})}
    if len(values) > 0 {
        set.Add(values...)
    }
    return set
}

func (set *Set) Add(items ...interface{}) {
    for _, item := range items {
        set.items[item] = itemExists
    }
}

func (set *Set) Remove(items ...interface{}) {
    for _, item := range items {
        delete(set.items, item)
    }
}

func (set *Set) Contains(items ...interface{}) bool {
    for _, item := range items {
        if _, contains := set.items[item]; !contains {
            return false
        }
    }
    return true
}

func (set *Set) Empty() bool {
    return set.Size() == 0
}

func (set *Set) Size() int {
    return len(set.items)
}

func (set *Set) Clear() {
    set.items = make(map[interface{}]struct{})
}

func (set *Set) Values() []interface{} {
    values := make([]interface{}, set.Size())
    count := 0
    for item := range set.items {
        values[count] = item
        count++
    }
    return values
}

func (set *Set) String() string {
    str := "HashSet\n"
    items := []string{}
    for k := range set.items {
        items = append(items, fmt.Sprintf("%v", k))
    }
    str += strings.Join(items, ", ")
    return str
}
