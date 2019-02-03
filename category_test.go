package hpcmodel_test

import (
	"nwpc-hpc-model-go"
	"testing"
)

func TestQueryCategory_Equal(t *testing.T) {
	tests := []struct {
		c1     hpcmodel.QueryCategory
		c2     hpcmodel.QueryCategory
		result bool
	}{
		{
			hpcmodel.QueryCategory{
				ID:                      "slurm.account",
				DisplayName:             "Account",
				Label:                   "ACCOUNT",
				RecordParserClass:       "TokenRecordParser",
				RecordParserArguments:   []string{"0", "|"},
				PropertyClass:           "StringProperty",
				PropertyCreateArguments: []string{},
			},
			hpcmodel.QueryCategory{
				ID:                      "slurm.account",
				DisplayName:             "Account",
				Label:                   "ACCOUNT",
				RecordParserClass:       "TokenRecordParser",
				RecordParserArguments:   []string{"0", "|"},
				PropertyClass:           "StringProperty",
				PropertyCreateArguments: []string{},
			},
			true,
		},
		{
			hpcmodel.QueryCategory{
				ID:                      "slurm.account",
				Label:                   "ACCOUNT",
				RecordParserClass:       "TokenRecordParser",
				RecordParserArguments:   []string{"0", "|"},
				PropertyClass:           "StringProperty",
				PropertyCreateArguments: []string{},
			},
			hpcmodel.QueryCategory{
				ID:                      "slurm.account",
				Label:                   "ACCOUNT",
				RecordParserClass:       "TokenRecordParser",
				RecordParserArguments:   []string{"1", "|"},
				PropertyClass:           "StringProperty",
				PropertyCreateArguments: []string{},
			},
			false,
		},
	}
	for _, test := range tests {
		if test.c1.Equal(&test.c2) != test.result {
			t.Errorf("c1 %v == c2 %v, should %t", test.c1, test.c2, test.result)
		}
	}
}

var categoryList = hpcmodel.QueryCategoryList{
	CategoryList: []hpcmodel.QueryCategory{
		{
			ID:    "slurm.account",
			Label: "ACCOUNT",
		},
		{
			ID:    "slurm.job_id",
			Label: "JOBID",
		},
		{
			ID:    "slurm.partition",
			Label: "PARTITION",
		},
		{
			ID:    "slurm.command",
			Label: "COMMAND",
		},
		{
			ID:    "slurm.state",
			Label: "STATE",
		},
		{
			ID:    "slurm.submit_time",
			Label: "SUBMIT_TIME",
		},
		{
			ID:    "slurm.work_dir",
			Label: "WORK_DIR",
		},
	},
}

func TestQueryCategoryList_ContainsID(t *testing.T) {
	tests := []struct {
		id     string
		result bool
	}{
		{"slurm.account", true},
		{"slurm.state", true},
		{"llq.user", false},
	}
	for _, test := range tests {
		if categoryList.ContainsID(test.id) != test.result {
			t.Errorf("contains %s is %t, should %t", test.id, !test.result, test.result)
		}
	}
}

func TestQueryCategoryList_IndexFromId(t *testing.T) {
	tests := []struct {
		id     string
		result int
	}{
		{"slurm.account", 0},
		{"slurm.state", 4},
		{"llq.user", -1},
	}
	for _, test := range tests {
		index := categoryList.IndexFromId(test.id)
		if index != test.result {
			t.Errorf("index for %s is %d, should %d", test.id, index, test.result)
		}
	}
}

func TestQueryCategoryList_CategoryFromId(t *testing.T) {
	tests := []struct {
		id     string
		result *hpcmodel.QueryCategory
	}{
		{"slurm.account", &categoryList.CategoryList[0]},
		{"slurm.state", &categoryList.CategoryList[4]},
		{"llq.user", nil},
	}
	for _, test := range tests {
		category := categoryList.CategoryFromId(test.id)
		if !category.Equal(test.result) {
			t.Errorf("category for %s is %v, should %v",
				test.id, category, test.result)
		}
	}
}

func TestQueryCategoryList_ContainsLabel(t *testing.T) {
	tests := []struct {
		label  string
		result bool
	}{
		{"ACCOUNT", true},
		{"STATE", true},
		{"USER", false},
	}
	for _, test := range tests {
		if categoryList.ContainsLabel(test.label) != test.result {
			t.Errorf("contains %s is %t, should %t", test.label, !test.result, test.result)
		}
	}
}

func TestQueryCategoryList_IndexFromLabel(t *testing.T) {
	tests := []struct {
		label  string
		result int
	}{
		{"ACCOUNT", 0},
		{"STATE", 4},
		{"USER", -1},
	}
	for _, test := range tests {
		index := categoryList.IndexFromLabel(test.label)
		if index != test.result {
			t.Errorf("index for %s is %d, should %d", test.label, index, test.result)
		}
	}
}

func TestQueryCategoryList_CategoryFromLabel(t *testing.T) {
	tests := []struct {
		label  string
		result *hpcmodel.QueryCategory
	}{
		{"ACCOUNT", &categoryList.CategoryList[0]},
		{"STATE", &categoryList.CategoryList[4]},
		{"USER", nil},
	}
	for _, test := range tests {
		category := categoryList.CategoryFromLabel(test.label)
		if !category.Equal(test.result) {
			t.Errorf("category for %s is %v, should %v",
				test.label, category, test.result)
		}
	}
}
