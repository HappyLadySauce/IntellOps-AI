package gvar

import (
	"gofly/utils/tools/gconv"
)

// Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer`
// with type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.
//
// See gconv.Scan.
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Scan(v.Val(), pointer, mapping...)
}
