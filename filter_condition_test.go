package hpcmodel_test

import (
	"github.com/perillaroc/nwpc-hpc-model-go"
	"testing"
)

func TestStringEqualValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue string
		data          string
		result        bool
	}{
		{
			"nwp",
			"nwp",
			true,
		},
		{
			"nwp",
			"nwp_qu",
			false,
		},
		{
			"nwp",
			"eps_nwp",
			false,
		},
		{
			"nwp",
			"pos_xp",
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.StringEqualValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%s != %s", test.data, test.expectedValue)
		}
	}
}

func TestStringInValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValues []string
		data           string
		result         bool
	}{
		{
			[]string{"nwp", "nwp_qu", "nwp_pd"},
			"nwp",
			true,
		},
		{
			[]string{"nwp", "nwp_qu", "nwp_pd"},
			"nwp_pd",
			true,
		},
		{
			[]string{"nwp", "nwp_qu", "nwp_pd"},
			"nwp_qu",
			true,
		},
		{
			[]string{"nwp", "nwp_qu", "nwp_pd"},
			"pos_xp",
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.StringInValueChecker{
			ExpectedValues: test.expectedValues,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%s not in %v", test.data, test.expectedValues)
		}
	}
}

func TestStringContainChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue string
		data          string
		result        bool
	}{
		{
			"nwp",
			"nwp",
			true,
		},
		{
			"nwp",
			"nwp_qu",
			true,
		},
		{
			"nwp",
			"eps_nwp",
			true,
		},
		{
			"nwp",
			"pos_xp",
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.StringContainChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%s != %s", test.data, test.expectedValue)
		}
	}
}
