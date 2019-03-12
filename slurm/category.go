package slurm

import (
	"github.com/perillaroc/hpc-model-go"
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
	for index, label := range tokens {
		category := ql.CategoryFromLabel(label)
		if category != nil {
			category.RecordParserArguments = []string{strconv.Itoa(index), sep}
			recordParser, err := hpcmodel.BuildRecordParser(category)
			if err != nil {
				log.Fatalf("build record parser failed: %v", err)
				continue
			}
			category.ParseRecord = recordParser
		}
	}
}
