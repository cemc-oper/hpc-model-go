package hpcmodel_test

import (
	"testing"
	"time"

	hpcmodel "github.com/cemc-oper/hpc-model-go"
)

func TestStringProperty_SetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "text", input: "nwpc"},
		{name: "empty", input: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p hpcmodel.StringProperty
			p.SetValue(tt.input)

			if p.Value != tt.input {
				t.Errorf("Value = %q, want %q", p.Value, tt.input)
			}
			if p.Data != tt.input {
				t.Errorf("Data = %q, want %q", p.Data, tt.input)
			}
			if p.Text != tt.input {
				t.Errorf("Text = %q, want %q", p.Text, tt.input)
			}
		})
	}
}

func TestNumberProperty_SetValue(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string
		wantText  string
		wantData  float64
	}{
		{name: "integer", input: "1", wantValue: "1", wantText: "1", wantData: 1},
		{name: "decimal", input: "1.25", wantValue: "1.25", wantText: "1.25", wantData: 1.25},
		{name: "invalid", input: "not-a-number", wantValue: "not-a-number", wantText: "not-a-number", wantData: 0},
		{name: "empty", input: "", wantValue: "", wantText: "", wantData: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p hpcmodel.NumberProperty
			p.SetValue(tt.input)

			if p.Value != tt.wantValue {
				t.Errorf("Value = %q, want %q", p.Value, tt.wantValue)
			}
			if p.Data != tt.wantData {
				t.Errorf("Data = %v, want %v", p.Data, tt.wantData)
			}
			if p.Text != tt.wantText {
				t.Errorf("Text = %q, want %q", p.Text, tt.wantText)
			}
		})
	}
}

func TestDateTimeProperty_SetValue(t *testing.T) {
	validTime := time.Date(2019, time.January, 29, 19, 45, 0, 0, time.UTC)
	tests := []struct {
		name      string
		input     string
		wantValue string
		wantText  string
		wantData  time.Time
	}{
		{
			name:      "valid",
			input:     "2019-01-29T19:45:00",
			wantValue: "2019-01-29T19:45:00",
			wantText:  "2019-01-29 19:45:00",
			wantData:  validTime,
		},
		{name: "not available", input: "N/A", wantValue: "N/A", wantText: "N/A", wantData: time.Time{}},
		{name: "empty", input: "", wantValue: "", wantText: "", wantData: time.Time{}},
		{name: "invalid", input: "not-a-time", wantValue: "not-a-time", wantText: "not-a-time", wantData: time.Time{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p hpcmodel.DateTimeProperty
			p.SetValue(tt.input)

			if p.Value != tt.wantValue {
				t.Errorf("Value = %q, want %q", p.Value, tt.wantValue)
			}
			if !p.Data.Equal(tt.wantData) {
				t.Errorf("Data = %v, want %v", p.Data, tt.wantData)
			}
			if p.Text != tt.wantText {
				t.Errorf("Text = %q, want %q", p.Text, tt.wantText)
			}
		})
	}
}

func TestTimestampProperty_SetValue(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string
		wantText  string
		wantData  time.Time
	}{
		{
			name:      "UTC timestamp",
			input:     "1547906480",
			wantValue: "1547906480",
			wantText:  "2019-01-19 14:01:20",
			wantData:  time.Date(2019, time.January, 19, 14, 1, 20, 0, time.UTC),
		},
		{
			name:      "epoch",
			input:     "0",
			wantValue: "0",
			wantText:  "1970-01-01 00:00:00",
			wantData:  time.Unix(0, 0).UTC(),
		},
		{name: "invalid", input: "invalid", wantValue: "invalid", wantText: "invalid", wantData: time.Time{}},
		{name: "empty", input: "", wantValue: "", wantText: "", wantData: time.Time{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p hpcmodel.TimestampProperty
			p.SetValue(tt.input)

			if p.Value != tt.wantValue {
				t.Errorf("Value = %q, want %q", p.Value, tt.wantValue)
			}
			if !p.Data.Equal(tt.wantData) {
				t.Errorf("Data = %v, want %v", p.Data, tt.wantData)
			}
			if p.Data.Location() != time.UTC {
				t.Errorf("Data location = %v, want UTC", p.Data.Location())
			}
			if p.Text != tt.wantText {
				t.Errorf("Text = %q, want %q", p.Text, tt.wantText)
			}
		})
	}
}

func TestTimeStringProperty_SetValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantData time.Duration
	}{
		{name: "minutes and seconds", input: "32:22", wantData: 32*time.Minute + 22*time.Second},
		{name: "hours minutes and seconds", input: "20:53:37", wantData: 20*time.Hour + 53*time.Minute + 37*time.Second},
		{name: "days hours minutes and seconds", input: "2-03:06:23", wantData: 2*24*time.Hour + 3*time.Hour + 6*time.Minute + 23*time.Second},
		{name: "zero", input: "0:00", wantData: 0},
		{name: "not available", input: "N/A", wantData: 0},
		{name: "empty", input: "", wantData: 0},
		{name: "invalid marker", input: "INVALID", wantData: 0},
		{name: "too many segments", input: "1:02:03:04", wantData: 0},
		{name: "invalid numbers", input: "ab:cd", wantData: 0},
		{name: "invalid days", input: "x-01:00:00", wantData: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p hpcmodel.TimeStringProperty
			p.SetValue(tt.input)

			if p.Value != tt.input {
				t.Errorf("Value = %q, want %q", p.Value, tt.input)
			}
			if p.Data != tt.wantData {
				t.Errorf("Data = %v, want %v", p.Data, tt.wantData)
			}
			if p.Text != tt.input {
				t.Errorf("Text = %q, want %q", p.Text, tt.input)
			}
		})
	}
}

