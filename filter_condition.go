package hpcmodel

import "strings"

type FilterCondition interface {
	isFit(item *Item) bool
}

type StringPropertyFilterCondition struct {
	ID      string
	Checker StringValueChecker
}

func (f *StringPropertyFilterCondition) isFit(item *Item) bool {
	prop := item.GetProperty(f.ID)
	if prop == nil {
		return false
	}
	sp, ok := prop.(*StringProperty)
	if !ok {
		return false
	}
	return f.Checker.CheckValue(sp.Data)
}

type StringValueChecker interface {
	CheckValue(s string) bool
}

type StringEqualValueChecker struct {
	ExpectedValue string
}

func (c *StringEqualValueChecker) CheckValue(s string) bool {
	if s == c.ExpectedValue {
		return true
	} else {
		return false
	}
}

type StringInValueChecker struct {
	ExpectedValues []string
}

func (c *StringInValueChecker) CheckValue(s string) bool {
	for _, v := range c.ExpectedValues {
		if s == v {
			return true
		}
	}
	return false
}

type StringContainChecker struct {
	ExpectedValue string
}

func (c *StringContainChecker) CheckValue(s string) bool {
	return strings.Contains(s, c.ExpectedValue)
}
