package hpcmodel_test

import (
	"github.com/perillaroc/nwpc-hpc-model-go"
	"testing"
	"time"
)

func TestItem(t *testing.T) {
	var item hpcmodel.Item

	var prop hpcmodel.StringProperty
	prop.SetValue("wdp")

	item.Props = append(item.Props, &prop)
}

func TestBuildItem_StringProperty(t *testing.T) {
	type TestProperty struct {
		id    string
		value string
		data  string
		text  string
	}

	type TestCase struct {
		categoryList hpcmodel.QueryCategoryList
		records      []string
		properties   []TestProperty
	}

	var tests = []TestCase{
		{
			hpcmodel.QueryCategoryList{
				CategoryList: []*hpcmodel.QueryCategory{
					{
						ID:          "squeue.account",
						DisplayName: "account",
						Label:       "Account",
						ParseRecord: &hpcmodel.TokenRecordParser{
							Index: 0,
							Sep:   "|",
						},
						PropertyClass:           "StringProperty",
						PropertyCreateArguments: []string{},
					},
					{
						ID:          "squeue.job_id",
						DisplayName: "JOB ID",
						Label:       "JOBID",
						ParseRecord: &hpcmodel.TokenRecordParser{
							Index: 7,
							Sep:   "|",
						},
						PropertyClass:           "StringProperty",
						PropertyCreateArguments: []string{},
					},
				},
			},
			[]string{"nwp_cy|N/A1|0|2019-02-16T09:05:58|(null)|nwpep|OK|7352547|GRAPES|GRAPES|15-00:00:00|100G||/g3/nwp_cy/Eps/wangjzh/GRAPES_MEPS_NCEP/ens_sh/run_filter.sh|0.00000011641532|normal1|None||R|nwp_cy|(null)|(null)||0|*:*:*|7352547|cmbc0011|1|1||7352547|10203|*|*|*|N/A|14-23:57:23|2:37|cmbc0011|0|serial|500|cmbc0011|2019-02-01T09:05:58|RUNNING|1020302|2019-02-01T09:05:56|(null)|N/A|(null)|/g3/nwp_cy/Eps/wangjzh/GRAPES_MEPS_NCEP/ens_sh"},
			[]TestProperty{
				{
					"squeue.account",
					"nwp_cy",
					"nwp_cy",
					"nwp_cy",
				},
				{
					"squeue.job_id",
					"7352547",
					"7352547",
					"7352547",
				},
			},
		},
	}
	for _, test := range tests {
		item, err := hpcmodel.BuildItem(test.records, test.categoryList)
		if err != nil {
			t.Errorf("build item failed: %v", err)
		}
		for _, propertyTest := range test.properties {
			p := item.GetProperty(propertyTest.id)
			if p == nil {
				t.Errorf("property %s is not found", propertyTest.id)
			}

			sp, ok := p.(*hpcmodel.StringProperty)
			if !ok {
				t.Errorf("property %s is not StringProperty", propertyTest.id)
			}

			if sp.Value != propertyTest.value {
				t.Errorf("p.Text != %s", propertyTest.value)
			}
			if sp.Data != propertyTest.data {
				t.Errorf("p.Data != %s", propertyTest.data)
			}
			if sp.Text != propertyTest.text {
				t.Errorf("p.Text != %s", propertyTest.text)
			}
		}
	}
}

