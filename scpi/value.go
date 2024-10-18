package scpi

//import "fmt"

type valueType int

const (
	VAL_BOOL valueType = iota + 1
	VAL_NUMBER
	VAL_STRING
)

type value struct {
	vType   valueType
	boolval bool
	numval  float64
	strval  string
}

func newNumberValue(val float64) value {
	return value{
		vType:  VAL_NUMBER,
		numval: val,
	}
}

// type valueArray struct {
// 	values []Value
// }
//
// func initValueArray() valueArray {
// 	return valueArray{
// 		values: make([]Value, 0),
// 	}
// }
//
// func (va *valueArray) writeValue(value Value) {
// 	va.values = append(va.values, value)
// }
//
// func printValue(value Value) {
// 	fmt.Printf("%g", value)
// }
