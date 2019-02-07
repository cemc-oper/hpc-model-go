package slurm

import (
	"fmt"
	"github.com/perillaroc/nwpc-hpc-model-go"
)

type Model struct {
	hpcmodel.Model
}

func BuildModel(record []string, categoryList QueryCategoryList, sep string) (*Model, error) {
	if len(record) == 0 {
		return nil, nil
	}

	titleLine := record[0]
	categoryList.UpdateTokenIndex(titleLine, sep)
	lines := record[1:]
	model := new(Model)
	if len(lines) == 0 {
		return model, nil
	}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		fmt.Printf("[BuildModel]**%v**\n", line)
		item, err := hpcmodel.BuildItem([]string{line}, categoryList.QueryCategoryList)
		if err != nil {
			fmt.Printf("build item error: %v\n", err)
			continue
		}
		model.Items = append(model.Items, *item)
	}
	return model, nil
}
