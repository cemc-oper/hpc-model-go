package hpcmodel

import (
	"fmt"
	"strconv"
	"time"
)

type Property interface {
	SetValue(value string)
	SetCategory(category QueryCategory)
	PropertyID() string
}

func BuildProperty(records []string, category QueryCategory) (Property, error) {
	var p Property
	switch category.PropertyClass {
	case "StringProperty":
		p = &StringProperty{}
	case "NumberProperty":
		p = &NumberProperty{}
	case "DateTimeProperty":
		p = &DateTimeProperty{}
	case "TimestampProperty":
		p = &TimestampProperty{}
	default:
		return nil, fmt.Errorf("error PropertyClass: %s", category.PropertyClass)
	}
	p.SetCategory(category)
	value := category.ParseRecord.Parse(records)
	p.SetValue(value)
	return p, nil
}

type StringProperty struct {
	Category QueryCategory
	Value    string
	Text     string
	Data     string
}

func (p *StringProperty) SetValue(value string) {
	p.Value = value
	p.Text = value
	p.Data = value
}

func (p *StringProperty) SetCategory(category QueryCategory) {
	p.Category = category
}

func (p *StringProperty) PropertyID() string {
	return p.Category.ID
}

type NumberProperty struct {
	Category QueryCategory
	Value    string
	Text     string
	Data     float64
}

func (p *NumberProperty) SetValue(value string) {
	data, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	p.Value = value
	p.Text = value
	p.Data = data
}

func (p *NumberProperty) SetCategory(category QueryCategory) {
	p.Category = category
}

func (p *NumberProperty) PropertyID() string {
	return p.Category.ID
}

type DateTimeProperty struct {
	Category   QueryCategory
	TimeFormat string
	Value      string
	Text       string
	Data       time.Time
}

func (p *DateTimeProperty) SetValue(value string) {
	const timeFormat = "2006-01-02T15:04:05"
	data, err := time.Parse(timeFormat, value)
	if err != nil {
		panic(err)
	}
	p.Value = value
	p.Text = data.Format("2006-01-02 15:04:05")
	p.Data = data
}

func (p *DateTimeProperty) SetCategory(category QueryCategory) {
	p.Category = category
}

func (p *DateTimeProperty) PropertyID() string {
	return p.Category.ID
}

type TimestampProperty struct {
	Category QueryCategory
	Value    string
	Text     string
	Data     time.Time
}

func (p *TimestampProperty) SetValue(value string) {
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	data := time.Unix(f, 0).UTC()

	p.Value = value
	p.Text = data.Format("2006-01-02 15:04:05")
	p.Data = data
}

func (p *TimestampProperty) SetCategory(category QueryCategory) {
	p.Category = category
}

func (p *TimestampProperty) PropertyID() string {
	return p.Category.ID
}
