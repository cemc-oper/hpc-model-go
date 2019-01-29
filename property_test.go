package hpcmodel_test

import (
	"nwpc-hpc-model-go"
	"testing"
	"time"
)

func TestStringProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  string
		text  string
	}{
		{"nwpc", "nwpc", "nwpc", "nwpc"},
	}

	for _, test := range tests {
		var p hpcmodel.StringProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestNumberProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  string
		text  string
	}{
		{"nwpc", "nwpc", "nwpc", "nwpc"},
	}

	for _, test := range tests {
		var p hpcmodel.StringProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestDateTimeProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  time.Time
		text  string
	}{
		{
			"2019-01-29T19:45:00",
			"2019-01-29T19:45:00",
			time.Date(2019, time.January, 29, 19, 45, 0, 0, time.UTC),
			"2019-01-29 19:45:00",
		},
	}

	for _, test := range tests {
		var p hpcmodel.DateTimeProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestTimestampProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  time.Time
		text  string
	}{
		{
			"1547906480",
			"1547906480",
			time.Date(2019, time.January, 19, 14, 01, 20, 0, time.UTC),
			"2019-01-19 14:01:20",
		},
	}

	for _, test := range tests {
		var p hpcmodel.TimestampProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}
