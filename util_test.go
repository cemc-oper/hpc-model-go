package hpcmodel_test

import (
	. "github.com/perillaroc/nwpc-hpc-model-go"
	"testing"
	"time"
)

func TestCreateSorter(t *testing.T) {
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
					Data: "eps_xp",
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
					Data: "nwp_xp",
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

	accountSort := func(item1, item2 *Item) bool {
		p1 := item1.GetProperty("squeue.account").(*StringProperty)
		p2 := item2.GetProperty("squeue.account").(*StringProperty)
		return p1.Data < p2.Data
	}

	jobidSort := func(item1, item2 *Item) bool {
		p1 := item1.GetProperty("squeue.jobid").(*StringProperty)
		p2 := item2.GetProperty("squeue.jobid").(*StringProperty)
		return p1.Data < p2.Data
	}

	sorter := CreateSorter(accountSort, jobidSort)
	sorter.Sort(items)

	if items[0].GetProperty("squeue.jobid").(*StringProperty).Data != "2" {
		t.Errorf("item 0 is not 2")
	}

	if items[1].GetProperty("squeue.jobid").(*StringProperty).Data != "1" {
		t.Errorf("item 1 is not 1")
	}

	if items[2].GetProperty("squeue.jobid").(*StringProperty).Data != "3" {
		t.Errorf("item 2 is not 3")
	}

}

func TestCreateStringPropertyLessFunc(t *testing.T) {
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
					Data: "eps_xp",
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
					Data: "nwp_xp",
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

	accountSort := CreatePropertyLessFunc("squeue.account")
	jobidSort := CreatePropertyLessFunc("squeue.jobid")

	sorter := CreateSorter(accountSort, jobidSort)
	sorter.Sort(items)

	if items[0].GetProperty("squeue.jobid").(*StringProperty).Data != "2" {
		t.Errorf("item 0 is not 2")
	}

	if items[1].GetProperty("squeue.jobid").(*StringProperty).Data != "1" {
		t.Errorf("item 1 is not 1")
	}

	if items[2].GetProperty("squeue.jobid").(*StringProperty).Data != "3" {
		t.Errorf("item 2 is not 3")
	}
}
