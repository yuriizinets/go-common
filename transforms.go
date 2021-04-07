package common

import (
	"encoding/json"
	"io/ioutil"
)

// JSONDumps is a helper to create json string from Go variable.
// Throws panic on error
func JSONDumps(val interface{}) string {
	databytes, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return string(databytes)
}

// JSONDump is a helper to create json file from Go variable.
// Throws panic on error
func JSONDump(val interface{}, filepath string) {
	databytes, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath, databytes, 0644)
	if err != nil {
		panic(err)
	}
}

// JSONDumps is a helper to create map[string]interface{} from string.
// Throws panic on error
func JSONLoads(val string, target interface{}) {
	err := json.Unmarshal([]byte(val), target)
	if err != nil {
		panic(err)
	}
}

// JSONLoadsMap is a helper to create map[string]interface{} from string.
// Throws panic on error
func JSONLoadsMap(val string) map[string]interface{} {
	var target map[string]interface{}
	err := json.Unmarshal([]byte(val), &target)
	if err != nil {
		panic(err)
	}
	return target
}

// JSONLoad is a helper to create map[string]interface{} from file content.
// Throws panic on error
func JSONLoad(filepath string, target interface{}) {
	databytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(databytes, target)
	if err != nil {
		panic(err)
	}
}

// JSONLoadMap is a helper to create map[string]interface{} from file content.
// Throws panic on error
func JSONLoadMap(filepath string) map[string]interface{} {
	databytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var target map[string]interface{}
	err = json.Unmarshal(databytes, &target)
	if err != nil {
		panic(err)
	}
	return target
}

// CastISliceToInt is a helper to cast []interface{} to []int
func CastISliceToInt(val []interface{}) []int {
	_val := []int{}
	for _, v := range val {
		_val = append(_val, v.(int))
	}
	return _val
}

// CastISliceToIntPointers is a helper to cast []interface{} to []*int
func CastISliceToIntPointers(val []interface{}) []*int {
	_val := []*int{}
	for _, v := range val {
		_val = append(_val, v.(*int))
	}
	return _val
}

// CastISliceToFloat32 is a helper to cast []interface{} to []float32
func CastISliceToFloat32(val []interface{}) []float32 {
	_val := []float32{}
	for _, v := range val {
		_val = append(_val, v.(float32))
	}
	return _val
}

// CastISliceToFloat32Pointers is a helper to cast []interface{} to []*float32
func CastISliceToFloat32Pointers(val []interface{}) []*float32 {
	_val := []*float32{}
	for _, v := range val {
		_val = append(_val, v.(*float32))
	}
	return _val
}

// CastISliceToFloat64 is a helper to cast []interface{} to []float64
func CastISliceToFloat64(val []interface{}) []float64 {
	_val := []float64{}
	for _, v := range val {
		_val = append(_val, v.(float64))
	}
	return _val
}

// CastISliceToFloat64Pointers is a helper to cast []interface{} to []*float64
func CastISliceToFloat64Pointers(val []interface{}) []*float64 {
	_val := []*float64{}
	for _, v := range val {
		_val = append(_val, v.(*float64))
	}
	return _val
}

// CastISliceToString is a helper to cast []interface{} to []string
func CastISliceToString(val []interface{}) []string {
	_val := []string{}
	for _, v := range val {
		_val = append(_val, v.(string))
	}
	return _val
}

// CastISliceToStringPointers is a helper to cast []interface{} to []*string
func CastISliceToStringPointers(val []interface{}) []*string {
	_val := []*string{}
	for _, v := range val {
		_val = append(_val, v.(*string))
	}
	return _val
}
