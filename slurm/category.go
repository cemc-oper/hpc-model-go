package slurm

import (
	"fmt"
	"nwpc-hpc-model-go"
	"strconv"
	"strings"
)

type QueryCategoryList struct {
	hpcmodel.QueryCategoryList
}

func (ql *QueryCategoryList) UpdateTokenIndex(titleLine string, sep string) {
	titleLine = strings.TrimSpace(titleLine)
	tokens := strings.Split(titleLine, sep)
	for index, label := range tokens {
		category := ql.CategoryFromLabel(label)
		if category != nil {
			category.RecordParserArguments = []string{strconv.Itoa(index), sep}
			recordParser, err := hpcmodel.BuildRecordParser(category)
			if err != nil {
				fmt.Errorf("build record parser failed: %v", err)
				continue
			}
			category.ParseRecord = recordParser
		}
	}
}
