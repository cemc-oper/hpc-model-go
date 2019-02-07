package hpcmodel

import "sort"

type LessFunc func(item1, item2 *Item) bool

func CreateStringPropertyLessFunc(id string) LessFunc {
	return func(item1, item2 *Item) bool {
		p1 := item1.GetProperty(id).(*StringProperty)
		p2 := item2.GetProperty(id).(*StringProperty)
		return p1.Data < p2.Data
	}
}

func CreateNumberPropertyLessFunc(id string) LessFunc {
	return func(item1, item2 *Item) bool {
		p1 := item1.GetProperty(id).(*NumberProperty)
		p2 := item2.GetProperty(id).(*NumberProperty)
		return p1.Data < p2.Data
	}
}

func CreateDateTimePropertyLessFunc(id string) LessFunc {
	return func(item1, item2 *Item) bool {
		p1 := item1.GetProperty(id).(*DateTimeProperty)
		p2 := item2.GetProperty(id).(*DateTimeProperty)
		return p1.Data.Before(p2.Data)
	}
}

type ItemSorter struct {
	items []Item
	less  []LessFunc
}

func (sorter *ItemSorter) Sort(items []Item) {
	sorter.items = items
	sort.Sort(sorter)
}

func CreateSorter(less ...LessFunc) *ItemSorter {
	return &ItemSorter{
		less: less,
	}
}

func (sorter *ItemSorter) Len() int {
	return len(sorter.items)
}

func (sorter *ItemSorter) Swap(i, j int) {
	sorter.items[i], sorter.items[j] = sorter.items[j], sorter.items[i]
}

func (sorter *ItemSorter) Less(i, j int) bool {
	p, q := &sorter.items[i], &sorter.items[j]
	var k int
	for k = 0; k < len(sorter.less)-1; k++ {
		less := sorter.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return sorter.less[k](p, q)
}
