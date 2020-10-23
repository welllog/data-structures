package readblacktree

import (
    "fmt"
    "testing"
)

func TestRbTree(t *testing.T) {
    tree := NewWithIntComparator()
    tree.Put(5, 5)
    tree.Put(3, 3)
    tree.Put(8, 8)
    tree.Put(2, 2)
    tree.Put(4, 4)
    tree.Put(7, 7)
    tree.Put(9, 9)
    tree.Remove(3)
    fmt.Println(tree.String())
    
    tree.PreOrder(Print)
    fmt.Println()
    tree.InOrder(Print)
    fmt.Println()
    tree.TailOrder(Print)
    fmt.Println()
    
    iter := tree.Iterator()
    for iter.Next() {
        fmt.Print(iter.Key(), " ")
    }
}

func Print(key, value interface{}) bool {
    fmt.Print(key, " ")
    return true
}
