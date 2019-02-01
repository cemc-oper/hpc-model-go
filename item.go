package hpcmodel

type Item struct {
	Props []Property
}

func BuildItem(item *Item, records []string, categoryList QueryCategoryList) {
	for _, category := range categoryList.CategoryList {
		switch category.PropertyClass {
		case "StringProperty":
			property := StringProperty{}
			BuildProperty(&property, records, category)
			item.Props = append(item.Props, &property)
		case "NumberProperty":
			property := NumberProperty{}
			BuildProperty(&property, records, category)
			item.Props = append(item.Props, &property)
		case "DateTimeProperty":
			property := DateTimeProperty{}
			BuildProperty(&property, records, category)
			item.Props = append(item.Props, &property)
		case "TimestampProperty":
			property := TimestampProperty{}
			BuildProperty(&property, records, category)
			item.Props = append(item.Props, &property)
		}
	}
}

func GetProperty(item *Item, propertyID string) Property {
	for _, prop := range item.Props {
		propID, err := GetPropertyID(prop)
		if err != nil {
			continue
		}
		if propID != propertyID {
			continue
		}
		return prop
	}
	return nil
}
