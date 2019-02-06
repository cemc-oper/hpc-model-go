package hpcmodel

type Filter struct {
	Conditions []FilterCondition
}

func (f *Filter) Filter(items []Item) []Item {
	var targetItems []Item
	for _, item := range items {
		isFit := true
		for _, condition := range f.Conditions {
			if !condition.IsFit(&item) {
				isFit = false
				break
			}
		}
		if isFit {
			targetItems = append(targetItems, item)
		}
	}
	return targetItems
}
