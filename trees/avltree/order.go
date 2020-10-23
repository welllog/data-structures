package avltree

import "github.com/welllog/data-structures/trees"

// 广度优先
func (t *Tree) BreadFirst(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    queue := make([]*Node, 0, 40)
    queue = append(queue, t.root)
    
    var n *Node
    for len(queue) > 0 {
        n = queue[0]
        queue = queue[1:]
        
        if !f(n.key, n.value) {
            return
        }
        
        if n.children[0] != nil {
            queue = append(queue, n.children[0])
        }
        
        if n.children[1] != nil {
            queue = append(queue, n.children[1])
        }
    }
}

// 前序遍历
func (t *Tree) PreOrder(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    stack := make([]*Node, 0, 40)
    stack = append(stack, t.root)
    
    var (
        n *Node
        last int
    )
    for len(stack) > 0 {
        last = len(stack) - 1
        n = stack[last]
        
        if !f(n.key, n.value) {
            return
        }
        stack = stack[:last]
    
        if n.children[1] != nil {
            stack = append(stack, n.children[1])
        }
        
        if n.children[0] != nil {
            stack = append(stack, n.children[0])
        }
    }
}

// 中序遍历
func (t *Tree) InOrder(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    stack := make([]*Node, 0, 40)
    var (
        n *Node
        last int
    )
    n = t.root
    for len(stack) > 0 || n != nil {
        for n != nil {
            stack = append(stack, n)
            n = n.children[0]
        }
        
        last = len(stack) - 1
        n = stack[last]
        if !f(n.key, n.value) {
            return
        }
        
        stack = stack[:last]
        
        n = n.children[1]
    }
}

// 后续遍历
func (t *Tree) TailOrder(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    stack := make([]*Node, 0, 40)
    var (
        n *Node
        cn *Node
        last int
    )
    n = t.root
    for len(stack) > 0 || n != nil {
        for n != nil {
            stack = append(stack, n)
            n = n.children[0]
        }
    
        last = len(stack) - 1
        n = stack[last]
        
        if n.children[1] == nil || n.children[1] == cn {
            if !f(n.key, n.value) {
                return
            }
            stack = stack[:last]
            cn = n
            n = nil
        } else {
            n = n.children[1]
        }
    }
}
