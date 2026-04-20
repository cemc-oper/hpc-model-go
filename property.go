package hpcmodel

import (
	"fmt"
	"strconv"
	"strings"
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
	case "TimeStringProperty":
		p = &TimeStringProperty{}
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
	p.Value = value
	p.Text = value
	data, err := strconv.ParseFloat(value, 64)
	if err != nil {
		p.Data = 0
		return
	}
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
	p.Value = value
	const timeFormat = "2006-01-02T15:04:05"
	data, err := time.Parse(timeFormat, value)
	if err != nil {
		p.Text = value
		p.Data = time.Time{}
		return
	}
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
	p.Value = value
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		p.Text = value
		p.Data = time.Time{}
		return
	}
	data := time.Unix(f, 0).UTC()
	p.Text = data.Format("2006-01-02 15:04:05")
	p.Data = data
}

func (p *TimestampProperty) SetCategory(category QueryCategory) {
	p.Category = category
}

func (p *TimestampProperty) PropertyID() string {
	return p.Category.ID
}

type TimeStringProperty struct {
	Category   QueryCategory
	TimeFormat string
	Value      string
	Text       string
	Data       time.Duration
}

func (p *TimeStringProperty) SetValue(value string) {
	p.Value = value
	p.Text = value

	if value == "N/A" || value == "" || value == "INVALID" {
		p.Data = 0
		return
	}

	var days int
	s := value
	if i := strings.IndexByte(s, '-'); i >= 0 {
		var err error
		days, err = strconv.Atoi(s[:i])
		if err != nil {
			p.Data = 0
			return
		}
		s = s[i+1:]
	}

	parts := strings.Split(s, ":")
	if len(parts) != 2 && len(parts) != 3 {
		p.Data = 0
		return
	}

	// MM:SS -> 0:MM:SS
	if len(parts) == 2 {
		parts = append([]string{"0"}, parts...)
	}

	hh, err := strconv.Atoi(parts[0])
	if err != nil {
		p.Data = 0
		return
	}
	mm, err := strconv.Atoi(parts[1])
	if err != nil {
		p.Data = 0
		return
	}
	ss, err := strconv.Atoi(parts[2])
	if err != nil {
		p.Data = 0
		return
	}

	p.Data = time.Duration(days)*24*time.Hour +
		time.Duration(hh)*time.Hour +
		time.Duration(mm)*time.Minute +
		time.Duration(ss)*time.Second
}

func (p *TimeStringProperty) SetCategory(category QueryCategory) {
	p.Category = category
}

func (p *TimeStringProperty) PropertyID() string {
	return p.Category.ID
}
