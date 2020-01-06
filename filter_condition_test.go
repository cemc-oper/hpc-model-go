package hpcmodel_test

import (
	. "github.com/nwpc-oper/hpc-model-go"
	"testing"
	"time"
)

func TestStringPropertyFilterCondition_IsFit(t *testing.T) {
	item := Item{
		Props: []Property{
			&StringProperty{
				Category: QueryCategory{
					ID: "squeue.account",
				},
				Data: "nwp_xp",
			},
			&StringProperty{
				Category: QueryCategory{
					ID: "squeue.partition",
				},
				Data: "serial",
			},
			&StringProperty{
				Category: QueryCategory{
					ID: "squeue.state",
				},
				Data: "RUNNING",
			},
		},
	}

	tests := []struct {
		condition *StringPropertyFilterCondition
		result    bool
	}{
		{
			condition: &StringPropertyFilterCondition{
				ID: "squeue.account",
				Checker: &StringEqualValueChecker{
					ExpectedValue: "nwp_xp",
				},
			},
			result: true,
		},
		{
			condition: &StringPropertyFilterCondition{
				ID: "squeue.partition",
				Checker: &StringEqualValueChecker{
					ExpectedValue: "serial",
				},
			},
			result: true,
		},
		{
			condition: &StringPropertyFilterCondition{
				ID: "squeue.partition",
				Checker: &StringEqualValueChecker{
					ExpectedValue: "serial_op",
				},
			},
			result: false,
		},
		{
			condition: &StringPropertyFilterCondition{
				ID: "squeue.account",
				Checker: &StringInValueChecker{
					ExpectedValues: []string{"nwp", "nwp_xp"},
				},
			},
			result: true,
		},
		{
			condition: &StringPropertyFilterCondition{
				ID: "squeue.account",
				Checker: &StringContainChecker{
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
	item := Item{
		Props: []Property{
			&NumberProperty{
				Category: QueryCategory{
					ID: "squeue.cpus",
				},
				Data: 18,
			},
			&NumberProperty{
				Category: QueryCategory{
					ID: "squeue.nodes",
				},
				Data: 32,
			},
		},
	}

	tests := []struct {
		condition *NumberPropertyFilterCondition
		result    bool
	}{
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberEqualValueChecker{
					ExpectedValue: 18,
				},
			},
			result: true,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberEqualValueChecker{
					ExpectedValue: 20,
				},
			},
			result: false,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberGreaterValueChecker{
					ExpectedValue: 15,
				},
			},
			result: true,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberGreaterValueChecker{
					ExpectedValue: 30,
				},
			},
			result: false,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberLessValueChecker{
					ExpectedValue: 30,
				},
			},
			result: true,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberLessValueChecker{
					ExpectedValue: 10,
				},
			},
			result: false,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberInValueChecker{
					ExpectedValues: []float64{18, 20},
				},
			},
			result: true,
		},
		{
			condition: &NumberPropertyFilterCondition{
				ID: "squeue.cpus",
				Checker: &NumberInValueChecker{
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
	item := Item{
		Props: []Property{
			&DateTimeProperty{
				Category: QueryCategory{
					ID: "squeue.submit_time",
				},
				Data: time.Date(2019, time.February, 6, 18, 00, 00, 00, time.UTC),
			},
			&DateTimeProperty{
				Category: QueryCategory{
					ID: "squeue.start_time",
				},
				Data: time.Date(2019, time.February, 6, 19, 00, 00, 00, time.UTC),
			},
		},
	}

	tests := []struct {
		condition *DateTimePropertyFilterCondition
		result    bool
	}{
		{
			condition: &DateTimePropertyFilterCondition{
				ID: "squeue.submit_time",
				Checker: &DateTimeEqualValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 18, 00, 00, 00, time.UTC),
				},
			},
			result: true,
		},
		{
			condition: &DateTimePropertyFilterCondition{
				ID: "squeue.submit_time",
				Checker: &DateTimeEqualValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 19, 00, 00, 00, time.UTC),
				},
			},
			result: false,
		},
		{
			condition: &DateTimePropertyFilterCondition{
				ID: "squeue.start_time",
				Checker: &DateTimeAfterValueChecker{
					ExpectedValue: time.Date(2019, time.February, 6, 18, 30, 00, 00, time.UTC),
				},
			},
			result: true,
		},
		{
			condition: &DateTimePropertyFilterCondition{
				ID: "squeue.start_time",
				Checker: &DateTimeBeforeValueChecker{
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
