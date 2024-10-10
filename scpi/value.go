package scpi

import "fmt"

type Value float64

type valueArray struct {
	values []Value
}

func initValueArray() valueArray {
	return valueArray{
		values: make([]Value, 0),
	}
}

func (va *valueArray) writeValue(value Value) {
	va.values = append(va.values, value)
}

func printValue(value Value) {
	fmt.Printf("%g", value)
}
