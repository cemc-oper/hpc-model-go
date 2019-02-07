package hpcmodel

import "fmt"

type FilterCondition interface {
	IsFit(item *Item) bool
}

type StringPropertyFilterCondition struct {
	ID      string
	Checker StringValueChecker
}

func (f *StringPropertyFilterCondition) IsFit(item *Item) bool {
	prop := item.GetProperty(f.ID)
	if prop == nil {
		fmt.Printf("property not found: %s\n", f.ID)
		return false
	}
	sp, ok := prop.(*StringProperty)
	if !ok {
		fmt.Printf("property is not string: %s\n", f.ID)
		return false
	}
	return f.Checker.CheckValue(sp.Data)
}

// Number Property

type NumberPropertyFilterCondition struct {
	ID      string
	Checker NumberValueChecker
}

func (f *NumberPropertyFilterCondition) IsFit(item *Item) bool {
	prop := item.GetProperty(f.ID)
	if prop == nil {
		return false
	}
	sp, ok := prop.(*NumberProperty)
	if !ok {
		return false
	}
	return f.Checker.CheckValue(sp.Data)
}

// DateTime property

type DateTimePropertyFilterCondition struct {
	ID      string
	Checker DateTimeValueChecker
}

func (f *DateTimePropertyFilterCondition) IsFit(item *Item) bool {
	prop := item.GetProperty(f.ID)
	if prop == nil {
		return false
	}
	dp, ok := prop.(*DateTimeProperty)
	if !ok {
		return false
	}
	return f.Checker.CheckValue(dp.Data)
}
