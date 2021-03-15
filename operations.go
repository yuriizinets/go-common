package common

import "strings"

// Sum is a function for values sum. Examples:
// Sum(10, 2) = 12
// Sum(10, 2, 3) = 15
// Sum("Foo Bar", " Baz") = "Foo Bar Baz"
func Sum(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		sum := int(0)
		for _, val := range vals {
			sum += val.(int)
		}
		return sum
	case float32:
		sum := float32(0)
		for _, val := range vals {
			sum += val.(float32)
		}
		return sum
	case float64:
		sum := float64(0)
		for _, val := range vals {
			sum += val.(float64)
		}
		return sum
	case string:
		sum := string("")
		for _, val := range vals {
			sum += val.(string)
		}
		return sum
	default:
		panic("Type is not supported")
	}
}

// Sub is a function for values substraction. Examples:
// Sub(10, 2) = 8
// Sub(10, 2, 3) = 5
// Sub("Foo Bar Baz", "Bar ") = "Foo Baz"
func Sub(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		sub := int(vals[0].(int))
		for _, val := range vals[1:] {
			sub -= val.(int)
		}
		return sub
	case float32:
		sub := float32(vals[0].(float32))
		for _, val := range vals[1:] {
			sub -= val.(float32)
		}
		return sub
	case float64:
		sub := float64(vals[0].(float64))
		for _, val := range vals[1:] {
			sub -= val.(float64)
		}
		return sub
	case string:
		sub := string(vals[0].(string))
		for _, val := range vals[1:] {
			sub = strings.ReplaceAll(sub, val.(string), "")
		}
		return sub
	default:
		panic("Type is not supported")
	}
}

// Mult is a function to multiply values. Examples:
// Mult(10, 2) = 20
// Mult(10, 2, 3) = 60
// Mult("Foo", 2) = "FooFoo"
func Mult(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		sum := int(1)
		for _, val := range vals {
			sum *= val.(int)
		}
		return sum
	case float32:
		sum := float32(1)
		for _, val := range vals {
			sum *= val.(float32)
		}
		return sum
	case float64:
		sum := float64(1)
		for _, val := range vals {
			sum *= val.(float64)
		}
		return sum
	case string:
		sum := string("")
		for i := 0; i < vals[1].(int); i++ {
			sum += vals[0].(string)
		}
		return sum
	default:
		panic("Type is not supported")
	}
}

// Div is a function to divide values. Always use the same type for all parameters! Examples:
// Div(10, 2) 5
// Div(50, 10, 2) 2
// Div(50.0, 10.0, 2.0) 2.5
func Div(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		div := int(vals[0].(int))
		for _, val := range vals[1:] {
			div /= val.(int)
		}
		return div
	case float32:
		div := float32(vals[0].(float32))
		for _, val := range vals[1:] {
			div /= val.(float32)
		}
		return div
	case float64:
		div := float64(vals[0].(float64))
		for _, val := range vals[1:] {
			div /= val.(float64)
		}
		return div
	default:
		panic("Type is not supported")
	}
}
