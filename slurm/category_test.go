package slurm_test

import (
	hpcmodel "github.com/cemc-oper/hpc-model-go"
	"github.com/cemc-oper/hpc-model-go/slurm"
	"strconv"
	"testing"
)

func TestQueryCategoryList_UpdateTokenIndex(t *testing.T) {
	categoryList := &slurm.QueryCategoryList{
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

	tests := []struct {
		label string
		index int
		sep   string
	}{
		{
			"ACCOUNT",
			0,
			"|",
		},
		{
			"JOBID",
			26,
			"|",
		},
	}

	titleLine := "ACCOUNT|GRES|MIN_CPUS|MIN_TMP_DISK|END_TIME|FEATURES|GROUP|OVER_SUBSCRIBE" +
		"|JOBID|NAME|COMMENT|TIME_LIMIT|MIN_MEMORY|REQ_NODES|COMMAND|PRIORITY|QOS|REASON||ST|" +
		"USER|RESERVATION|WCKEY|EXC_NODES|NICE|S:C:T|JOBID|EXEC_HOST|CPUS|NODES|DEPENDENCY|ARRAY_JOB_ID|" +
		"GROUP|SOCKETS_PER_NODE|CORES_PER_SOCKET|THREADS_PER_CORE|ARRAY_TASK_ID|TIME_LEFT|TIME|" +
		"NODELIST|CONTIGUOUS|PARTITION|PRIORITY|NODELIST(REASON)|START_TIME|STATE|USER|SUBMIT_TIME|" +
		"LICENSES|CORE_SPEC|SCHEDNODES|WORK_DIR"

	categoryList.UpdateTokenIndex(titleLine, "|")
	for _, test := range tests {
		category := categoryList.CategoryFromLabel(test.label)
		if category == nil {
			t.Errorf("can't find category with label %s", test.label)
		}
		args := category.RecordParserArguments
		if len(args) != 2 {
			t.Errorf("RecordParserArguments length is not 2")
		}
		if args[0] != strconv.Itoa(test.index) {
			t.Errorf("index is %s, should %d", args[0], test.index)
		}
		if args[1] != test.sep {
			t.Errorf("sep is %s, should %s", args[0], test.sep)
		}
	}
}

func TestBuildModel2(t *testing.T) {
	categoryList := &slurm.QueryCategoryList{
		QueryCategoryList: hpcmodel.QueryCategoryList{
			CategoryList: []*hpcmodel.QueryCategory{
				{
					ID:                      "sinfo.partition",
					DisplayName:             "Partition",
					Label:                   "PARTITION",
					PropertyClass:           "StringProperty",
					PropertyCreateArguments: []string{},
					RecordParserClass:       "TokenRecordParser",
				},
				{
					ID:                      "squeue.avail",
					DisplayName:             "Avail",
					Label:                   "AVAIL",
					PropertyClass:           "StringProperty",
					PropertyCreateArguments: []string{},
					RecordParserClass:       "TokenRecordParser",
				},
			},
		},
	}

	tests := []struct {
		label string
		index int
		sep   string
	}{
		{
			"PARTITION",
			0,
			"",
		},
		{
			"AVAIL",
			1,
			"",
		},
	}

	titleLine := "PARTITION            AVAIL       NODES(A/I/O/T)                  CPUS(A/I/O/T)"

	categoryList.UpdateTokenIndex(titleLine, "")
	for _, test := range tests {
		category := categoryList.CategoryFromLabel(test.label)
		if category == nil {
			t.Errorf("can't find category with label %s", test.label)
		}
		args := category.RecordParserArguments
		if len(args) != 2 {
			t.Errorf("RecordParserArguments length is not 2")
		}
		if args[0] != strconv.Itoa(test.index) {
			t.Errorf("index is %s, should %d", args[0], test.index)
		}
		if args[1] != test.sep {
			t.Errorf("sep is %s, should %s", args[0], test.sep)
		}
	}
}
