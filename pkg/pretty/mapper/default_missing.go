package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/mapper"
)

type DefaultMissing struct {
	Field   string
	Default interface{}
}

func (d DefaultMissing) FromInternal(data map[string]interface{}) {
}

func (d DefaultMissing) ToInternal(data map[string]interface{}) error {
	if _, ok := data[d.Field]; !ok && data != nil {
		data[d.Field] = d.Default
	}

	return nil
}

func (d DefaultMissing) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return mapper.ValidateField(d.Field, schema)
}
