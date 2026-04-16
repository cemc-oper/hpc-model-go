package hpcmodel_test

import (
	"testing"
	"time"

	hpcmodel "github.com/cemc-oper/hpc-model-go"
)

func TestStringProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  string
		text  string
	}{
		{"nwpc", "nwpc", "nwpc", "nwpc"},
	}

	for _, test := range tests {
		var p hpcmodel.StringProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestNumberProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  string
		text  string
	}{
		{"nwpc", "nwpc", "nwpc", "nwpc"},
	}

	for _, test := range tests {
		var p hpcmodel.StringProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestDateTimeProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  time.Time
		text  string
	}{
		{
			"2019-01-29T19:45:00",
			"2019-01-29T19:45:00",
			time.Date(2019, time.January, 29, 19, 45, 0, 0, time.UTC),
			"2019-01-29 19:45:00",
		},
	}

	for _, test := range tests {
		var p hpcmodel.DateTimeProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestTimestampProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  time.Time
		text  string
	}{
		{
			"1547906480",
			"1547906480",
			time.Date(2019, time.January, 19, 14, 01, 20, 0, time.UTC),
			"2019-01-19 14:01:20",
		},
	}

	for _, test := range tests {
		var p hpcmodel.TimestampProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestBuildProperty_StringProperty(t *testing.T) {
	var tests = []struct {
		category hpcmodel.QueryCategory
		records  []string
		value    string
		data     string
		text     string
	}{
		{
			hpcmodel.QueryCategory{
				ID:          "id1",
				DisplayName: "owner",
				Label:       "owner",
				ParseRecord: &hpcmodel.TokenRecordParser{
					Index: 0,
					Sep:   "|",
				},
				PropertyClass: "StringProperty",
			},
			[]string{"nwp_xp|(null)|1|0|NONE"},
			"nwp_xp",
			"nwp_xp",
			"nwp_xp",
		},
	}
	for _, test := range tests {
		p, err := hpcmodel.BuildProperty(test.records, test.category)
		if err != nil {
			t.Errorf("build property failed: %v", err)
		}
		sp, ok := p.(*hpcmodel.StringProperty)
		if !ok {
			t.Errorf("property is not StringProperty")
		}
		if sp.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if sp.Data != test.data {
			t.Errorf("p.Data != %s", test.data)
		}
		if sp.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}

	}
}

func TestTimeStringProperty_SetValue(t *testing.T) {
	var tests = []struct {
		s     string
		value string
		data  time.Duration
		text  string
	}{
		{
			"32:22",
			"32:22",
			32*time.Minute + 22*time.Second,
			"32:22",
		},
		{
			"20:53:37",
			"20:53:37",
			20*time.Hour + 53*time.Minute + 37*time.Second,
			"20:53:37",
		},
		{
			"2-03:06:23",
			"2-03:06:23",
			2*24*time.Hour + 3*time.Hour + 6*time.Minute + 23*time.Second,
			"2-03:06:23",
		},
		{
			"0:00",
			"0:00",
			0,
			"0:00",
		},
	}

	for _, test := range tests {
		var p hpcmodel.TimeStringProperty
		p.SetValue(test.s)
		if p.Value != test.value {
			t.Errorf("p.Value != %s", test.value)
		}
		if p.Data != test.data {
			t.Errorf("p.Data != %v", test.data)
		}
		if p.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}

func TestBuildProperty_NumberProperty(t *testing.T) {
	var tests = []struct {
		category hpcmodel.QueryCategory
		records  []string
		value    string
		data     float64
		text     string
	}{
		{
			hpcmodel.QueryCategory{
				ID:          "id1",
				DisplayName: "owner",
				Label:       "owner",
				ParseRecord: &hpcmodel.TokenRecordParser{
					Index: 2,
					Sep:   "|",
				},
				PropertyClass: "NumberProperty",
			},
			[]string{"nwp_xp|(null)|1|0|NONE"},
			"1",
			1,
			"1",
		},
	}
	for _, test := range tests {
		p, err := hpcmodel.BuildProperty(test.records, test.category)
		if err != nil {
			t.Errorf("build property failed: %v", err)
		}
		np, ok := p.(*hpcmodel.NumberProperty)
		if !ok {
			t.Errorf("property is not NumberProperty")
		}
		if np.Value != test.value {
			t.Errorf("p.Text != %s", test.value)
		}
		if np.Data != test.data {
			t.Errorf("p.Data != %f", test.data)
		}
		if np.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}

	}
}

func TestBuildProperty_TimeStringProperty(t *testing.T) {
	var tests = []struct {
		category hpcmodel.QueryCategory
		records  []string
		value    string
		data     time.Duration
		text     string
	}{
		{
			hpcmodel.QueryCategory{
				ID:          "id1",
				DisplayName: "time",
				Label:       "time",
				ParseRecord: &hpcmodel.TokenRecordParser{
					Index: 1,
					Sep:   "|",
				},
				PropertyClass: "TimeStringProperty",
			},
			[]string{"job|20:53:37|running"},
			"20:53:37",
			20*time.Hour + 53*time.Minute + 37*time.Second,
			"20:53:37",
		},
		{
			hpcmodel.QueryCategory{
				ID:          "id2",
				DisplayName: "time_left",
				Label:       "time_left",
				ParseRecord: &hpcmodel.TokenRecordParser{
					Index: 0,
					Sep:   "|",
				},
				PropertyClass: "TimeStringProperty",
			},
			[]string{"2-03:06:23"},
			"2-03:06:23",
			2*24*time.Hour + 3*time.Hour + 6*time.Minute + 23*time.Second,
			"2-03:06:23",
		},
	}
	for _, test := range tests {
		p, err := hpcmodel.BuildProperty(test.records, test.category)
		if err != nil {
			t.Errorf("build property failed: %v", err)
		}
		tp, ok := p.(*hpcmodel.TimeStringProperty)
		if !ok {
			t.Errorf("property is not TimeStringProperty")
		}
		if tp.Value != test.value {
			t.Errorf("p.Value != %s", test.value)
		}
		if tp.Data != test.data {
			t.Errorf("p.Data != %v", test.data)
		}
		if tp.Text != test.text {
			t.Errorf("p.Text != %s", test.text)
		}
	}
}
