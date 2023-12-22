package hpcmodel_test

import (
	hpcmodel "github.com/cemc-oper/hpc-model-go"
	"testing"
	"time"
)

func TestFilter_Filter(t *testing.T) {
	items := []hpcmodel.Item{
		{
			Props: []hpcmodel.Property{
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.jobid",
					},
					Data: "1",
				},
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
				&hpcmodel.DateTimeProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.submit_time",
					},
					Data: time.Date(2019, time.February, 6, 18, 00, 00, 00, time.UTC),
				},
			},
		},
		{
			Props: []hpcmodel.Property{
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.jobid",
					},
					Data: "2",
				},
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.account",
					},
					Data: "nwp",
				},
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.partition",
					},
					Data: "serial_op",
				},
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.state",
					},
					Data: "RUNNING",
				},
				&hpcmodel.DateTimeProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.submit_time",
					},
					Data: time.Date(2019, time.February, 6, 20, 00, 00, 00, time.UTC),
				},
			},
		},
		{
			Props: []hpcmodel.Property{
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.jobid",
					},
					Data: "3",
				},
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.account",
					},
					Data: "windroc",
				},
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.partition",
					},
					Data: "serial_op",
				},
				&hpcmodel.StringProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.state",
					},
					Data: "PENDING",
				},
				&hpcmodel.DateTimeProperty{
					Category: hpcmodel.QueryCategory{
						ID: "squeue.submit_time",
					},
					Data: time.Date(2019, time.February, 6, 19, 00, 00, 00, time.UTC),
				},
			},
		},
	}

	tests := []struct {
		filter hpcmodel.Filter
		result []string
	}{
		{
			filter: hpcmodel.Filter{
				Conditions: []hpcmodel.FilterCondition{
					&hpcmodel.StringPropertyFilterCondition{
						ID: "squeue.account",
						Checker: &hpcmodel.StringEqualValueChecker{
							ExpectedValue: "nwp_xp",
						},
					},
				},
			},
			result: []string{"1"},
		},
		{
			filter: hpcmodel.Filter{
				Conditions: []hpcmodel.FilterCondition{
					&hpcmodel.StringPropertyFilterCondition{
						ID: "squeue.account",
						Checker: &hpcmodel.StringContainChecker{
							ExpectedValue: "nwp",
						},
					},
					&hpcmodel.StringPropertyFilterCondition{
						ID: "squeue.partition",
						Checker: &hpcmodel.StringEqualValueChecker{
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
			prop := targetItem.GetProperty("squeue.jobid").(*hpcmodel.StringProperty)
			if prop.Data != id {
				t.Errorf("target id %s != requred id %s", prop.Data, id)
			}
		}
	}
}
