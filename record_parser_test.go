package hpcmodel_test

import (
	"github.com/perillaroc/nwpc-hpc-model-go"
	"testing"
)

func TestBuildRecordParser_TokenRecordParser(t *testing.T) {
	category := &hpcmodel.QueryCategory{
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

func TestTokenRecordParser_SetArguments(t *testing.T) {
	var tests = []struct {
		arguments []string
		index     int
		sep       string
	}{
		{
			[]string{"1", "|"},
			1,
			"|",
		},
		{
			[]string{"2"},
			2,
			"",
		},
	}

	for _, test := range tests {
		parser := new(hpcmodel.TokenRecordParser)
		err := parser.SetArguments(test.arguments)
		if err != nil {
			t.Errorf("set arguments %v failed: %v", test.arguments, err)
		}
		if parser.Index != test.index {
			t.Errorf("set arguments %v Index is %d, requred %d", test.arguments, parser.Index, test.index)
		}
		if parser.Sep != test.sep {
			t.Errorf("set arguments %v Sep is %s, requred %s", test.arguments, parser.Sep, test.sep)
		}
	}
}
