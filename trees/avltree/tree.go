package avltree

import (
    "fmt"
    "github.com/welllog/data-structures/trees"
    "github.com/welllog/data-structures/utils"
)

func assertTreeImplementation() {
    var _ trees.Tree = (*Tree)(nil)
}

type Tree struct {
    root *Node
    comparator utils.Comparator
    size int
}

type Node struct {
    key interface{}
    value interface{}
    parent *Node
    children [2]*Node
    height int
}

func (n *Node) GetValue() interface{} {
    return n.value
}

func (n *Node) Left() *Node {
    return n.children[0]
}

func (n *Node) Right() *Node {
    return n.children[1]
}

func (n *Node) String() string {
    return fmt.Sprintf("%v", n.key)
}

func NewWith(comparator utils.Comparator) *Tree {
    return &Tree{comparator: comparator}
}

func NewWithIntComparator() *Tree {
    return &Tree{comparator: utils.IntComparator}
}

func NewWithStringComparator() *Tree {
    return &Tree{comparator: utils.StringComparator}
}

func (t *Tree) IsBalanced() bool {
    return t.isBalanced(t.root)
}

func (t *Tree) isBalanced(n *Node) bool {
    if n == nil {
        return true
    }
    
    balanceFactor := t.getBalanceFactor(n)
    if balanceFactor > 1 || balanceFactor < -1 {
        return false
    }
    return t.isBalanced(n.children[0]) && t.isBalanced(n.children[1])
}

func (t *Tree) IsBST() bool {
    var keys []interface{}
    inOrder(t.root, keys)
    
    for i := 1; i < len(keys); i++ {
        if t.comparator(keys[i-1], keys[i]) == 1 {
            return false
        }
    }
    return true
}

func (t *Tree) Put(key interface{}, val interface{}) {
    t.put(key, val, nil, &t.root)
}

func (t *Tree) Remove(key interface{}) {
    t.remove(key, &t.root)
}

func (t *Tree) Get(key interface{}) (value interface{}, found bool) {
    n := t.root
    for n != nil {
        c := t.comparator(key, n.key)
        switch {
        case c == 0:
            return n.value, true
        case c < 0:
            n = n.children[0]
        case c > 0:
            n = n.children[1]
        }
    }
    return nil, false
}

func (t *Tree) GetNode(key interface{}) *Node {
    n := t.root
    for n != nil {
        c := t.comparator(key, n.key)
        switch {
        case c == 0:
            return n
        case c < 0:
            n = n.children[0]
        case c > 0:
            n = n.children[1]
        }
    }
    return nil
}

func (t *Tree) Empty() bool {
    return t.size == 0
}

func (t *Tree) Size() int {
    return t.size
}

func (t *Tree) Clear() {
    t.root = nil
    t.size = 0
}

func (t *Tree) Values() []interface{} {
    values := make([]interface{}, t.size)
    it := t.Iterator()
    for i := 0; it.Next(); i++ {
        values[i] = it.Value()
    }
    return values
}

func (t *Tree) Left() *Node {
    return t.bottom(0)
}

func (t *Tree) Right() *Node {
    return t.bottom(1)
}

// 返回小于等于key的最大节点
func (t *Tree) Floor(key interface{}) (floor *Node, found bool) {
    found = false
    n := t.root
    for n != nil {
        c := t.comparator(key, n.key)
        switch {
        case c == 0:
            return n, true
        case c < 0:
            n = n.children[0]
        case c > 0:
            floor, found = n, true
            n = n.children[1]
        }
    }
    if found {
        return
    }
    return nil, false
}

// 返回大于等于key的最大节点
func (t *Tree) Ceiling(key interface{}) (floor *Node, found bool) {
    found = false
    n := t.root
    for n != nil {
        c := t.comparator(key, n.key)
        switch {
        case c == 0:
            return n, true
        case c < 0:
            floor, found = n, true
            n = n.children[0]
        case c > 0:
            n = n.children[1]
        }
    }
    if found {
        return
    }
    return nil, false
}

