package containers

// Container定义所有数据结构都需要实现的基础方法
type Container interface {
    Empty() bool
    Size() int
    Clear()
    Values() []interface{}
}
