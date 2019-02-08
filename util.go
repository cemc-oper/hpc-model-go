package hpcmodel

import (
	"fmt"
	"sort"
)

type LessFunc func(item1, item2 *Item) bool

func CreatePropertyLessFunc(id string) LessFunc {
	return func(item1, item2 *Item) bool {
		p1 := item1.GetProperty(id)
		p2 := item2.GetProperty(id)
		switch prop1 := p1.(type) {
		case *StringProperty:
			prop2 := p2.(*StringProperty)
			return prop1.Data < prop2.Data
		case *NumberProperty:
			prop2 := p2.(*NumberProperty)
			return prop1.Data < prop2.Data
		case *DateTimeProperty:
			prop2 := p2.(*DateTimeProperty)
			return prop1.Data.Before(prop2.Data)
		case *TimestampProperty:
			prop2 := p2.(*DateTimeProperty)
			return prop1.Data.Before(prop2.Data)
		case nil:
			panic(fmt.Errorf("prop %s not found", id))
		default:
			panic(fmt.Errorf("prop %s type not supported", id))
		}
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
