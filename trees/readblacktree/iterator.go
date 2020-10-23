package readblacktree

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
    return Iterator{tree: t, position: begin}
}

func (i *Iterator) Next() bool {
    if i.position == begin {
        left := i.tree.Left()
        if left == nil {
            goto end
        }
        i.node = left
        goto between
    }
    if i.position == end {
        goto end
    }
    if i.node.right != nil {
        i.node = i.node.right
        for i.node.left != nil {
            i.node = i.node.left
        }
        goto between
    }
    if i.node.parent != nil {
        node := i.node
        for i.node.parent != nil {
            i.node = i.node.parent
            if i.tree.comparator(node.key, i.node.key) <= 0 {
                goto between
            }
        }
    }

end:
    i.node = nil
    i.position = end
    return false

between:
    i.position = between
    return true
}

func (i *Iterator) Prev() bool {
    if i.position == begin {
        goto begin
    }
    if i.position == end {
        right := i.tree.Right()
        if right == nil {
            goto begin
        }
        i.node = right
        goto between
    }
    if i.node.left != nil {
        i.node = i.node.left
        for i.node.right != nil {
            i.node = i.node.right
        }
        goto between
    }
    if i.node.parent != nil {
        node := i.node
        for i.node.parent != nil {
            i.node = i.node.parent
            if i.tree.comparator(node.key, i.node.key) >= 0 {
                goto between
            }
        }
    }

begin:
    i.node = nil
    i.position = begin
    return false

between:
    i.position = between
    return true
}

func (i *Iterator) Value() interface{} {
    return i.node.value
}

func (i *Iterator) Key() interface{} {
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