func TestBuildProperty_StringProperty(t *testing.T) {
	tests := []struct {
		name      string
		category  hpcmodel.QueryCategory
		records   []string
		wantValue string
		wantData  string
		wantText  string
	}{
		{
			name: "string from token",
			category: hpcmodel.QueryCategory{
				ID:            "id1",
				DisplayName:   "owner",
				Label:         "owner",
				ParseRecord:   &hpcmodel.TokenRecordParser{Index: 0, Sep: "|"},
				PropertyClass: "StringProperty",
			},
			records:   []string{"nwp_xp|(null)|1|0|NONE"},
			wantValue: "nwp_xp",
			wantData:  "nwp_xp",
			wantText:  "nwp_xp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property, err := hpcmodel.BuildProperty(tt.records, tt.category)
			if err != nil {
				t.Fatalf("BuildProperty() error = %v", err)
			}
			p, ok := property.(*hpcmodel.StringProperty)
			if !ok {
				t.Fatalf("BuildProperty() type = %T, want *StringProperty", property)
			}
			if p.Value != tt.wantValue {
				t.Errorf("Value = %q, want %q", p.Value, tt.wantValue)
			}
			if p.Data != tt.wantData {
				t.Errorf("Data = %q, want %q", p.Data, tt.wantData)
			}
			if p.Text != tt.wantText {
				t.Errorf("Text = %q, want %q", p.Text, tt.wantText)
			}
		})
	}
}

func TestBuildProperty_NumberProperty(t *testing.T) {
	tests := []struct {
		name      string
		category  hpcmodel.QueryCategory
		records   []string
		wantValue string
		wantData  float64
		wantText  string
	}{
		{
			name: "number from token",
			category: hpcmodel.QueryCategory{
				ID:            "id1",
				DisplayName:   "owner",
				Label:         "owner",
				ParseRecord:   &hpcmodel.TokenRecordParser{Index: 2, Sep: "|"},
				PropertyClass: "NumberProperty",
			},
			records:   []string{"nwp_xp|(null)|1|0|NONE"},
			wantValue: "1",
			wantData:  1,
			wantText:  "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property, err := hpcmodel.BuildProperty(tt.records, tt.category)
			if err != nil {
				t.Fatalf("BuildProperty() error = %v", err)
			}
			p, ok := property.(*hpcmodel.NumberProperty)
			if !ok {
				t.Fatalf("BuildProperty() type = %T, want *NumberProperty", property)
			}
			if p.Value != tt.wantValue {
				t.Errorf("Value = %q, want %q", p.Value, tt.wantValue)
			}
			if p.Data != tt.wantData {
				t.Errorf("Data = %v, want %v", p.Data, tt.wantData)
			}
			if p.Text != tt.wantText {
				t.Errorf("Text = %q, want %q", p.Text, tt.wantText)
			}
		})
	}
}

func TestBuildProperty_TimeStringProperty(t *testing.T) {
	tests := []struct {
		name      string
		category  hpcmodel.QueryCategory
		records   []string
		wantValue string
		wantData  time.Duration
		wantText  string
	}{
		{
			name: "hours minutes and seconds",
			category: hpcmodel.QueryCategory{
				ID:            "id1",
				DisplayName:   "time",
				Label:         "time",
				ParseRecord:   &hpcmodel.TokenRecordParser{Index: 1, Sep: "|"},
				PropertyClass: "TimeStringProperty",
			},
			records:   []string{"job|20:53:37|running"},
			wantValue: "20:53:37",
			wantData:  20*time.Hour + 53*time.Minute + 37*time.Second,
			wantText:  "20:53:37",
		},
		{
			name: "days hours minutes and seconds",
			category: hpcmodel.QueryCategory{
				ID:            "id2",
				DisplayName:   "time_left",
				Label:         "time_left",
				ParseRecord:   &hpcmodel.TokenRecordParser{Index: 0, Sep: "|"},
				PropertyClass: "TimeStringProperty",
			},
			records:   []string{"2-03:06:23"},
			wantValue: "2-03:06:23",
			wantData:  2*24*time.Hour + 3*time.Hour + 6*time.Minute + 23*time.Second,
			wantText:  "2-03:06:23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			property, err := hpcmodel.BuildProperty(tt.records, tt.category)
			if err != nil {
				t.Fatalf("BuildProperty() error = %v", err)
			}
			p, ok := property.(*hpcmodel.TimeStringProperty)
			if !ok {
				t.Fatalf("BuildProperty() type = %T, want *TimeStringProperty", property)
			}
			if p.Value != tt.wantValue {
				t.Errorf("Value = %q, want %q", p.Value, tt.wantValue)
			}
			if p.Data != tt.wantData {
				t.Errorf("Data = %v, want %v", p.Data, tt.wantData)
			}
			if p.Text != tt.wantText {
				t.Errorf("Text = %q, want %q", p.Text, tt.wantText)
			}
		})
	}
}
