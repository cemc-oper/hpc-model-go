package slurm

import (
	"fmt"
	"strings"

	hpcmodel "github.com/cemc-oper/hpc-model-go"
)

type Model struct {
	hpcmodel.Model
}

func BuildModel(record []string, categoryList QueryCategoryList, sep string) (*Model, error) {
	model := new(Model)
	headerIndex := -1
	for i, line := range record {
		if strings.TrimSpace(line) != "" {
			headerIndex = i
			break
		}
	}
	if headerIndex == -1 {
		return model, nil
	}

	titleLine := strings.TrimSpace(record[headerIndex])
	headerTokens := splitTokens(titleLine, sep)
	if err := validateHeader(headerTokens, categoryList); err != nil {
		return nil, fmt.Errorf("line %d: %w", headerIndex+1, err)
	}
	categoryList.UpdateTokenIndex(titleLine, sep)

	for i := headerIndex + 1; i < len(record); i++ {
		line := record[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		tokens := splitTokens(line, sep)
		if len(tokens) != len(headerTokens) {
			return nil, fmt.Errorf("line %d: got %d columns, want %d", i+1, len(tokens), len(headerTokens))
		}
		item, err := hpcmodel.BuildItem([]string{line}, categoryList.QueryCategoryList)
		if err != nil {
			return nil, fmt.Errorf("line %d: build item: %w", i+1, err)
		}
		model.Items = append(model.Items, *item)
	}
	return model, nil
}

func validateHeader(headerTokens []string, categoryList QueryCategoryList) error {
	labelCounts := make(map[string]int)
	for _, label := range headerTokens {
		labelCounts[label]++
	}

	var missing []string
	for _, category := range categoryList.CategoryList {
		if category.LabelOccurrence < 0 {
			return fmt.Errorf("category %q has negative label occurrence %d", category.ID, category.LabelOccurrence)
		}
		if labelCounts[category.Label] <= category.LabelOccurrence {
			missing = append(missing, category.Label)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required columns: %s", strings.Join(missing, ", "))
	}
	return nil
}
