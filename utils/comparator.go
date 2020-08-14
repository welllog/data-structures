package utils

import "time"

// 比较器将进行类型断言（参见IntComparator的例子）。
// 如果a或b不属于断言类型，它将会panic。
//
// 应该返回一个数字。
// 负数, 则 a < b
// 零, 则 a == b
// 正数, 则 a > b
type Comparator func(a, b interface{}) int

func StringComparator(a, b interface{}) int {
    s1 := a.(string)
    s2 := b.(string)
    min := len(s2)
    if len(s1) < len(s2) {
        min = len(s1)
    }
    diff := 0
    for i := 0; i < min && diff == 0; i++ {
        diff = int(s1[i]) - int(s2[i])
    }
    if diff == 0 {
        diff = len(s1) - len(s2)
    }
    if diff < 0 {
        return -1
    }
    if diff > 0 {
        return 1
    }
    return 0
}

func IntComparator(a, b interface{}) int {
    i1 := a.(int)
    i2 := b.(int)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func Int8Comparator(a, b interface{}) int {
    i1 := a.(int8)
    i2 := b.(int8)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func Int16Comparator(a, b interface{}) int {
    i1 := a.(int16)
    i2 := b.(int16)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func Int32Comparator(a, b interface{}) int {
    i1 := a.(int32)
    i2 := b.(int32)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func Int64Comparator(a, b interface{}) int {
    i1 := a.(int64)
    i2 := b.(int64)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func UIntComparator(a, b interface{}) int {
    i1 := a.(uint)
    i2 := b.(uint)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func UInt8Comparator(a, b interface{}) int {
    i1 := a.(uint8)
    i2 := b.(uint8)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func UInt16Comparator(a, b interface{}) int {
    i1 := a.(uint16)
    i2 := b.(uint16)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func UInt32Comparator(a, b interface{}) int {
    i1 := a.(uint32)
    i2 := b.(uint32)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func UInt64Comparator(a, b interface{}) int {
    i1 := a.(uint64)
    i2 := b.(uint64)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func Float32Comparator(a, b interface{}) int {
    i1 := a.(float32)
    i2 := b.(float32)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func Float64Comparator(a, b interface{}) int {
    i1 := a.(float64)
    i2 := b.(float64)
    switch {
    case i1 > i2:
        return 1
    case i1 < i2:
        return -1
    default:
        return 0
    }
}

func ByteComparator(a, b interface{}) int {
    aAsserted := a.(byte)
    bAsserted := b.(byte)
    switch {
    case aAsserted > bAsserted:
        return 1
    case aAsserted < bAsserted:
        return -1
    default:
        return 0
    }
}

func RuneComparator(a, b interface{}) int {
    aAsserted := a.(rune)
    bAsserted := b.(rune)
    switch {
    case aAsserted > bAsserted:
        return 1
    case aAsserted < bAsserted:
        return -1
    default:
        return 0
    }
}

func TimeComparator(a, b interface{}) int {
    aAsserted := a.(time.Time)
    bAsserted := b.(time.Time)
    
    switch {
    case aAsserted.After(bAsserted):
        return 1
    case aAsserted.Before(bAsserted):
        return -1
    default:
        return 0
    }
}
