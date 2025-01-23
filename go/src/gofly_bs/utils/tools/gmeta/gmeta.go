package gmeta

import (
	"gofly/utils/tools/gstructs"
	"gofly/utils/tools/gvar"
)

// Meta is used as an embedded attribute for struct to enabled metadata feature.
type Meta struct{}

const (
	metaAttributeName = "Meta"       // metaAttributeName is the attribute name of metadata in struct.
	metaTypeName      = "gmeta.Meta" // metaTypeName is for type string comparison.
)

// Data retrieves and returns all metadata from `object`.
func Data(object interface{}) map[string]string {
	reflectType, err := gstructs.StructType(object)
	if err != nil {
		return nil
	}
	if field, ok := reflectType.FieldByName(metaAttributeName); ok {
		if field.Type.String() == metaTypeName {
			return gstructs.ParseTag(string(field.Tag))
		}
	}
	return map[string]string{}
}

// Get retrieves and returns specified metadata by `key` from `object`.
func Get(object interface{}, key string) *gvar.Var {
	v, ok := Data(object)[key]
	if !ok {
		return nil
	}
	return gvar.New(v)
}
