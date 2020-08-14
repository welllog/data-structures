package containers

// 迭代器提供有状态的迭代器。

// IteratorWithIndex是有序容器的有状态迭代器，其值可以通过索引获取。
type IteratorWithIndex interface {
    // Next将迭代器移动到下一个元素，如果容器中存在下一个元素，则返回true。
    // 如果Next()返回true，那么可以通过Index()和Value()来获取下一个元素的索引和值。
    // 如果Next()是第一次被调用，那么它将把迭代器指向第一个元素，如果它存在的话。
    // 修改迭代器的状态。
    Next() bool
    
    // Value 返回当前元素的值。
    // 不修改迭代器的状态。
    Value() interface{}
    
    // Index 返回当前元素的索引。
    // 不修改迭代器的状态。
    Index() int
    
    // Begin 将迭代器重置为初始状态（one-before-first）。
    // 如果有的话，调用Next()来获取第一个元素。
    Begin()
    
    // First将迭代器移动到第一个元素，如果容器中存在第一个元素，则返回true。
    // 如果First()返回true，那么可以通过Index()和Value()检索第一个元素的索引和值。
    // 修改迭代器的状态。
    First() bool
}

// IteratorWithKey是一个有状态的迭代器，用于有序容器的元素是键值对。
type IteratorWithKey interface {
    // Next将迭代器移动到下一个元素，如果容器中存在下一个元素，则返回true。
    // 如果Next()返回true，那么可以通过Key()和Value()来获取下一个元素的key和value。
    // 如果Next()是第一次被调用，那么它将把迭代器指向第一个元素，如果它存在的话。
    // 修改迭代器的状态。
    Next() bool
    
    // Value 返回当前元素的值。
    // 不修改迭代器的状态。
    Value() interface{}
    
    // Key 返回当前元素的键。
    // 不修改迭代器的状态。
    Key() interface{}
    
    // Begin 将迭代器重置为初始状态（one-before-first）。
    // 如果有的话，调用Next()来获取第一个元素。
    Begin()
    
    // First将迭代器移动到第一个元素，如果容器中存在第一个元素，则返回true。
    // 如果First()返回true，那么可以通过Key()和Value()检索第一个元素的key和value。
    // 修改迭代器的状态。
    First() bool
}

// ReverseIteratorWithIndex是一个有状态的迭代器，用于有序容器，其值可以通过索引来获取。
//
// 本质上它与IteratorWithIndex相同，但提供了额外的功能。
//
// Prev()函数用于启用反向遍历。
//
// Last()函数用于将迭代器移动到最后一个元素。
//
// End()函数将迭代器移过最后一个元素（one-past-the-end）。
type ReverseIteratorWithIndex interface {
    // Prev将迭代器移动到前一个元素，如果容器中存在前一个元素，则返回true。
    // 如果Prev()返回true，那么可以通过Index()和Value()来获取前一个元素的索引和值。
    // 修改迭代器的状态。
    Prev() bool
    
    // End 将迭代器移到最后一个元素之后（one-past-the-end）。
    // 调用Prev()来获取最后一个元素（如果有的话）。
    End()
    
    // Last将迭代器移动到最后一个元素，如果容器中存在最后一个元素，则返回true。
    // 如果Last()返回true，那么可以通过Index()和Value()检索最后一个元素的索引和值。
    // 修改迭代器的状态。
    Last() bool
    
    IteratorWithIndex
}

// ReverseIteratorWithKey是一个有序容器的状态迭代器，其元素是键值对。
//
// 本质上它与IteratorWithKey相同，但提供了额外的功能。
//
// Prev()函数用于启用反向遍历。
//
// Last()函数，将迭代器移动到最后一个元素。
type ReverseIteratorWithKey interface {
    // Prev将迭代器移动到前一个元素，如果容器中存在前一个元素，则返回true。
    // 如果Prev()返回true，那么可以通过Key()和Value()来获取前一个元素的key和value。
    // 修改迭代器的状态。
    Prev() bool
    
    // End 将迭代器移到最后一个元素之后（one-past-the-end）。
    // 调用Prev()来获取最后一个元素（如果有的话）。
    End()
    
    // Last将迭代器移动到最后一个元素，如果容器中存在最后一个元素，则返回true。
    // 如果Last()返回true，那么可以通过Key()和Value()检索最后一个元素的key和value。
    // 修改迭代器的状态。
    Last() bool
    
    IteratorWithKey
}