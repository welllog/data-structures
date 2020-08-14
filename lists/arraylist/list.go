package arraylist

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
    elements []interface{}
    size     int
}

const (
    growthFactor = float32(2.0)  // growth by 100%
    shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

func New(values ...interface{}) *List {
    list := &List{}
    if len(values) > 0 {
        list.Add(values...)
    }
    return list
}

func (l *List) Add(values ...interface{}) {
    l.growBy(len(values))
    for _, value := range values {
        l.elements[l.size] = value
        l.size++
    }
}

func (l *List) Get(index int) (interface{}, bool) {
    if !l.withinRange(index) {
        return nil, false
    }
    
    return l.elements[index], true
}

func (l *List) Remove(index int) {
    if !l.withinRange(index) {
        return
    }
    
    l.elements[index] = nil
    copy(l.elements[index:], l.elements[index+1:l.size]) // shift to the left by one (slow operation, need ways to optimize this)
    l.size--
    
    l.shrink()
}

func (l *List) Contains(values ...interface{}) bool {
    for _, searchValue := range values {
        found := false
        for i := 0; i < l.size; i++ {
            if searchValue == l.elements[i] {
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
    newElements := make([]interface{}, l.size, l.size)
    copy(newElements, l.elements[:l.size])
    return newElements
}

func (l *List) IndexOf(value interface{}) int {
    for i := 0; i < l.size; i++ {
        if value == l.elements[i] {
            return i
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
    l.elements = []interface{}{}
}

func (l *List) Sort(comparator utils.Comparator) {
    if len(l.elements) < 2 {
        return
    }
    utils.Sort(l.elements[:l.size], comparator)
}

func (l *List) Swap(i, j int) {
    if l.withinRange(i) && l.withinRange(j) {
        l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
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
    
    length := len(values)
    l.growBy(length)
    l.size += length
    copy(l.elements[index+length:], l.elements[index:l.size-length])
    copy(l.elements[index:], values)
}

func (l *List) Set(index int, value interface{}) {
    if !l.withinRange(index) {
        if index == l.size {
            l.Add(value)
        }
        return
    }
    
    l.elements[index] = value
}

func (l *List) String() string {
    str := "ArrayList\n"
    values := []string{}
    for _, value := range l.elements[:l.size] {
        values = append(values, fmt.Sprintf("%v", value))
    }
    str += strings.Join(values, ", ")
    return str
}

func (l *List) withinRange(index int) bool {
    return index >= 0 && index < l.size
}

func (l *List) resize(cap int) {
    newElements := make([]interface{}, cap, cap)
    copy(newElements, l.elements)
    l.elements = newElements
}

// 必要时扩展数组，即增加n个元素就会达到容量。
func (l *List) growBy(n int) {
    //  当容量达到时，按增长系数growthFactor增长，并增加元素数量。
    currentCapacity := cap(l.elements)
    if l.size+n >= currentCapacity {
        newCapacity := int(growthFactor * float32(currentCapacity+n))
        l.resize(newCapacity)
    }
}

// 必要时收缩数组，即当大小为当前容量的shrinkFactor百分比时。
func (l *List) shrink() {
    if shrinkFactor == 0.0 {
        return
    }
    // 当尺寸达到shrinkFactor*capacity时，收缩
    currentCapacity := cap(l.elements)
    if l.size <= int(float32(currentCapacity)*shrinkFactor) {
        l.resize(l.size)
    }
}
