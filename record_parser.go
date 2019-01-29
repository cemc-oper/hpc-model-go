package hpcmodel

import "strings"

type RecordParser interface {
	Parse(records []string) string
}

type TokenRecordParser struct {
	Index int
	Sep   string
}

func (p *TokenRecordParser) Parse(records []string) string {
	record := records[0]
	tokens := strings.Split(record, p.Sep)
	return tokens[p.Index]
}
