package slurm_test

import (
	"strings"
	"testing"

	hpcmodel "github.com/cemc-oper/hpc-model-go"
	"github.com/cemc-oper/hpc-model-go/slurm"
)

func newModelTestCategoryList() slurm.QueryCategoryList {
	return slurm.QueryCategoryList{
		QueryCategoryList: hpcmodel.QueryCategoryList{
			CategoryList: []*hpcmodel.QueryCategory{
				{
					ID:                      "squeue.account",
					DisplayName:             "account",
					Label:                   "ACCOUNT",
					PropertyClass:           "StringProperty",
					PropertyCreateArguments: []string{},
					RecordParserClass:       "TokenRecordParser",
				},
				{
					ID:                      "squeue.job_id",
					DisplayName:             "JOB ID",
					Label:                   "JOBID",
					PropertyClass:           "StringProperty",
					PropertyCreateArguments: []string{},
					RecordParserClass:       "TokenRecordParser",
				},
			},
		},
	}
}

func mustBuildModel(t *testing.T, records []string, categoryList slurm.QueryCategoryList, sep string) *slurm.Model {
	t.Helper()
	model, err := slurm.BuildModel(records, categoryList, sep)
	if err != nil {
		t.Fatalf("BuildModel() error = %v", err)
	}
	if model == nil {
		t.Fatal("BuildModel() returned a nil model with nil error")
	}
	return model
}

func TestBuildModel(t *testing.T) {
	categoryList := newModelTestCategoryList()
	records := []string{
		`ACCOUNT|TRES_PER_NODEMIN_CPUS|MIN_TMP_DISK|END_TIME|FEATURES|GROUP|OVER_SUBSCRIBE|JOBID|NAME|COMMENT|TIME_LIMIT|MIN_MEMORY|REQ_NODES|COMMAND|PRIORITY|QOS|REASON||ST|USER|RESERVATION|WCKEY|EXC_NODES|NICE|S:C:T|JOBID|EXEC_HOST|CPUS|NODES|DEPENDENCY|ARRAY_JOB_ID|GROUP|SOCKETS_PER_NODE|CORES_PER_SOCKET|THREADS_PER_CORE|ARRAY_TASK_ID|TIME_LEFT|TIME|NODELIST|CONTIGUOUS|PARTITION|PRIORITY|NODELIST(REASON)|START_TIME|STATE|UID|SUBMIT_TIME|LICENSES|CORE_SPEC|SCHEDNODES|WORK_DIR`,
		`chendh|N/A32|0|2019-02-13T20:30:06|(null)|nwpc|OK|5836542|GRAPES|GRAPES|15-00:00:00|5G||/g8/JOB_TMP/chendh/ShCu/RUN_24/grapes.sbatch|0.00000011641532|middle|None||R|chendh|(null)|(null)||0|*:*:*|5836542|cmac0819|1024|32||5836542|10201|*|*|*|N/A|12-11:23:44|2-12:36:16|cmac[0819-0832,0835-0850,0897-0898]|0|normal|500|cmac[0819-0832,0835-0850,0897-0898]|2019-01-29T20:30:06|RUNNING|1020104|2019-01-29T08:41:45|(null)|N/A|(null)|/g8/JOB_TMP/chendh/ShCu/RUN_24`,
		`lijl|N/A32|0|2019-02-13T00:37:51|(null)|csmd|OK|5831234|BCC800|BCCCSM|15-00:00:00|5G||/g6/lijl/BCC_CSMv3.v20190124/p25_2/build.csh|0.00000011641532|normal|None||R|lijl|(null)|(null)||0|*:*:*|5831234|cmac0243|704|22||5831234|10304|*|*|*|N/A|11-15:31:29|3-08:28:31|cmac[0243-0248,1155-1170]|0|normal|500|cmac[0243-0248,1155-1170]|2019-01-29T00:37:51|RUNNING|1030410|2019-01-29T00:37:49|(null)|N/A|(null)|/g6/lijl/BCC_CSMv3.v20190124/p25_2`,
		`wutw|N/A32|0|2019-02-07T10:28:35|(null)|csmd|OK|5720780|BCCCSM|WRF|15-00:00:00|5G||/g6/wutw/BCC_CSMv3_T266p13_OK/T266L56.deep.955/build.csh|0.00000011641532|normal|None||R|wutw|(null)|(null)||0|*:*:*|5720780|cmac0272|576|18||5720780|10304|*|*|*|N/A|6-01:22:13|8-22:37:47|cmac[0272-0273,0499-0514]|0|normal|500|cmac[0272-0273,0499-0514]|2019-01-23T10:28:35|RUNNING|1030402|2019-01-23T10:28:35|(null)|N/A|(null)|/g6/wutw/BCC_CSMv3_T266p13_OK/T266L56.deep.955`,
	}

	type propertyTest struct {
		ID   string
		text string
	}
	tests := []struct {
		index      int
		properties []propertyTest
	}{
		{index: 0, properties: []propertyTest{{ID: "squeue.account", text: "chendh"}, {ID: "squeue.job_id", text: "5836542"}}},
		{index: 2, properties: []propertyTest{{ID: "squeue.account", text: "wutw"}, {ID: "squeue.job_id", text: "5720780"}}},
	}

	model := mustBuildModel(t, records, categoryList, "|")
	if got, want := len(model.Items), 3; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
	for _, test := range tests {
		item := model.Items[test.index]
		for _, pt := range test.properties {
			property := item.GetProperty(pt.ID)
			if property == nil {
				t.Fatalf("property %s not found", pt.ID)
			}
			var text string
			switch p := property.(type) {
			case *hpcmodel.StringProperty:
				text = p.Text
			case *hpcmodel.NumberProperty:
				text = p.Text
			case *hpcmodel.DateTimeProperty:
				text = p.Text
			case *hpcmodel.TimestampProperty:
				text = p.Text
			case *hpcmodel.TimeStringProperty:
				text = p.Text
			default:
				t.Fatalf("property %s has unsupported type %T", pt.ID, property)
			}
			if text != pt.text {
				t.Errorf("property %s text = %q, want %q", pt.ID, text, pt.text)
			}
		}
	}
}

