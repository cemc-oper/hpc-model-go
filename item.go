package hpcmodel

type Item struct {
	Props []Property
}

func (item *Item) AddProp(p Property) {
	item.Props = append(item.Props, p)
}

func (item *Item) GetProperty(propertyID string) Property {
	for _, prop := range item.Props {
		propID := prop.PropertyID()
		if propID != propertyID {
			continue
		}
		return prop
	}
	return nil
}

func BuildItem(records []string, categoryList QueryCategoryList) *Item {
	item := new(Item)
	var p Property
	for _, category := range categoryList.CategoryList {
		switch category.PropertyClass {
		case "StringProperty":
			p = &StringProperty{}
		case "NumberProperty":
			p = &NumberProperty{}
		case "DateTimeProperty":
			p = &DateTimeProperty{}
		case "TimestampProperty":
			p = &TimestampProperty{}
		}
		BuildProperty(p, records, category)
		item.AddProp(p)
	}
	return item
}
