package avltree

import "github.com/welllog/data-structures/containers"

func assertIteratorImplementation() {
    var _ containers.ReverseIteratorWithKey = (*Iterator)(nil)
}

type Iterator struct {
    tree *Tree
    node *Node
    position position
}

type position byte

const (
    begin, between, end position = 0, 1, 2
)

func (t *Tree) Iterator() Iterator {
    return Iterator{tree: t, node: t.root, position: begin}
}

func (i *Iterator) Next() bool {
    switch i.position {
    case begin:
        i.position = between
        i.node = i.tree.Left()
    case between:
        i.node = i.node.Next()
    }
    
    if i.node == nil {
        i.position = end
        return false
    }
    return true
}

func (i *Iterator) Prev() bool {
    switch i.position {
    case end:
        i.position = between
        i.node = i.tree.Right()
    case between:
        i.node = i.node.Prev()
    }
    
    if i.node == nil {
        i.position = begin
        return false
    }
    return true
}

func (i *Iterator) Value() interface{} {
    if i.node == nil {
        return nil
    }
    return i.node.value
}

func (i *Iterator) Key() interface{} {
    if i.node == nil {
        return nil
    }
    return i.node.key
}

func (i *Iterator) Begin() {
    i.node = nil
    i.position = begin
}

func (i *Iterator) End() {
    i.node = nil
    i.position = end
}

func (i *Iterator) First() bool {
    i.Begin()
    return i.Next()
}

func (i *Iterator) Last() bool {
    i.End()
    return i.Prev()
}