func (t *Tree) String() string {
    str := "AVLTree\n"
    if !t.Empty() {
        output(t.root, "", true, &str)
    }
    return str
}

func (t *Tree) put(key interface{}, val interface{}, parent *Node, np **Node) {
    // 优化方式：采用栈替代递归，且如果当前节点的高度没有改变，则停止向上回溯父节点
    n := *np
    if n == nil {
        t.size++
        *np = &Node{key: key, value: val, parent: parent, height: 1}
        return
    }
    
    c := t.comparator(key, n.key)
    if c == 0 { // 替换当前节点值
        n.key = key
        n.value = val
        return
    }
    
    if c < 0 {
        c = -1
    } else {
        c = 1
    }
    
    a := (c + 1) / 2
    t.put(key, val, n, &n.children[a])
    
    lh := t.getHeight(n.children[0])
    rh := t.getHeight(n.children[1])
    n.height = Max(lh, rh) + 1 // 更新高度
    
    b := lh - rh
    if b >= -1 && b <= 1 { // 未失衡
        return
    }
    
    // a == 0 左子树高   a == 1 右子树高
    if t.getBalanceFactor(n.children[a]) == c { // a:0 => c:-1, LR     a:1 => c:1, RL
        n.children[a] = t.rotate(a^1, n.children[a])
    } // LL RR
    *np = t.rotate(a, n)
    
    
    //if b > 1 { // 左子树高
    //    if t.getBalanceFactor(n.children[0]) > 0 { // LL
    //        *np = t.rotate(0, n)
    //    } else { // LR 左子左旋，再右旋
    //        n.children[0] = t.rotate(1, n.children[0])
    //        *np = t.rotate(0, n)
    //    }
    //} else { // 右子树高
    //    if t.getBalanceFactor(n.children[1]) < 0 { // RR
    //        *np = t.rotate(1, n)
    //    } else { // RL 右子右旋再左旋
    //        n.children[1] = t.rotate(0, n.children[1])
    //        *np = t.rotate(1, n)
    //    }
    //}
}

func (t *Tree) remove(key interface{}, np **Node) {
    // 优化方式：采用栈替代递归，且如果当前节点的高度没有改变，且平衡值在[-1, 1]区间则停止向上回溯父节点
    n := *np
    if n == nil {
        return
    }
    
    c := t.comparator(key, n.key)
    if c == 0 {
        t.size--
        if n.children[0] == nil { // 左子树为空
            if n.children[1] != nil {
                n.children[1].parent = n.parent
            }
            *np = n.children[1]
        } else if n.children[1] == nil { // 右子树为空
            n.children[0].parent = n.parent
            *np = n.children[0]
        } else { // 左右子树都不为空，找到该节点的右子树最小节点 替换该节点
            brn := t.min(n.children[1])
            t.remove(brn.key, &n.children[1])
            n.key = brn.key
            n.value = brn.value
        }
        return
    }
    
    if c < 0 {
        c = -1
    } else {
        c = 1
    }
    
    a := (c + 1) / 2
    t.remove(key, &n.children[a])
    
    lh := t.getHeight(n.children[0])
    rh := t.getHeight(n.children[1])
    n.height = Max(lh, rh) + 1 // 更新高度
    
    b := lh - rh
    if b >= -1 && b <= 1 { // 未失衡
        return
    }
    
    // a == 0 移除左子树节点失衡，右子树高   a == 1 移除右子树节点失衡，左子树高
    if t.getBalanceFactor(n.children[a^1]) == -c { // a:0 => c:-1, RL     a:1 => c:1, LR
        n.children[a^1] = t.rotate(a, n.children[a^1])
    } // RR LL
    *np = t.rotate(a^1, n)
    
    //if b > 1 { // 左子树高
    //    if t.getBalanceFactor(n.children[0]) > 0 { // LL
    //        *np = t.rotate(0, n)
    //    } else { // LR
    //        n.children[0] = t.rotate(1, n.children[0])
    //        *np = t.rotate(0, n)
    //    }
    //} else { // 右子树高
    //    if t.getBalanceFactor(n.children[1]) < 0 { // RR
    //        *np = t.rotate(1, n)
    //    } else { // RL
    //        n.children[1] = t.rotate(0, n.children[1])
    //        *np = t.rotate(1, n)
    //    }
    //}
}

