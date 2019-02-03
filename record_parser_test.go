package hpcmodel_test

import (
	"nwpc-hpc-model-go"
	"testing"
)

func TestBuildRecordParser_TokenRecordParser(t *testing.T) {
	category := hpcmodel.QueryCategory{
		RecordParserClass:     "TokenRecordParser",
		RecordParserArguments: []string{"1", "|"},
	}
	_, err := hpcmodel.BuildRecordParser(category)
	if err != nil {
		t.Errorf("Build TokenRecordParser failed.")
	}
}

func TestTokenRecordParser_Parse(t *testing.T) {
	var tests = []struct {
		index   int
		sep     string
		records []string
		result  string
	}{
		{
			0,
			"|",
			[]string{"nwp_xp|(null)|1|0|NONE"},
			"nwp_xp",
		},
		{
			1,
			"|",
			[]string{"nwp_xp|(null)|1|0|NONE"},
			"(null)",
		},
		{
			2,
			"|",
			[]string{"nwp_xp|(null)|1|0|NONE"},
			"1",
		},
		{
			3,
			"|",
			[]string{"nwp_xp|(null)|1|0|NONE"},
			"0",
		},
	}

	for _, test := range tests {
		p := &hpcmodel.TokenRecordParser{
			Index: test.index,
			Sep:   test.sep,
		}
		result := p.Parse(test.records)
		if result != test.result {
			t.Errorf("result is error: %d, result is %s, expect is %s",
				test.index, result, test.result)
		}
	}
}
