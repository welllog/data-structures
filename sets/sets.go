package sets

import "github.com/welllog/data-structures/containers"

type Set interface {
    Add(elements ...interface{})
    Remove(elements ...interface{})
    Contains(elements ...interface{}) bool
    
    containers.Container
}
