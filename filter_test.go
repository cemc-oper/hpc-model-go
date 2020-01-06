package hpcmodel_test

import (
	. "github.com/nwpc-oper/hpc-model-go"
	"testing"
	"time"
)

func TestFilter_Filter(t *testing.T) {
	items := []Item{
		{
			Props: []Property{
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.jobid",
					},
					Data: "1",
				},
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
				&DateTimeProperty{
					Category: QueryCategory{
						ID: "squeue.submit_time",
					},
					Data: time.Date(2019, time.February, 6, 18, 00, 00, 00, time.UTC),
				},
			},
		},
		{
			Props: []Property{
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.jobid",
					},
					Data: "2",
				},
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.account",
					},
					Data: "nwp",
				},
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.partition",
					},
					Data: "serial_op",
				},
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.state",
					},
					Data: "RUNNING",
				},
				&DateTimeProperty{
					Category: QueryCategory{
						ID: "squeue.submit_time",
					},
					Data: time.Date(2019, time.February, 6, 20, 00, 00, 00, time.UTC),
				},
			},
		},
		{
			Props: []Property{
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.jobid",
					},
					Data: "3",
				},
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.account",
					},
					Data: "windroc",
				},
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.partition",
					},
					Data: "serial_op",
				},
				&StringProperty{
					Category: QueryCategory{
						ID: "squeue.state",
					},
					Data: "PENDING",
				},
				&DateTimeProperty{
					Category: QueryCategory{
						ID: "squeue.submit_time",
					},
					Data: time.Date(2019, time.February, 6, 19, 00, 00, 00, time.UTC),
				},
			},
		},
	}

	tests := []struct {
		filter Filter
		result []string
	}{
		{
			filter: Filter{
				Conditions: []FilterCondition{
					&StringPropertyFilterCondition{
						ID: "squeue.account",
						Checker: &StringEqualValueChecker{
							ExpectedValue: "nwp_xp",
						},
					},
				},
			},
			result: []string{"1"},
		},
		{
			filter: Filter{
				Conditions: []FilterCondition{
					&StringPropertyFilterCondition{
						ID: "squeue.account",
						Checker: &StringContainChecker{
							ExpectedValue: "nwp",
						},
					},
					&StringPropertyFilterCondition{
						ID: "squeue.partition",
						Checker: &StringEqualValueChecker{
							ExpectedValue: "serial_op",
						},
					},
				},
			},
			result: []string{"2"},
		},
	}

	for _, test := range tests {
		targetItems := test.filter.Filter(items)
		for index, id := range test.result {
			targetItem := targetItems[index]
			prop := targetItem.GetProperty("squeue.jobid").(*StringProperty)
			if prop.Data != id {
				t.Errorf("target id %s != requred id %s", prop.Data, id)
			}
		}
	}
}