func (t *Tree) min(n *Node) *Node {
    for {
        if n.children[0] == nil {
            return n
        }
        n = n.children[0]
    }
}

func (t *Tree) getHeight(n *Node) int {
    if n == nil {
        return 0
    }
    return n.height
}

func (t *Tree) getBalanceFactor(n *Node) int {
    if n == nil {
        return 0
    }
    return t.getHeight(n.children[0]) - t.getHeight(n.children[1])
}

// 0 右旋   1 左旋
func (t *Tree) rotate(flag int, n *Node) *Node {
    cn := n.children[flag]
    n.children[flag] = cn.children[flag^1]
    if n.children[flag] != nil {
        n.children[flag].parent = n
    }
    cn.children[flag^1] = n
    cn.parent = n.parent
    n.parent = cn
    
    // 更新高度
    n.height = Max(t.getHeight(n.children[0]), t.getHeight(n.children[1])) + 1
    cn.height = Max(t.getHeight(cn.children[0]), t.getHeight(cn.children[1])) + 1
    
    return cn
}

//func (t *Tree) rightRotate(n *Node) *Node {
//    cn := n.children[0]
//    n.children[0] = cn.children[1]
//    if n.children[0] != nil {
//        n.children[0].parent = n
//    }
//    cn.children[1] = n
//    cn.parent = n.parent
//    n.parent = cn
//
//    // 更新高度
//    n.height = Max(t.getHeight(n.children[0]), t.getHeight(n.children[1])) + 1
//    cn.height = Max(t.getHeight(cn.children[0]), t.getHeight(cn.children[1])) + 1
//
//    return cn
//}
//
//func (t *Tree) leftRotate(n *Node) *Node {
//    cn := n.children[1]
//    n.children[1] = cn.children[0]
//    if n.children[1] != nil {
//        n.children[1].parent = n
//    }
//    cn.children[0] = n
//    cn.parent = n.parent
//    n.parent = cn
//
//    // 更新高度
//    n.height = Max(t.getHeight(n.children[0]), t.getHeight(n.children[1])) + 1
//    cn.height = Max(t.getHeight(cn.children[0]), t.getHeight(cn.children[1])) + 1
//    return cn
//}

func (t *Tree) bottom(d int) *Node {
    n := t.root
    if n == nil {
        return nil
    }
    
    for c := n.children[d]; c != nil; c = n.children[d] {
        n = c
    }
    return n
}

func (n *Node) Prev() *Node {
    return n.walk1(0)
}

// Next returns the next element in an inorder
// walk of the AVL tree.
func (n *Node) Next() *Node {
    return n.walk1(1)
}

func (n *Node) walk1(a int) *Node {
    if n == nil {
        return nil
    }
    
    if n.children[a] != nil {
        n = n.children[a]
        for n.children[a^1] != nil {
            n = n.children[a^1]
        }
        return n
    }
    
    p := n.parent
    for p != nil && p.children[a] == n {
        n = p
        p = p.parent
    }
    return p
}

func output(node *Node, prefix string, isTail bool, str *string) {
    if node.children[1] != nil {
        newPrefix := prefix
        if isTail {
            newPrefix += "│   "
        } else {
            newPrefix += "    "
        }
        output(node.children[1], newPrefix, false, str)
    }
    *str += prefix
    if isTail {
        *str += "└── "
    } else {
        *str += "┌── "
    }
    *str += node.String() + "\n"
    if node.children[0] != nil {
        newPrefix := prefix
        if isTail {
            newPrefix += "    "
        } else {
            newPrefix += "│   "
        }
        output(node.children[0], newPrefix, true, str)
    }
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func inOrder(n *Node, keys []interface{}) {
    if n == nil {
        return
    }
    
    inOrder(n.children[0], keys)
    keys = append(keys, n.key)
    inOrder(n.children[1], keys)
}