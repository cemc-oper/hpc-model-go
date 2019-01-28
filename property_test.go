package hpcmodel_test

import (
	"nwpc-hpc-model-go"
	"testing"
)

func TestStringProperty(t *testing.T) {
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
