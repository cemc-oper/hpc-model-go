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

func splitTokens(line string, sep string) []string {
	if sep == "" || sep == " " {
		return strings.Fields(line)
	}
	return strings.Split(line, sep)
}

func (ql *QueryCategoryList) UpdateTokenIndex(titleLine string, sep string) {
	titleLine = strings.TrimSpace(titleLine)
	tokens := splitTokens(titleLine, sep)
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
