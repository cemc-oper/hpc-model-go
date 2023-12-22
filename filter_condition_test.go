package hpcmodel_test

import (
	hpcmodel "github.com/cemc-oper/hpc-model-go"
	"testing"
	"time"
)

func TestStringPropertyFilterCondition_IsFit(t *testing.T) {
	item := hpcmodel.Item{
		Props: []hpcmodel.Property{
			&hpcmodel.StringProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.account",
				},
				Data: "nwp_xp",
			},
			&hpcmodel.StringProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.partition",
				},
				Data: "serial",
			},
			&hpcmodel.StringProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.state",
				},
				Data: "RUNNING",
			},
		},
	}

	tests := []struct {
		condition *hpcmodel.StringPropertyFilterCondition
		result    bool
	}{
		{
			condition: &hpcmodel.StringPropertyFilterCondition{
				ID: "squeue.account",
				Checker: &hpcmodel.StringEqualValueChecker{
					ExpectedValue: "nwp_xp",
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.StringPropertyFilterCondition{
				ID: "squeue.partition",
				Checker: &hpcmodel.StringEqualValueChecker{
					ExpectedValue: "serial",
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.StringPropertyFilterCondition{
				ID: "squeue.partition",
				Checker: &hpcmodel.StringEqualValueChecker{
					ExpectedValue: "serial_op",
				},
			},
			result: false,
		},
		{
			condition: &hpcmodel.StringPropertyFilterCondition{
				ID: "squeue.account",
				Checker: &hpcmodel.StringInValueChecker{
					ExpectedValues: []string{"nwp", "nwp_xp"},
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.StringPropertyFilterCondition{
				ID: "squeue.account",
				Checker: &hpcmodel.StringContainChecker{
					ExpectedValue: "nwp",
				},
			},
			result: true,
		},
	}

	for _, test := range tests {
		if test.condition.IsFit(&item) != test.result {
			t.Errorf("condtion failed: %v", test.condition)
		}
	}
}

func TestNumberPropertyFilterCondition_IsFit(t *testing.T) {
	item := hpcmodel.Item{
		Props: []hpcmodel.Property{
			&hpcmodel.NumberProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.cpus",
				},
				Data: 18,
			},
			&hpcmodel.NumberProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.nodes",
				},
				Data: 32,
			},
		},
	}

	tests := []struct {
		condition *hpcmodel.NumberPropertyFilterCondition
		result    bool
	}{
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberEqualValueChecker{
					ExpectedValue: 18,
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberEqualValueChecker{
					ExpectedValue: 20,
				},
			},
			result: false,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberGreaterValueChecker{
					ExpectedValue: 15,
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberGreaterValueChecker{
					ExpectedValue: 30,
				},
			},
			result: false,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberLessValueChecker{
					ExpectedValue: 30,
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberLessValueChecker{
					ExpectedValue: 10,
				},
			},
			result: false,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberInValueChecker{
					ExpectedValues: []float64{18, 20},
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &hpcmodel.NumberInValueChecker{
					ExpectedValues: []float64{19, 20},
				},
			},
			result: false,
		},
	}

	for _, test := range tests {
		if test.condition.IsFit(&item) != test.result {
			t.Errorf("condtion failed: %v", test.condition)
		}
	}
}

func TestDateTimePropertyFilterCondition_IsFit(t *testing.T) {
	item := hpcmodel.Item{
		Props: []hpcmodel.Property{
			&hpcmodel.DateTimeProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.submit_time",
				},
				Data: time.Date(2019, time.February, 6, 18, 00, 00, 00, time.UTC),
			},
			&hpcmodel.DateTimeProperty{
				Category: hpcmodel.QueryCategory{
					ID: "squeue.start_time",
				},
				Data: time.Date(2019, time.February, 6, 19, 00, 00, 00, time.UTC),
			},
		},
	}

	tests := []struct {
		condition *hpcmodel.DateTimePropertyFilterCondition
		result    bool
	}{
		{
			condition: &hpcmodel.DateTimePropertyFilterCondition{
				ID: "squeue.submit_time",
				Checker: &hpcmodel.DateTimeEqualValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 18, 00, 00, 00, time.UTC),
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.DateTimePropertyFilterCondition{
				ID: "squeue.submit_time",
				Checker: &hpcmodel.DateTimeEqualValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 19, 00, 00, 00, time.UTC),
				},
			},
			result: false,
		},
		{
			condition: &hpcmodel.DateTimePropertyFilterCondition{
				ID: "squeue.start_time",
				Checker: &hpcmodel.DateTimeAfterValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 18, 30, 00, 00, time.UTC),
				},
			},
			result: true,
		},
		{
			condition: &hpcmodel.DateTimePropertyFilterCondition{
				ID: "squeue.start_time",
				Checker: &hpcmodel.DateTimeBeforeValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 18, 30, 00, 00, time.UTC),
				},
			},
			result: false,
		},
	}

	for _, test := range tests {
		if test.condition.IsFit(&item) != test.result {
			t.Errorf("condtion failed: %v", test.condition)
		}
	}
}
