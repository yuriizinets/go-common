package common

// Min is a function to find minimal value from the list. Always use the same type for all parameters! Examples:
// Min(3, 5, 2, 5) 2
// Min(10.3, 2.3, 3.5) 2.3
// Min("Fooo", "Bar") "Bar" (by length)
func Min(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		min := vals[0].(int)
		for _, v := range vals {
			if v.(int) < min {
				min = v.(int)
			}
		}
		return min
	case float32:
		min := vals[0].(float32)
		for _, v := range vals {
			if v.(float32) < min {
				min = v.(float32)
			}
		}
		return min
	case float64:
		min := vals[0].(float64)
		for _, v := range vals {
			if v.(float64) < min {
				min = v.(float64)
			}
		}
		return min
	case string:
		min := vals[0].(string)
		for _, v := range vals {
			if len(v.(string)) < len(min) {
				min = v.(string)
			}
		}
		return min
	default:
		panic("Type is not supported")
	}
}

// Max is a function to find maximum value from the list. Always use the same type for all parameters! Examples:
// Max(3, 5, 2, 5) 5
// Max(10.3, 2.3, 3.5) 10.3
// Max("Fooo", "Bar") "Fooo" (by length)
func Max(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		min := vals[0].(int)
		for _, v := range vals {
			if v.(int) > min {
				min = v.(int)
			}
		}
		return min
	case float32:
		min := vals[0].(float32)
		for _, v := range vals {
			if v.(float32) > min {
				min = v.(float32)
			}
		}
		return min
	case float64:
		min := vals[0].(float64)
		for _, v := range vals {
			if v.(float64) > min {
				min = v.(float64)
			}
		}
		return min
	case string:
		min := vals[0].(string)
		for _, v := range vals {
			if len(v.(string)) > len(min) {
				min = v.(string)
			}
		}
		return min
	default:
		panic("Type is not supported")
	}
}

// Avg is a function to find average value from the list. Always use the same type for all parameters! Examples:
// Avg(3, 5, 2, 5) 3
// Avg(3.0, 5.0, 2.0, 5.0) 3.75
// Avg("Fooo", "Bar") 3.5 (float64, by length)
func Avg(vals ...interface{}) interface{} {
	switch vals[0].(type) {
	case int:
		sum := Sum(vals...).(int)
		return sum / len(vals)
	case float32:
		sum := Sum(vals...).(float32)
		return sum / float32(len(vals))
	case float64:
		sum := Sum(vals...).(float64)
		return sum / float64(len(vals))
	case string:
		lngs := []interface{}{}
		for _, val := range vals {
			lngs = append(lngs, float64(len(val.(string))))
		}
		return Avg(lngs...)
	default:
		panic("Type is not supported")
	}
}

// In is a function to check value is in a list. Always use the same type for all parameters! Examples:
// In("foo", []string{"foo", "bar"}) true
// In(3, []int{4, 5}) false
func In(val interface{}, vals interface{}) bool {
	switch val.(type) {
	case int:
		_vals := vals.([]int)
		for _, v := range _vals {
			if val.(int) == v {
				return true
			}
		}
		return false
	case float32:
		_vals := vals.([]float32)
		for _, v := range _vals {
			if val.(float32) == v {
				return true
			}
		}
		return false
	case float64:
		_vals := vals.([]float64)
		for _, v := range _vals {
			if val.(float64) == v {
				return true
			}
		}
		return false
	case string:
		_vals := vals.([]string)
		for _, v := range _vals {
			if val.(string) == v {
				return true
			}
		}
		return false
	default:
		panic("Type is not supported")
	}
}
