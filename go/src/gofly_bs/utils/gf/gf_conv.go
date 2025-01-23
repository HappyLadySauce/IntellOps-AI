package gf

import (
	"gofly/utils/tools/gconv"
)

// any数据类型转成int
func Int(i interface{}) int {
	return gconv.Int(i)
}

// any数据类型转成int8
func Int8(i interface{}) int8 {
	return gconv.Int8(i)
}

// any数据类型转成int16
func Int16(i interface{}) int16 {
	return gconv.Int16(i)
}

// any数据类型转成int32
func Int32(i interface{}) int32 {
	return gconv.Int32(i)
}

// any数据类型转成int64
func Int64(i interface{}) int64 {
	return gconv.Int64(i)
}

// any数据类型转成uint
func Uint(i interface{}) uint {
	return gconv.Uint(i)
}

// any数据类型转成uint8
func Uint8(i interface{}) uint8 {
	return gconv.Uint8(i)
}

// any数据类型转成uint16
func Uint16(i interface{}) uint16 {
	return gconv.Uint16(i)
}

// any数据类型转成uint32
func Uint32(i interface{}) uint32 {
	return gconv.Uint32(i)
}

// any数据类型转成uint64
func Uint64(i interface{}) uint64 {
	return gconv.Uint64(i)
}

// any数据类型转成float32
func Float32(i interface{}) float32 {
	return gconv.Float32(i)
}

// any数据类型转成float64
func Float64(i interface{}) float64 {
	return gconv.Float64(i)
}

// any数据类型转成bool
func Bool(i interface{}) bool {
	return gconv.Bool(i)
}

// any数据类型转成string
func String(i interface{}) string {
	return gconv.String(i)
}

// any数据类型转成 []byte
func Bytes(i interface{}) []byte {
	return gconv.Bytes(i)
}

// any数据类型转成 []string 数组
func Strings(i interface{}) []string {
	return gconv.Strings(i)
}

// any数据类型转成 []int 数组
func Ints(i interface{}) []int {
	return gconv.Ints(i)
}

// any数据类型转成 []float64 数组
func Floats(i interface{}) []float64 {
	return gconv.Floats(i)
}

// any数据类型转成 []interface{}  数组
func Interfaces(i interface{}) []interface{} {
	return gconv.Interfaces(i)
}
