package gtype

// New is alias of NewInterface.
// See NewInterface.
func New(value ...interface{}) *Interface {
	return NewInterface(value...)
}
