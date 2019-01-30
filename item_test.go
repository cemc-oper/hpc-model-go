package hpcmodel_test

import (
	"nwpc-hpc-model-go"
	"testing"
)

func TestItem(t *testing.T) {
	var item hpcmodel.Item

	var prop hpcmodel.StringProperty
	prop.SetValue("wdp")

	item.Props = append(item.Props, &prop)
}
