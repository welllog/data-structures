package containers

// Enumerable提供了Ruby启发的（each、select、map、find、any？等）容器函数。

// EnumerableWithIndex提供了有序容器的函数，这些容器的值可以通过索引来获取。
type EnumerableWithIndex interface {
    // Each对每个元素调用一次给定的函数，传递该元素的索引和值。
    Each(func(index int, value interface{}))
    
    // Map对每个元素调用一次给定的函数，并返回一个
    // 容器中包含指定函数返回的值。
    // TODO希望能得到帮助，如何在容器中实现这一点(不想在链式时使用断言类型)
    // Map(func(index int, value interface{}) interface{}) Container
    
    // Select返回一个新的容器，其中包含所有元素，而给定的函数对这些元素返回一个真值。
    // TODO需要帮助，如何在容器中强制执行这一点（不想在链式时使用断言的类型
    // Select(func(index int, value interface{}) bool) Container
    
    // Any 将容器的每个元素传递给给定的函数
    // 如果函数对任何一个元素都返回true，则返回true。
    Any(func(index int, value interface{}) bool) bool
    
    // All 将容器的每个元素传递给给定的函数
    // 如果函数对所有元素都返回true，则返回true。
    All(func(index int, value interface{}) bool) bool
    
    // Find将容器中的每个元素传递给给定函数，并返回
    // 函数为真的第一个(index,value)，否则为-1,nil。
    // 如果没有符合标准的元素。
    Find(func(index int, value interface{}) bool) (int, interface{})
}

// EnumerableWithKey提供了有序容器的功能，其值的元素是键/值对。
type EnumerableWithKey interface {
    // Each对每个元素调用一次给定的函数，并传递该元素的键和值。
    Each(func(key interface{}, value interface{}))
    
    // Map对每个元素调用一次给定的函数，并返回一个容器。
    // 包含由给定函数返回的键/值对的值。
    // TODO需要帮助，如何在容器中实现这一点（不希望在链式连接的时候使用断言类型）。
    // Map(func(key interface{}, value interface{}) (interface{}, interface{})) Container
    
    // Select返回一个新的容器，其中包含所有元素，而给定的函数对这些元素返回一个真值。
    // TODO需要帮助，如何在容器中强制执行这一点（不想在链式时使用断言的类型
    // Select(func(key interface{}, value interface{}) bool) Container

    // Any将容器中的每个元素传递给给定的函数
    // 如果函数对任何一个元素返回true，则返回true。
    Any(func(key interface{}, value interface{}) bool) bool

    // All将容器中的每个元素传递给给定的函数
    // 如果函数对所有元素都返回true，则返回true。
    All(func(key interface{}, value interface{}) bool) bool
    
    // Find将容器中的每个元素传递给给定函数，并返回
    // 函数为真的第一个(key,value)，如果没有元素符合标准，则为nil,nil。
    Find(func(key interface{}, value interface{}) bool) (interface{}, interface{})
}