func TestBuildModelEmptyInput(t *testing.T) {
	tests := []struct {
		name    string
		records []string
	}{
		{name: "nil", records: nil},
		{name: "empty slice", records: []string{}},
		{name: "blank lines", records: []string{"", "  ", "\t"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := mustBuildModel(t, tt.records, newModelTestCategoryList(), "|")
			if len(model.Items) != 0 {
				t.Errorf("len(Items) = %d, want 0", len(model.Items))
			}
		})
	}
}

func TestBuildModelSkipsBlankLines(t *testing.T) {
	records := []string{"ACCOUNT|JOBID", "", "  ", "alice|1001", "\t", ""}
	model := mustBuildModel(t, records, newModelTestCategoryList(), "|")
	if got, want := len(model.Items), 1; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
	property := model.Items[0].GetProperty("squeue.job_id").(*hpcmodel.StringProperty)
	if property.Data != "1001" {
		t.Errorf("job ID = %q, want %q", property.Data, "1001")
	}
}

func TestBuildModelSkipsLeadingBlankLines(t *testing.T) {
	records := []string{"", "  ", "ACCOUNT|JOBID", "alice|1001"}
	model := mustBuildModel(t, records, newModelTestCategoryList(), "|")
	if got, want := len(model.Items), 1; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
}

func TestBuildModelMissingColumns(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		missing []string
	}{
		{name: "one column", header: "ACCOUNT", missing: []string{"JOBID"}},
		{name: "multiple columns", header: "USER", missing: []string{"ACCOUNT", "JOBID"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := slurm.BuildModel([]string{tt.header}, newModelTestCategoryList(), "|")
			if err == nil {
				t.Fatal("BuildModel() error = nil, want missing column error")
			}
			for _, label := range tt.missing {
				if !strings.Contains(err.Error(), label) {
					t.Errorf("error = %q, want it to contain %q", err, label)
				}
			}
		})
	}
}

func TestBuildModelMissingLabelOccurrence(t *testing.T) {
	categoryList := newModelTestCategoryList()
	categoryList.CategoryList[0].LabelOccurrence = 1
	_, err := slurm.BuildModel([]string{"ACCOUNT|JOBID"}, categoryList, "|")
	if err == nil {
		t.Fatal("BuildModel() error = nil, want missing occurrence error")
	}
	if !strings.Contains(err.Error(), "ACCOUNT") {
		t.Errorf("error = %q, want it to contain ACCOUNT", err)
	}
}

func TestBuildModelColumnCountMismatch(t *testing.T) {
	tests := []struct {
		name        string
		record      string
		wantColumns string
	}{
		{name: "extra columns", record: "alice|1001|extra", wantColumns: "3 columns"},
		{name: "missing columns", record: "alice", wantColumns: "1 columns"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := slurm.BuildModel([]string{"ACCOUNT|JOBID", tt.record}, newModelTestCategoryList(), "|")
			if err == nil {
				t.Fatal("BuildModel() error = nil, want column count error")
			}
			if !strings.Contains(err.Error(), "line 2") || !strings.Contains(err.Error(), tt.wantColumns) || !strings.Contains(err.Error(), "want 2") {
				t.Errorf("error = %q, want line number and column counts", err)
			}
		})
	}
}

func TestBuildModelColumnCountErrorReportsOriginalLine(t *testing.T) {
	records := []string{"ACCOUNT|JOBID", "alice|1001", "bob"}
	_, err := slurm.BuildModel(records, newModelTestCategoryList(), "|")
	if err == nil {
		t.Fatal("BuildModel() error = nil, want column count error")
	}
	if !strings.Contains(err.Error(), "line 3") {
		t.Errorf("error = %q, want line 3", err)
	}
}

func TestBuildModelDelimiterOnlyRecord(t *testing.T) {
	model := mustBuildModel(t, []string{"ACCOUNT|JOBID", "|"}, newModelTestCategoryList(), "|")
	if got, want := len(model.Items), 1; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
	for _, id := range []string{"squeue.account", "squeue.job_id"} {
		property := model.Items[0].GetProperty(id).(*hpcmodel.StringProperty)
		if property.Data != "" {
			t.Errorf("%s = %q, want an empty value", id, property.Data)
		}
	}
}

func TestBuildModelPropagatesItemErrors(t *testing.T) {
	categoryList := slurm.QueryCategoryList{
		QueryCategoryList: hpcmodel.QueryCategoryList{
			CategoryList: []*hpcmodel.QueryCategory{
				{
					ID:                "squeue.account",
					Label:             "ACCOUNT",
					PropertyClass:     "UnsupportedProperty",
					RecordParserClass: "TokenRecordParser",
				},
			},
		},
	}
	_, err := slurm.BuildModel([]string{"ACCOUNT", "alice"}, categoryList, "|")
	if err == nil {
		t.Fatal("BuildModel() error = nil, want property construction error")
	}
	if !strings.Contains(err.Error(), "line 2") || !strings.Contains(err.Error(), "UnsupportedProperty") {
		t.Errorf("error = %q, want line number and property class", err)
	}
}
