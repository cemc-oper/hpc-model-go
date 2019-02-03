package hpcmodel

type QueryCategory struct {
	ID                      string
	DisplayName             string
	Label                   string
	ParseRecord             RecordParser
	RecordParserClass       string
	RecordParserArguments   []string
	PropertyClass           string
	PropertyCreateArguments []string
}

func (q1 *QueryCategory) Equal(q2 *QueryCategory) bool {
	if q1 == nil && q2 == nil {
		return true
	}
	if q1.ID != q2.ID {
		return false
	}
	if q1.DisplayName != q2.DisplayName {
		return false
	}
	if q1.RecordParserClass != q2.RecordParserClass {
		return false
	}

	if len(q1.RecordParserArguments) != len(q2.RecordParserArguments) {
		return false
	}

	for i, arg := range q1.RecordParserArguments {
		if arg != q2.RecordParserArguments[i] {
			return false
		}
	}

	if q1.PropertyClass != q2.PropertyClass {
		return false
	}

	if len(q1.PropertyCreateArguments) != len(q2.PropertyCreateArguments) {
		return false
	}

	for i2, arg2 := range q1.PropertyCreateArguments {
		if arg2 != q2.PropertyCreateArguments[i2] {
			return false
		}
	}

	return true
}

type QueryCategoryList struct {
	CategoryList []QueryCategory
}

func (ql *QueryCategoryList) ContainsID(id string) bool {
	for _, category := range ql.CategoryList {
		if category.ID == id {
			return true
		}
	}
	return false
}

func (ql *QueryCategoryList) IndexFromId(id string) int {
	for index, category := range ql.CategoryList {
		if category.ID == id {
			return index
		}
	}
	return -1
}

func (ql *QueryCategoryList) CategoryFromId(id string) *QueryCategory {
	for _, category := range ql.CategoryList {
		if category.ID == id {
			return &category
		}
	}
	return nil
}

func (ql *QueryCategoryList) ContainsLabel(label string) bool {
	for _, category := range ql.CategoryList {
		if category.Label == label {
			return true
		}
	}
	return false
}

func (ql *QueryCategoryList) IndexFromLabel(label string) int {
	for index, category := range ql.CategoryList {
		if category.Label == label {
			return index
		}
	}
	return -1
}

func (ql *QueryCategoryList) CategoryFromLabel(label string) *QueryCategory {
	for _, category := range ql.CategoryList {
		if category.Label == label {
			return &category
		}
	}
	return nil
}
