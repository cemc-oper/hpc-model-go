package slurm

import (
	hpcmodel "github.com/cemc-oper/hpc-model-go"
	"log"
	"strconv"
	"strings"
)

type QueryCategoryList struct {
	hpcmodel.QueryCategoryList
}

func (ql *QueryCategoryList) UpdateTokenIndex(titleLine string, sep string) {
	titleLine = strings.TrimSpace(titleLine)
	var tokens []string
	if sep == "" || sep == " " {
		tokens = strings.Fields(titleLine)
	} else {
		tokens = strings.Split(titleLine, sep)
	}
	labelOccurrences := make(map[string]int)
	for index, label := range tokens {
		occurrence := labelOccurrences[label]
		labelOccurrences[label] = occurrence + 1
		category := ql.CategoryFromLabel(label)
		if category == nil || occurrence != category.LabelOccurrence {
			continue
		}
		category.RecordParserArguments = []string{strconv.Itoa(index), sep}
		recordParser, err := hpcmodel.BuildRecordParser(category)
		if err != nil {
			log.Fatalf("build record parser failed: %v", err)
		}
		category.ParseRecord = recordParser
	}
}
