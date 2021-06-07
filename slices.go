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
	case []int:
		arr := vals[0].([]int)
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		return min
	case []float32:
		arr := vals[0].([]float32)
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		return min
	case []float64:
		arr := vals[0].([]float64)
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		return min
	case []string:
		arr := vals[0].([]string)
		min := vals[0].(string)
		for _, v := range arr {
			if len(v) < len(min) {
				min = v
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
	case []int:
		arr := vals[0].([]int)
		min := arr[0]
		for _, v := range arr {
			if v > min {
				min = v
			}
		}
		return min
	case []float32:
		arr := vals[0].([]float32)
		min := arr[0]
		for _, v := range arr {
			if v > min {
				min = v
			}
		}
		return min
	case []float64:
		arr := vals[0].([]float64)
		min := arr[0]
		for _, v := range arr {
			if v > min {
				min = v
			}
		}
		return min
	case []string:
		arr := vals[0].([]string)
		min := vals[0].(string)
		for _, v := range arr {
			if len(v) > len(min) {
				min = v
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
	case []int:
		sum := Sum(vals[0]).(int)
		return sum / len(vals[0].([]int))
	case []float32:
		sum := Sum(vals[0]).(float32)
		return sum / float32(len(vals[0].([]float32)))
	case []float64:
		sum := Sum(vals[0]).(float64)
		return sum / float64(len(vals[0].([]float64)))
	case []string:
		arr := vals[0].([]string)
		lngs := []interface{}{}
		for _, val := range arr {
			lngs = append(lngs, float64(len(val)))
		}
		return Avg(lngs...)
	default:
		panic("Type is not supported")
	}
}

// In is a function to check value is in the list. Always use the same type for all parameters! Examples:
// In("foo", []string{"foo", "bar"}) true
// In(3, []int{4, 5}) false
func In(val interface{}, vals interface{}) bool {
	switch val.(type) {
	case int:
		_vals := vals.([]int)
		// Convert, if needed
		if _, ok := vals.([]int); ok {
			_vals = vals.([]int)
		} else if _, ok := vals.([]*int); ok {
			_vals = SUnpoint(vals.([]*int)).([]int)
		}
		for _, v := range _vals {
			if val.(int) == v {
				return true
			}
		}
		return false
	case float32:
		_vals := vals.([]float32)
		// Convert, if needed
		if _, ok := vals.([]float32); ok {
			_vals = vals.([]float32)
		} else if _, ok := vals.([]*float32); ok {
			_vals = SUnpoint(vals.([]*float32)).([]float32)
		}
		for _, v := range _vals {
			if val.(float32) == v {
				return true
			}
		}
		return false
	case float64:
		_vals := vals.([]float64)
		// Convert, if needed
		if _, ok := vals.([]float64); ok {
			_vals = vals.([]float64)
		} else if _, ok := vals.([]*float64); ok {
			_vals = SUnpoint(vals.([]*float64)).([]float64)
		}
		for _, v := range _vals {
			if val.(float64) == v {
				return true
			}
		}
		return false
	case string:
		_vals := []string{}
		// Convert, if needed
		if _, ok := vals.([]string); ok {
			_vals = vals.([]string)
		} else if _, ok := vals.([]*string); ok {
			_vals = SUnpoint(vals.([]*string)).([]string)
		}
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

// Remove is a function to remove value from the list. Always use the same type for all parameters! Examples:
// Remove("foo", []string{"foo", "bar"}) []string{"bar"}
// Remove(3, []int{4, 3, 5}) []int{4, 5}
func Remove(val interface{}, vals interface{}) interface{} {
	switch val.(type) {
	case int:
		newslice := []int{}
		for _, v := range vals.([]int) {
			if val.(int) != v {
				newslice = append(newslice, v)
			}
		}
		return newslice
	case float32:
		newslice := []float32{}
		for _, v := range vals.([]float32) {
			if val.(float32) != v {
				newslice = append(newslice, v)
			}
		}
		return newslice
	case float64:
		newslice := []float64{}
		for _, v := range vals.([]float64) {
			if val.(float64) != v {
				newslice = append(newslice, v)
			}
		}
		return newslice
	case string:
		newslice := []string{}
		for _, v := range vals.([]string) {
			if val.(string) != v {
				newslice = append(newslice, v)
			}
		}
		return newslice
	default:
		panic("Type is not supported")
	}
}

// Deduplicate is a function to remove duplicates from the list. Always use the same type for all parameters! Examples:
// Deduplicate([]string{"foo", "foo", "bar"}) []string{"foo", "bar"}
// Deduplicate([]int{4, 3, 3, 5}) []int{4, 3, 5}
func Deduplicate(vals interface{}) interface{} {
	switch vals.(type) {
	case []int:
		keys := map[int]bool{}
		list := []int{}
		for _, v := range vals.([]int) {
			if _, ok := keys[v]; !ok {
				keys[v] = true
				list = append(list, v)
			}
		}
		return list
	case []float32:
		keys := map[float32]bool{}
		list := []float32{}
		for _, v := range vals.([]float32) {
			if _, ok := keys[v]; !ok {
				keys[v] = true
				list = append(list, v)
			}
		}
		return list
	case []float64:
		keys := map[float64]bool{}
		list := []float64{}
		for _, v := range vals.([]float64) {
			if _, ok := keys[v]; !ok {
				keys[v] = true
				list = append(list, v)
			}
		}
		return list
	case []string:
		keys := map[string]bool{}
		list := []string{}
		for _, v := range vals.([]string) {
			if _, ok := keys[v]; !ok {
				keys[v] = true
				list = append(list, v)
			}
		}
		return list
	default:
		panic("Type is not supported")
	}
}

// SUnpoint is a function to convert slice of pointers to normal slice
// f.e. []*string -> []string
func SUnpoint(slice interface{}) interface{} {
	switch slice := slice.(type) {
	case []*int:
		newslice := []int{}
		for _, val := range slice {
			newslice = append(newslice, *val)
		}
		return newslice
	case []*float32:
		newslice := []float32{}
		for _, val := range slice {
			newslice = append(newslice, *val)
		}
		return newslice
	case []*float64:
		newslice := []float64{}
		for _, val := range slice {
			newslice = append(newslice, *val)
		}
		return newslice
	case []*string:
		newslice := []string{}
		for _, val := range slice {
			newslice = append(newslice, *val)
		}
		return newslice
	default:
		panic("Type is not supported")
	}
}
