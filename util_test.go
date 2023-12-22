package hpcmodel_test

import (
	hpcmodel "github.com/cemc-oper/hpc-model-go"
	"testing"
	"time"
)

func TestCreateSorter(t *testing.T) {
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
					Data: "eps_xp",
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
					Data: "nwp_xp",
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

	accountSort := func(item1, item2 *hpcmodel.Item) bool {
		p1 := item1.GetProperty("squeue.account").(*hpcmodel.StringProperty)
		p2 := item2.GetProperty("squeue.account").(*hpcmodel.StringProperty)
		return p1.Data < p2.Data
	}

	jobidSort := func(item1, item2 *hpcmodel.Item) bool {
		p1 := item1.GetProperty("squeue.jobid").(*hpcmodel.StringProperty)
		p2 := item2.GetProperty("squeue.jobid").(*hpcmodel.StringProperty)
		return p1.Data < p2.Data
	}

	sorter := hpcmodel.CreateSorter(accountSort, jobidSort)
	sorter.Sort(items)

	if items[0].GetProperty("squeue.jobid").(*hpcmodel.StringProperty).Data != "2" {
		t.Errorf("item 0 is not 2")
	}

	if items[1].GetProperty("squeue.jobid").(*hpcmodel.StringProperty).Data != "1" {
		t.Errorf("item 1 is not 1")
	}

	if items[2].GetProperty("squeue.jobid").(*hpcmodel.StringProperty).Data != "3" {
		t.Errorf("item 2 is not 3")
	}

}

func TestCreateStringPropertyLessFunc(t *testing.T) {
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
					Data: "eps_xp",
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
					Data: "nwp_xp",
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

	accountSort := hpcmodel.CreatePropertyLessFunc("squeue.account")
	jobidSort := hpcmodel.CreatePropertyLessFunc("squeue.jobid")

	sorter := hpcmodel.CreateSorter(accountSort, jobidSort)
	sorter.Sort(items)

	if items[0].GetProperty("squeue.jobid").(*hpcmodel.StringProperty).Data != "2" {
		t.Errorf("item 0 is not 2")
	}

	if items[1].GetProperty("squeue.jobid").(*hpcmodel.StringProperty).Data != "1" {
		t.Errorf("item 1 is not 1")
	}

	if items[2].GetProperty("squeue.jobid").(*hpcmodel.StringProperty).Data != "3" {
		t.Errorf("item 2 is not 3")
	}
}
