package readblacktree

import (
    "github.com/welllog/data-structures/trees"
)

func (t *Tree) BreadFirst(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    queue := make([]*Node, 0, 40)
    queue = append(queue, t.root)
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        if !f(node.key, node.value) {
            return
        }
        
        if node.left != nil {
            queue = append(queue, node.left)
        }
        if node.right != nil {
            queue = append(queue, node.right)
        }
    }
}

// 前序遍历
/**
func PreOrder(node *Node) {
    fmt.Println(node.key, node.value)
    PreOrder(node.left)
    PreOrder(node.right)
}
改为非递归，手动模拟栈
 */
func (t *Tree) PreOrder(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    stack := make([]*Node, 0, 40)
    stack = append(stack, t.root)
    for len(stack) > 0 {
        l := len(stack)
        node := stack[l - 1]
        if !f(node.key, node.value) {
            return
        }
        
        stack = stack[:l - 1]
        
        if node.right != nil {
            stack = append(stack, node.right)
        }
        
        if node.left != nil {
            stack = append(stack, node.left)
        }
    }
}

// 中序遍历
/**
    a
   / \
   b  c
  / \ / \
 d  e f  g
func InOrder(node *Node) {
    InOrder(node.left)
    fmt.Println(node.key, node.value)
    InOrder(node.right)
}
 */
func (t *Tree) InOrder(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    stack := make([]*Node, 0, 40)
    node := t.root
    for len(stack) > 0 || node != nil {
        for node != nil{
            stack = append(stack, node)
            node = node.left
        }
        l := len(stack)
        node = stack[l - 1]
        
        if !f(node.key, node.value) {
            return
        }
        stack = stack[:l - 1]

        node = node.right
    }
}

/**
    a
   / \
   b  c
  / \ / \
 d  e f  g
func TailOrder(node *Node) {
    TailOrder(node.left)
    TailOrder(node.right)
    fmt.Println(node.key, node.value)
}
 */
func (t *Tree) TailOrder(f trees.NodeHandler) {
    if t.root == nil {
        return
    }
    stack := make([]*Node, 0, 40)
    node := t.root
    var cnode *Node
    for len(stack) > 0 || node != nil {
        for node != nil{
            stack = append(stack, node)
            node = node.left
        }
        l := len(stack)
        node = stack[l - 1]
        if node.right == nil || node.right == cnode {
            if !f(node.key, node.value) {
                return
            }
            stack = stack[:l - 1]
            cnode = node
            node = nil
        } else {
            node = node.right
        }
    }
}
