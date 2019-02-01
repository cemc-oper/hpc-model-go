package hpcmodel

type QueryCategory struct {
	Id                      string
	DisplayName             string
	Label                   string
	ParseRecord             RecordParser
	PropertyClass           string
	PropertyCreateArguments []string
}

type QueryCategoryList struct {
	CategoryList []QueryCategory
}
