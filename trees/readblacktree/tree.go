package readblacktree

import (
    "fmt"
    "github.com/welllog/data-structures/utils"
)

type color bool

const (
    black, red color = true, false
)

/**
红黑树特征：
1. 根节点为黑色
2. 叶子节点(nil)为黑色
3. 红色节点的子节点必须为黑色
4. 一个节点到该节点的子孙节点的所有路径包含相同数目的黑节点
*/
type Tree struct {
    root *Node
    size int
    comparator utils.Comparator
}

type Node struct {
    key interface{}
    value interface{}
    color color
    left *Node
    right *Node
    parent *Node
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

/**
插入节点：
1. 插入的节点(除根节点)一定为红色
2  父节点为黑色，什么都不做
3. case1: 父节点、叔父节点都为红色，将父节点，叔父节点改为黑色。将祖父节点改为红色，并以祖父节点为当前节点
4. case2: 父节点为红色，叔父节点为黑色，当前节点为父节点的右孩子，父节点为祖父节点左孩子，则以父节点为当前节点并左旋
当前节点为父节点左孩子，父节点为祖父节点右孩子，则以父节点进行右旋。 得到case3
5. case3: 父节点为红色，叔父节点为黑色，将父节点改为黑色，祖父节点改为红色，当前节点为左孩子，以祖父节点为当前节点进行右旋，
当前节点为右孩子，以祖父节点进行左旋
 */
func (t *Tree) Put(key, value interface{}) {
    var insertNode *Node
    if t.root == nil {
        // 防止key非comparator定义的类型
        t.comparator(key, key)
        t.root = &Node{key: key, value: value, color: red}
        insertNode = t.root
    } else {
        node := t.root
        loop := true
        for loop {
            c := t.comparator(key, node.key)
            switch {
            case c == 0:
                node.key = key
                node.value = value
                return
            case c < 0:
                if node.left == nil {
                    node.left = &Node{key: key, value: value, color: red, parent: node}
                    insertNode = node.left
                    loop = false
                } else {
                    node = node.left
                }
            case c > 0:
                if node.right == nil {
                    node.right = &Node{key: key, value: value, color: red, parent: node}
                    insertNode = node.right
                    loop = false
                } else {
                    node = node.right
                }
            }
        }
        insertNode.parent = node
    }
    t.insertCase1(insertNode)
    t.size++
}

/**
删除节点：（1无子节点，2有一个子节点，3有两个子节点）3可转化为1，2
1. 被删除节点有两个孩子，找到这个节点的中序后继，进行替换。转换为删除后继节点（该后继节点无子节点或只有一个子节点）
...
 */
func (t *Tree) Remove(key interface{}) {
    var child *Node
    node := t.lookup(key)
    if node == nil {
        return
    }
    if node.left != nil && node.right != nil { // 两个子节点，先进行替换
        pred := node.left.maximumNode()
        node.key = pred.key
        node.value = pred.value
        node = pred
    }
    // 两个子节点为空，将任意一个作为孩子
    if node.left == nil {
        child = node.right
    } else {
        child = node.left
    }
    // 当前节点为红色，不用改变属性。 当前节点为黑色，子节点为红色，修改子节点为黑色。 当前节点子节点都为黑色，分为六种情况进行处理
    if node.color == black {
        if nodeColor(child) == red {
            child.color = black
        } else {
            t.deleteCase1(node)
        }
    }
    t.replaceNode(node, child)
    
    t.size--
}

func (t *Tree) Get(key interface{}) (value interface{}, found bool) {
    node := t.lookup(key)
    if node != nil {
        return node.value, true
    }
    return nil, false
}

func (t *Tree) Empty() bool {
    return t.size == 0
}

func (t *Tree) Size() int {
    return t.size
}

func (t *Tree) Keys() []interface{} {
    keys := make([]interface{}, t.size)
    it := t.Iterator()
    for i := 0; it.Next(); i++ {
        keys[i] = it.Key()
    }
    return keys
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
    node := t.root
    child := node.left
    for child != nil {
        node = child
        child = node.left
    }
    return node
}

func (t *Tree) Right() *Node {
    node := t.root
    child := node.right
    for child != nil {
        node = child
        child = node.right
    }
    return node
}

// <= key 的最大值
func (t *Tree) Floor(key interface{}) (floor *Node, found bool) {
    node := t.root
    for node != nil {
        c := t.comparator(key, node.key)
        switch {
        case c == 0:
            return node, true
        case c < 0:
            node = node.left
        case c > 0:
            floor, found = node, true
            node = node.right
        }
    }
    if found {
        return floor, true
    }
    return nil, false
}

// >= key的最小值
func (t *Tree) Ceiling(key interface{}) (ceiling *Node, found bool) {
    node := t.root
    for node != nil {
        c := t.comparator(key, node.key)
        switch {
        case c == 0:
            return node, true
        case c < 0:
            ceiling, found = node, true
            node = node.left
        case c > 0:
            node = node.right
        }
    }
    if found {
        return ceiling, true
    }
    return nil, false
}

func (t *Tree) Clear() {
    t.root = nil
    t.size = 0
}

func (t *Tree) String() string {
    str := "RedBlackTree\n"
    if !t.Empty() {
        output(t.root, "", true, &str)
    }
    return str
}

func output(node *Node, prefix string, isTail bool, str *string) {
    if node.right != nil {
        newPrefix := prefix
        if isTail {
            newPrefix += "│   "
        } else {
            newPrefix += "    "
        }
        output(node.right, newPrefix, false, str)
    }
    *str += prefix
    if isTail {
        *str += "└── "
    } else {
        *str += "┌── "
    }
    *str += node.String() + "\n"
    if node.left != nil {
        newPrefix := prefix
        if isTail {
            newPrefix += "    "
        } else {
            newPrefix += "│   "
        }
        output(node.left, newPrefix, true, str)
    }
}

func (t *Tree) lookup(key interface{}) *Node {
    node := t.root
    for node != nil {
        c := t.comparator(key, node.key)
        switch {
        case c == 0:
            return node
        case c > 0:
            node = node.right
        case c < 0:
            node = node.left
        }
    }
    return nil
}

func (t *Tree) replaceNode(old *Node, new *Node) {
    if old.parent == nil {
        t.root = new
    } else {
        if old == old.parent.left {
            old.parent.left = new
        } else {
            old.parent.right = new
        }
    }
    if new != nil {
        new.parent = old.parent
    }
}


// 根节点为黑色
func (t *Tree) insertCase1(node *Node) {
    if node.parent == nil {
        node.color = black
        return
    }
    t.insertCase2(node)
}

// 插入节点的父节点为黑色，不用处理
func (t *Tree) insertCase2(node *Node) {
    if node.parent.color == black {
        return
    }
    t.insertCase3(node)
}

// 父节点、叔父节点都为红色，将父节点，叔父节点改为黑色。将祖父节点改为红色，并以祖父节点为当前节点
func (t *Tree) insertCase3(node *Node) {
    uncle := node.uncle()
    if nodeColor(uncle) == red {
        node.parent.color = black
        uncle.color = black
        grandparent := node.grandparent()
        grandparent.color = red
        t.insertCase1(grandparent)
    } else {
        t.insertCase4(node)
    }
}

// 叔父节点为黑色。节点为右孩子，父节点为左孩子，以父节点左旋；或者节点为左孩子，父节点为右孩子，以父节点右旋。转化为下一种情况
func (t *Tree) insertCase4(node *Node) {
    grandparent := node.grandparent()
    if node == node.parent.right && node.parent == grandparent.left {
        t.rotateLeft(node.parent)
        node = node.left
    } else if node == node.parent.left && node.parent == grandparent.right {
        t.rotateRight(node.parent)
        node = node.right
    }
    t.insertCase5(node)
}

// 父节点为红色，叔父节点为黑色。将父节点置为黑色，祖父节点置为红色，以祖父节点进行旋转
func (t *Tree) insertCase5(node *Node) {
    node.parent.color = black
    grandparent := node.grandparent()
    grandparent.color = red
    if node == node.parent.left && node.parent == grandparent.left {
        t.rotateRight(grandparent)
    } else if node == node.parent.right && node.parent == grandparent.right {
        t.rotateLeft(grandparent)
    }
}

// 该节点为根，属性不变
func (t *Tree) deleteCase1(node *Node) {
    if node.parent == nil {
        return
    }
    t.deleteCase2(node)
}

// 兄弟节点为红色，父节点做左旋或右旋，兄弟节点变为祖父，对调父节点跟祖父节点的颜色，接着处理。旋转后新兄弟节点一定为黑色
func (t *Tree) deleteCase2(node *Node) {
    sibling := node.sibling()
    if nodeColor(sibling) == red {
        node.parent.color = red // 父节点原来一定为黑色,设为红色
        sibling.color = black
        if node == node.parent.left {
            t.rotateLeft(node.parent)
        } else {
            t.rotateRight(node.parent)
        }
    }
    t.deleteCase3(node)
}

// 父节点，兄弟节点，兄弟节点的子节点都为黑色。将兄弟节点绘为红色,在父节点上重新平衡
func (t *Tree) deleteCase3(node *Node) {
    sibling := node.sibling()
    if nodeColor(node.parent) == black &&
        nodeColor(sibling) == black &&
        nodeColor(sibling.left) == black &&
        nodeColor(sibling.right) == black {
        sibling.color = red
        t.deleteCase1(node.parent)
    } else {
        t.deleteCase4(node)
    }
}

// 父节点为红色，兄弟节点跟兄弟节点的子节点都为黑色，将父节点置为黑色，兄弟节点置为红色,重新达到平衡
func (t *Tree) deleteCase4(node *Node) {
    sibling := node.sibling()
    if nodeColor(node.parent) == red &&
        nodeColor(sibling) == black &&
        nodeColor(sibling.left) == black &&
        nodeColor(sibling.right) == black {
        sibling.color = red
        node.parent.color = black
    } else {
        t.deleteCase5(node)
    }
}

// 兄弟节点为黑色。当前节点为左孩子，兄弟节点左孩子为红色；当前节点为右孩子，兄弟节点右孩子为红色，兄弟节点进行右旋或左旋，转为下一种情况
func (t *Tree) deleteCase5(node *Node) {
    sibling := node.sibling()
    // 节点为左孩子，兄弟节点为黑色，兄弟节点的左孩子为红色，在兄弟节点上做右旋
    // 兄弟节点的左孩子成为了兄弟节点的父节点，当前节点的新兄弟。交换旧的兄弟节点跟新兄弟节点的颜色
    if node == node.parent.left &&
        nodeColor(sibling) == black &&
        nodeColor(sibling.left) == red &&
        nodeColor(sibling.right) == black {
        sibling.color = red
        sibling.left.color = black
        t.rotateRight(sibling)
    } else if node == node.parent.right &&
        nodeColor(sibling) == black &&
        nodeColor(sibling.right) == red &&
        nodeColor(sibling.left) == black {
        sibling.color = red
        sibling.right.color = black
        t.rotateLeft(sibling)
    }
    // 进入下一种处理
    t.deleteCase6(node)
}

// 兄弟节点为黑色。当前节点为左孩子，兄弟节点的右孩子为红色；当前节点为右孩子，兄弟节点左孩子为红色.在父节点上左旋或右旋
func (t *Tree) deleteCase6(node *Node) {
    sibling := node.sibling()
    sibling.color = nodeColor(node.parent)
    node.parent.color = black
    if node == node.parent.left && nodeColor(sibling.right) == red {
        sibling.right.color = black
        t.rotateLeft(node.parent)
    } else if nodeColor(sibling.left) == red {
        sibling.left.color = black
        t.rotateRight(node.parent)
    }
}

func (t *Tree) rotateLeft(node *Node) {
    right := node.right
    if node.parent != nil {
        if node == node.parent.left {
            node.parent.left = right
        } else {
            node.parent.right = right
        }
    } else {
        t.root = right
    }
    right.parent = node.parent
    node.right = right.left
    if right.left != nil {
        right.left.parent = node
    }
    right.left = node
    node.parent = right
}

func (t *Tree) rotateRight(node *Node) {
    left := node.left
    if node.parent != nil {
        if node == node.parent.left {
            node.parent.left = left
        } else {
            node.parent.right = left
        }
    } else {
        t.root = left
    }
    left.parent = node.parent
    node.left = left.right
    if left.right != nil {
        left.right.parent = node
    }
    left.right = node
    node.parent = left
}

func (n *Node) grandparent() *Node {
    if n != nil && n.parent != nil {
        return n.parent.parent
    }
    return nil
}

func (n *Node) uncle() *Node {
    if n == nil || n.parent == nil || n.parent.parent == nil {
        return nil
    }
    return n.parent.sibling()
}

func (n *Node) sibling() *Node {
    if n == nil || n.parent == nil {
        return nil
    }
    if n == n.parent.left {
        return n.parent.right
    }
    return n.parent.left
}

func (n *Node) maximumNode() *Node {
    if n == nil {
        return nil
    }
    for n.right != nil {
        n = n.right
    }
    return n
}

func (n *Node) String() string {
    return fmt.Sprintf("%v", n.key)
}

func nodeColor(node *Node) color {
    if node == nil {
        return black
    }
    return node.color
}