func TestBuildItem_NumberProperty(t *testing.T) {
	type TestProperty struct {
		id    string
		value string
		data  float64
		text  string
	}

	type TestCase struct {
		categoryList hpcmodel.QueryCategoryList
		records      []string
		properties   []TestProperty
	}

	var tests = []TestCase{
		{
			hpcmodel.QueryCategoryList{
				CategoryList: []*hpcmodel.QueryCategory{
					{
						ID:          "squeue.cpus",
						DisplayName: "Cpus",
						Label:       "CPUS",
						ParseRecord: &hpcmodel.TokenRecordParser{
							Index: 27,
							Sep:   "|",
						},
						PropertyClass:           "NumberProperty",
						PropertyCreateArguments: []string{},
					},
					{
						ID:          "squeue.nodes",
						DisplayName: "Nodes",
						Label:       "NODES",
						ParseRecord: &hpcmodel.TokenRecordParser{
							Index: 28,
							Sep:   "|",
						},
						PropertyClass:           "NumberProperty",
						PropertyCreateArguments: []string{},
					},
				},
			},
			[]string{"nwp_cy|N/A1|0|2019-02-16T09:05:58|(null)|nwpep|OK|7352547|GRAPES|GRAPES|15-00:00:00|100G||/g3/nwp_cy/Eps/wangjzh/GRAPES_MEPS_NCEP/ens_sh/run_filter.sh|0.00000011641532|normal1|None||R|nwp_cy|(null)|(null)||0|*:*:*|7352547|cmbc0011|1|1||7352547|10203|*|*|*|N/A|14-23:57:23|2:37|cmbc0011|0|serial|500|cmbc0011|2019-02-01T09:05:58|RUNNING|1020302|2019-02-01T09:05:56|(null)|N/A|(null)|/g3/nwp_cy/Eps/wangjzh/GRAPES_MEPS_NCEP/ens_sh"},
			[]TestProperty{
				{
					"squeue.cpus",
					"1",
					float64(1),
					"1",
				},
				{
					"squeue.nodes",
					"1",
					float64(1),
					"1",
				},
			},
		},
	}
	for _, test := range tests {
		item, err := hpcmodel.BuildItem(test.records, test.categoryList)
		if err != nil {
			t.Errorf("build item failed: %v", err)
		}
		for _, propertyTest := range test.properties {
			p := item.GetProperty(propertyTest.id)
			if p == nil {
				t.Errorf("property %s is not found", propertyTest.id)
			}

			sp, ok := p.(*hpcmodel.NumberProperty)
			if !ok {
				t.Errorf("property %s is not NumberProperty", propertyTest.id)
			}

			if sp.Value != propertyTest.value {
				t.Errorf("p.Text != %s", propertyTest.value)
			}
			if sp.Data != propertyTest.data {
				t.Errorf("p.Data != %f", propertyTest.data)
			}
			if sp.Text != propertyTest.text {
				t.Errorf("p.Text != %s", propertyTest.text)
			}
		}
	}
}

func TestBuildItem_DateTimeProperty(t *testing.T) {
	type TestProperty struct {
		id    string
		value string
		data  time.Time
		text  string
	}

	type TestCase struct {
		categoryList hpcmodel.QueryCategoryList
		records      []string
		properties   []TestProperty
	}

	var tests = []TestCase{
		{
			hpcmodel.QueryCategoryList{
				CategoryList: []*hpcmodel.QueryCategory{
					{
						ID:          "squeue.start_time",
						DisplayName: "Start time",
						Label:       "START_TIME",
						ParseRecord: &hpcmodel.TokenRecordParser{
							Index: 43,
							Sep:   "|",
						},
						PropertyClass:           "DateTimeProperty",
						PropertyCreateArguments: []string{},
					},
				},
			},
			[]string{"nwp_cy|N/A1|0|2019-02-16T09:05:58|(null)|nwpep|OK|7352547|GRAPES|GRAPES|15-00:00:00|100G||/g3/nwp_cy/Eps/wangjzh/GRAPES_MEPS_NCEP/ens_sh/run_filter.sh|0.00000011641532|normal1|None||R|nwp_cy|(null)|(null)||0|*:*:*|7352547|cmbc0011|1|1||7352547|10203|*|*|*|N/A|14-23:57:23|2:37|cmbc0011|0|serial|500|cmbc0011|2019-02-01T09:05:58|RUNNING|1020302|2019-02-01T09:05:56|(null)|N/A|(null)|/g3/nwp_cy/Eps/wangjzh/GRAPES_MEPS_NCEP/ens_sh"},
			[]TestProperty{
				{
					"squeue.start_time",
					"2019-02-01T09:05:58",
					time.Date(2019, time.February, 1, 9, 5, 58, 0, time.UTC),
					"2019-02-01 09:05:58",
				},
			},
		},
	}
	for _, test := range tests {
		item, err := hpcmodel.BuildItem(test.records, test.categoryList)
		if err != nil {
			t.Errorf("build item failed: %v", err)
		}
		for _, propertyTest := range test.properties {
			p := item.GetProperty(propertyTest.id)
			if p == nil {
				t.Errorf("property %s is not found", propertyTest.id)
			}

			sp, ok := p.(*hpcmodel.DateTimeProperty)
			if !ok {
				t.Errorf("property %s is not NumberProperty", propertyTest.id)
			}

			if sp.Value != propertyTest.value {
				t.Errorf("p.Text != %s", propertyTest.value)
			}
			if sp.Data != propertyTest.data {
				t.Errorf("p.Data != %s", propertyTest.data)
			}
			if sp.Text != propertyTest.text {
				t.Errorf("p.Text != %s", propertyTest.text)
			}
		}
	}
}
