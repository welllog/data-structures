package utils

import "sort"

// 根据给定的比较器对值进行排序（原地）。
//
// 使用Go的排序（大切片的快速排序和小切片的插入排序(希尔排序)的混合体）。
func Sort(values []interface{}, comparator Comparator) {
    sort.Sort(sortable{values, comparator})
}

type sortable struct {
    values []interface{}
    comparator Comparator
}

func (s sortable) Len() int {
    return len(s.values)
}

func (s sortable) Swap(i, j int) {
    s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s sortable) Less(i, j int) bool {
    return s.comparator(s.values[i], s.values[j]) < 0
}
