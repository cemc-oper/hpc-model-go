package hpcmodel

import (
	"strconv"
	"time"
)

type Property interface{}

type StringProperty struct {
	Value string
	Text  string
	Data  string
}

func (p *StringProperty) SetValue(value string) {
	p.Value = value
	p.Text = value
	p.Data = value
}

type NumberProperty struct {
	Value string
	Text  string
	Data  float64
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
	TimeFormat string
	Value      string
	Text       string
	Data       time.Time
}

func (p *DateTimeProperty) SetValue(value string) {
	const timeFormat = "2019-01-28T21:16:00"
	data, err := time.Parse(timeFormat, value)
	if err != nil {
		panic(err)
	}
	p.Value = value
	p.Text = value
	p.Data = data
}

type TimestampProperty struct {
	Value string
	Text  string
	Data  time.Time
}

func (p *TimestampProperty) SetValue(value string) {
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	data := time.Unix(f, 0)

	p.Value = value
	p.Text = value
	p.Data = data
}
