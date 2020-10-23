package avltree

import (
    "fmt"
    "testing"
)

func TestAvlTree(t *testing.T) {
    tree := NewWithIntComparator()
    tree.Put(1, 1)
    tree.Put(2, 2)
    tree.Put(3, 3)
    tree.Put(0, 0)
    tree.Put(-1, -1)
    tree.Put(4, 4)
    tree.Put(5, 5)
    tree.Put(6, 6)
    tree.Put(7, 7)
    tree.Put(8, 8)
    tree.Put(9, 9)
    fmt.Println(tree.String())
    tree.Put(10, 10)
    fmt.Println(tree.String())
    iterator := tree.Iterator()
    for iterator.Next() {
        fmt.Print(iterator.Value(), " ")
    }
    fmt.Println()
    fmt.Print("广度遍历：")
    tree.BreadFirst(func(key, value interface{}) bool {
        fmt.Print(key, " ")
        return true
    })
    fmt.Println()
    fmt.Print("前序遍历：")
    tree.PreOrder(func(key, value interface{}) bool {
        fmt.Print(key, " ")
        return true
    })
    fmt.Println()
    fmt.Print("中序遍历：")
    tree.InOrder(func(key, value interface{}) bool {
        fmt.Print(key, " ")
        return true
    })
    fmt.Println()
    fmt.Print("后序遍历：")
    tree.TailOrder(func(key, value interface{}) bool {
        fmt.Print(key, " ")
        return true
    })
    fmt.Println()
}
