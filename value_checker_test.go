package hpcmodel_test

import (
	"github.com/perillaroc/nwpc-hpc-model-go"
	"testing"
	"time"
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

func TestNumberEqualValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue float64
		data          float64
		result        bool
	}{
		{
			1,
			1,
			true,
		},
		{
			2,
			2,
			true,
		},
		{
			1,
			2,
			false,
		},
		{
			2,
			2.5,
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.NumberEqualValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%g != %g", test.data, test.expectedValue)
		}
	}
}

func TestNumberGreaterValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue float64
		data          float64
		result        bool
	}{
		{
			1,
			2,
			true,
		},
		{
			2,
			3,
			true,
		},
		{
			1,
			1,
			false,
		},
		{
			2,
			1,
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.NumberGreaterValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%g <= %g", test.data, test.expectedValue)
		}
	}
}

func TestNumberLessValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue float64
		data          float64
		result        bool
	}{
		{
			2,
			1,
			true,
		},
		{
			3,
			2,
			true,
		},
		{
			1,
			2,
			false,
		},
		{
			2,
			2.5,
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.NumberLessValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%g >= %g", test.data, test.expectedValue)
		}
	}
}

func TestNumberInValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValues []float64
		data           float64
		result         bool
	}{
		{
			[]float64{2, 3, 4},
			2,
			true,
		},
		{
			[]float64{2, 3, 4},
			3,
			true,
		},
		{
			[]float64{2, 3, 4},
			2.5,
			false,
		},
		{
			[]float64{2, 3, 4},
			1,
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.NumberInValueChecker{
			ExpectedValues: test.expectedValues,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%g not in %v", test.data, test.expectedValues)
		}
	}
}

func TestDateTimeEqualValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue time.Time
		data          time.Time
		result        bool
	}{
		{
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			true,
		},
		{
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			time.Date(
				2019, time.January, 5, 20, 4, 0, 0, time.UTC),
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.DateTimeEqualValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%v not equal %v", test.data, test.expectedValue)
		}
	}
}

func TestDateTimeAfterValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue time.Time
		data          time.Time
		result        bool
	}{
		{
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 5, 0, 0, time.UTC),
			true,
		},
		{
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			false,
		},
		{
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 3, 0, 0, time.UTC),
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.DateTimeAfterValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%v after %v %t, required %t",
				test.data, test.expectedValue, !test.result, test.result)
		}
	}
}

func TestDateTimeBeforeValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValue time.Time
		data          time.Time
		result        bool
	}{
		{
			time.Date(
				2019, time.February, 5, 20, 5, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			true,
		},
		{
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			false,
		},
		{
			time.Date(
				2019, time.February, 5, 20, 3, 0, 0, time.UTC),
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.DateTimeBeforeValueChecker{
			ExpectedValue: test.expectedValue,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%v before %v %t, required %t",
				test.data, test.expectedValue, !test.result, test.result)
		}
	}
}

func TestDateTimeInValueChecker_CheckValue(t *testing.T) {
	tests := []struct {
		expectedValues []time.Time
		data           time.Time
		result         bool
	}{
		{
			[]time.Time{
				time.Date(2019, time.February, 5, 20, 5, 0, 0, time.UTC),
				time.Date(2019, time.February, 5, 20, 4, 0, 0, time.UTC),
				time.Date(2019, time.February, 5, 20, 3, 0, 0, time.UTC),
			},
			time.Date(
				2019, time.February, 5, 20, 4, 0, 0, time.UTC),
			true,
		},
		{
			[]time.Time{
				time.Date(2019, time.February, 5, 20, 5, 0, 0, time.UTC),
				time.Date(2019, time.February, 5, 20, 4, 0, 0, time.UTC),
				time.Date(2019, time.February, 5, 20, 3, 0, 0, time.UTC),
			},
			time.Date(
				2019, time.February, 5, 20, 3, 0, 0, time.UTC),
			true,
		},
		{
			[]time.Time{
				time.Date(2019, time.February, 5, 20, 5, 0, 0, time.UTC),
				time.Date(2019, time.February, 5, 20, 4, 0, 0, time.UTC),
				time.Date(2019, time.February, 5, 20, 3, 0, 0, time.UTC),
			},
			time.Date(
				2019, time.February, 5, 20, 2, 0, 0, time.UTC),
			false,
		},
	}
	for _, test := range tests {
		checker := hpcmodel.DateTimeInValueChecker{
			ExpectedValues: test.expectedValues,
		}

		if checker.CheckValue(test.data) != test.result {
			t.Errorf("%v in %v %t, required %t",
				test.data, test.expectedValues, !test.result, test.result)
		}
	}
}
