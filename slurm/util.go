package slurm

import (
	hpcmodel "github.com/cemc-oper/hpc-model-go"
)

func CreatePropertyLessFunc(id string) hpcmodel.LessFunc {
	return func(item1, item2 *hpcmodel.Item) bool {
		p1 := item1.GetProperty(id)
		p2 := item2.GetProperty(id)
		switch prop1 := p1.(type) {
		case *JobStateProperty:
			prop2 := p2.(*JobStateProperty)
			return prop1.Data < prop2.Data
		default:
			return hpcmodel.CreatePropertyLessFunc(id)(item1, item2)
		}
	}
}
