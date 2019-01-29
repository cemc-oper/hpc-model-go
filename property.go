package hpcmodel

import (
	"strconv"
	"time"
)

type Property interface{}

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
