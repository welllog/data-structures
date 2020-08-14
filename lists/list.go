package lists

import (
    "github.com/welllog/data-structures/containers"
    "github.com/welllog/data-structures/utils"
)

type List interface {
    Get(index int) (interface{}, bool)
    Remove(index int)
    Add(values ...interface{})
    Contains(values ...interface{}) bool
    Sort(comparator utils.Comparator)
    Swap(index1, index2 int)
    Insert(index int, values ...interface{})
    Set(index int, value interface{})
    
    containers.Container
}