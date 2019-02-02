package hpcmodel

import (
	"fmt"
	"strconv"
	"strings"
)

type RecordParser interface {
	SetArguments(arguments []string) error
	Parse(records []string) string
}

func BuildRecordParser(category QueryCategory) RecordParser {
	var parser RecordParser
	arguments := category.RecordParserArguments
	switch category.RecordParserClass {
	case "TokenRecordParser":
		parser = &TokenRecordParser{}
		err := parser.SetArguments(arguments)
		if err != nil {
			return nil
		}
		return parser
	}
	return nil
}

type TokenRecordParser struct {
	Index int
	Sep   string
}

func (p *TokenRecordParser) SetArguments(arguments []string) error {
	l := len(arguments)
	if l == 0 || l > 2 {
		return fmt.Errorf("arguments must 1 or 2.")
	}
	var index int64
	var sep string
	var err error
	if l == 1 {
		index, err = strconv.ParseInt(arguments[0], 10, 64)
		if err != nil {
			return fmt.Errorf("parse arg[0] failed: %v", err)
		}
	} else if l == 2 {
		sep = arguments[1]
	}
	p.Index = int(index)
	p.Sep = sep
	return nil
}

func (p *TokenRecordParser) Parse(records []string) string {
	record := records[0]
	tokens := strings.Split(record, p.Sep)
	return tokens[p.Index]
}
