package trees

import "github.com/welllog/data-structures/containers"

type Tree interface {
    containers.Container
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
}

type NodeHandler func(key, value interface{}) bool
