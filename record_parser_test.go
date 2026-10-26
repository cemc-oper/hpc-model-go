package hpcmodel_test

import (
	"testing"

	hpcmodel "github.com/cemc-oper/hpc-model-go"
)

func TestBuildRecordParser_TokenRecordParser(t *testing.T) {
	category := &hpcmodel.QueryCategory{
		RecordParserClass:     "TokenRecordParser",
		RecordParserArguments: []string{"1", "|"},
	}

	recordParser, err := hpcmodel.BuildRecordParser(category)
	if err != nil {
		t.Fatalf("BuildRecordParser() error = %v", err)
	}
	p, ok := recordParser.(*hpcmodel.TokenRecordParser)
	if !ok {
		t.Fatalf("BuildRecordParser() type = %T, want *TokenRecordParser", recordParser)
	}
	if p.Index != 1 || p.Sep != "|" {
		t.Errorf("parser = {Index: %d, Sep: %q}, want {Index: 1, Sep: %q}", p.Index, p.Sep, "|")
	}
}

func TestTokenRecordParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		sep     string
		records []string
		want    string
	}{
		{
			name:    "first column",
			index:   0,
			sep:     "|",
			records: []string{"nwp_xp|(null)|1|0|NONE"},
			want:    "nwp_xp",
		},
		{
			name:    "middle column",
			index:   2,
			sep:     "|",
			records: []string{"nwp_xp|(null)|1|0|NONE"},
			want:    "1",
		},
		{
			name:    "last column",
			index:   4,
			sep:     "|",
			records: []string{"nwp_xp|(null)|1|0|NONE"},
			want:    "NONE",
		},
		{
			name:    "empty column between consecutive separators",
			index:   1,
			sep:     "|",
			records: []string{"alpha||gamma"},
			want:    "",
		},
		{
			name:    "trailing empty column",
			index:   2,
			sep:     "|",
			records: []string{"alpha|beta|"},
			want:    "",
		},
		{
			name:    "space separator uses fields",
			index:   1,
			sep:     " ",
			records: []string{"  alpha   beta gamma  "},
			want:    "beta",
		},
		{
			name:    "empty separator uses fields",
			index:   2,
			sep:     "",
			records: []string{"  alpha   beta gamma  "},
			want:    "gamma",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hpcmodel.TokenRecordParser{Index: tt.index, Sep: tt.sep}
			if got := p.Parse(tt.records); got != tt.want {
				t.Errorf("Parse(%q) with index %d and separator %q = %q, want %q", tt.records, tt.index, tt.sep, got, tt.want)
			}
		})
	}
}

func TestTokenRecordParser_SetArguments(t *testing.T) {
	tests := []struct {
		name      string
		arguments []string
		wantIndex int
		wantSep   string
	}{
		{name: "index and separator", arguments: []string{"1", "|"}, wantIndex: 1, wantSep: "|"},
		{name: "index only", arguments: []string{"2"}, wantIndex: 2, wantSep: ""},
		{name: "zero index", arguments: []string{"0", " "}, wantIndex: 0, wantSep: " "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := new(hpcmodel.TokenRecordParser)
			if err := parser.SetArguments(tt.arguments); err != nil {
				t.Fatalf("SetArguments(%v) error = %v", tt.arguments, err)
			}
			if parser.Index != tt.wantIndex {
				t.Errorf("Index = %d, want %d", parser.Index, tt.wantIndex)
			}
			if parser.Sep != tt.wantSep {
				t.Errorf("Sep = %q, want %q", parser.Sep, tt.wantSep)
			}
		})
	}
}

func TestTokenRecordParser_SetArgumentsErrors(t *testing.T) {
	tests := []struct {
		name      string
		arguments []string
	}{
		{name: "no arguments", arguments: nil},
		{name: "too many arguments", arguments: []string{"1", "|", "extra"}},
		{name: "invalid index", arguments: []string{"invalid", "|"}},
		{name: "negative index", arguments: []string{"-1", "|"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := new(hpcmodel.TokenRecordParser)
			if err := parser.SetArguments(tt.arguments); err == nil {
				t.Errorf("SetArguments(%v) error = nil, want an error", tt.arguments)
			}
		})
	}
}
