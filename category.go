package hpcmodel

type QueryCategory struct {
	Id          string
	DisplayName string
	Label       string
}

type QueryCategoryList struct {
	CategoryList []QueryCategory
}
