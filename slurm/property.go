package slurm

import hpcmodel "github.com/cemc-oper/hpc-model-go"

type JobState string

type JobStateProperty struct {
	Category hpcmodel.QueryCategory
	Value    string
	Text     string
	Data     JobState
}

func (p *JobStateProperty) SetValue(value string) {
	p.Value = value
	p.Text = value
	p.Data = JobState(value)
}

func (p *JobStateProperty) SetCategory(category hpcmodel.QueryCategory) {
	p.Category = category
}

func (p *JobStateProperty) PropertyID() string {
	return p.Category.ID
}
