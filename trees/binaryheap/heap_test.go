package binaryheap

import (
    "fmt"
    "testing"
)

func TestHeap(t *testing.T) {
    h := NewWithIntComparator()
    h.Push(5, 6, 7, 8, 9, -1, 1, 1, 2, 3, 4 )
    fmt.Println(h.Pop())
    fmt.Println(h.Pop())
    fmt.Println(h.Pop())
    fmt.Println(h.Pop())
}
