package binaryheap

import (
    "fmt"
    "github.com/welllog/data-structures/lists/arraylist"
    "github.com/welllog/data-structures/trees"
    "github.com/welllog/data-structures/utils"
    "strings"
)

func assertTreeImplementation() {
    var _ trees.Tree = (*Heap)(nil)
}

type Heap struct {
    list *arraylist.List
    comparator utils.Comparator
}

func NewWith(comparator utils.Comparator) *Heap {
    return &Heap{list: arraylist.New(), comparator: comparator}
}

func NewWithIntComparator() *Heap {
    return &Heap{list: arraylist.New(), comparator: utils.IntComparator}
}

func NewWithStringComparator() *Heap {
    return &Heap{list: arraylist.New(), comparator: utils.StringComparator}
}

func (h *Heap) Push(values ...interface{}) {
    l := len(values)
    if l == 0 {
        return
    }
    
    if l == 1 {
        h.list.Add(values[0])
        h.bubbleUp() // 冒泡上浮
        return
    }
    
    // 多个输入时，采用批量构建,效率更高
    for _, value := range values {
        h.list.Add(value)
    }
    index := h.list.Size() / 2 - 1
    for i := index; i >= 0; i-- {
        h.bubbleDownIndex(i)
    }
}

func (h *Heap) Pop() (value interface{}, ok bool) {
    value, ok = h.list.Get(0)
    if !ok {
        return
    }
    
    // 将末尾元素与根节点互换，然后删除末尾节点，根节点向下进行整形
    lastIndex := h.list.Size() - 1
    h.list.Swap(0, lastIndex)
    h.list.Remove(lastIndex)
    h.bubbleDownIndex(0)
    return
}

// 返回堆的顶点，而不进行移除
func (h *Heap) Peek() (value interface{}, ok bool) {
    return h.list.Get(0)
}

func (h *Heap) Size() int {
    return h.list.Size()
}

func (h *Heap) Empty() bool {
    return h.list.Empty()
}

func (h *Heap) Clear() {
    h.list.Clear()
}

func (h *Heap) Values() []interface{} {
    return h.list.Values()
}

func (heap *Heap) String() string {
    str := "BinaryHeap\n"
    values := []string{}
    for _, value := range heap.list.Values() {
        values = append(values, fmt.Sprintf("%v", value))
    }
    str += strings.Join(values, ", ")
    return str
}

// 跟父节点比较，如果不符合二叉堆特性，则进行交换。向上重复该操作
func (h *Heap) bubbleUp() {
    index := h.list.Size() - 1
    for parentIndex := (index - 1) >> 1; index > 0; parentIndex = (index - 1) >> 1 { // parentIndex = (childIndex - 1) / 2
        indexValue, _ := h.list.Get(index)
        parentValue, _ := h.list.Get(parentIndex)
        if h.comparator(parentValue, indexValue) <= 0 {
            break
        }
        h.list.Swap(index, parentIndex)
        index = parentIndex
    }
}

// 跟较大(较小)子节点比较，不符合二叉堆的特性，则进行交换。向下重复该操作
func (h *Heap) bubbleDownIndex(index int) {
    size := h.list.Size()
    for leftIndex := index << 1 + 1; leftIndex < size; leftIndex = index << 1 + 1 { // leftIndex = 2 * parentIndex + 1
        rightIndex := leftIndex + 1   // rightIndex = 2 * parentIndex + 2
        smallIndex := leftIndex
        smallVal,_ := h.list.Get(smallIndex)
        if rightIndex < size {
            rightVal,_ := h.list.Get(rightIndex)
            if h.comparator(rightVal, smallVal) < 0 {
                smallIndex = rightIndex
                smallVal = rightVal
            }
        }
        
        indexVal,_ := h.list.Get(index)
        if h.comparator(smallVal, indexVal) > 0 {
            break
        }
        h.list.Swap(index, smallIndex)
        index = smallIndex
    }
}

func (h *Heap) withinRange(index int) bool {
    return index >= 0 && index < h.list.Size()
}


