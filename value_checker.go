package hpcmodel

import (
	"fmt"
	"strings"
	"time"
)

type StringValueChecker interface {
	CheckValue(s string) bool
}

type StringEqualValueChecker struct {
	ExpectedValue string
}

func (c *StringEqualValueChecker) CheckValue(s string) bool {
	return s == c.ExpectedValue
}

type StringInValueChecker struct {
	ExpectedValues []string
}

func (c *StringInValueChecker) CheckValue(s string) bool {
	for _, v := range c.ExpectedValues {
		if s == v {
			return true
		}
		fmt.Printf("%s != %s: %d\n", s, v, strings.EqualFold(s, v))
	}
	fmt.Printf("StringInValueChecker false: %s not in %v\n", s, c.ExpectedValues)
	return false
}

type StringContainChecker struct {
	ExpectedValue string
}

func (c *StringContainChecker) CheckValue(s string) bool {
	return strings.Contains(s, c.ExpectedValue)
}

type NumberValueChecker interface {
	CheckValue(s float64) bool
}

type NumberEqualValueChecker struct {
	ExpectedValue float64
}

func (c *NumberEqualValueChecker) CheckValue(s float64) bool {
	return s == c.ExpectedValue
}

type NumberGreaterValueChecker struct {
	ExpectedValue float64
}

func (c *NumberGreaterValueChecker) CheckValue(s float64) bool {
	return s > c.ExpectedValue
}

type NumberLessValueChecker struct {
	ExpectedValue float64
}

func (c *NumberLessValueChecker) CheckValue(s float64) bool {
	return s < c.ExpectedValue
}

type NumberInValueChecker struct {
	ExpectedValues []float64
}

func (c *NumberInValueChecker) CheckValue(s float64) bool {
	for _, v := range c.ExpectedValues {
		if s == v {
			return true
		}
	}
	return false
}

type DateTimeValueChecker interface {
	CheckValue(t time.Time) bool
}

type DateTimeEqualValueChecker struct {
	ExpectedValue time.Time
}

func (c *DateTimeEqualValueChecker) CheckValue(t time.Time) bool {
	return t.Equal(c.ExpectedValue)
}

type DateTimeAfterValueChecker struct {
	ExpectedValue time.Time
}

func (c *DateTimeAfterValueChecker) CheckValue(t time.Time) bool {
	return t.After(c.ExpectedValue)
}

type DateTimeBeforeValueChecker struct {
	ExpectedValue time.Time
}

func (c *DateTimeBeforeValueChecker) CheckValue(t time.Time) bool {
	return t.Before(c.ExpectedValue)
}

type DateTimeInValueChecker struct {
	ExpectedValues []time.Time
}

func (c *DateTimeInValueChecker) CheckValue(t time.Time) bool {
	for _, v := range c.ExpectedValues {
		if t.Equal(v) {
			return true
		}
	}
	return false
}
