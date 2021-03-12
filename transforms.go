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

// JSONDumps is a helper to create map[string]interface{} from string.
// Throws panic on error
func JSONLoadsMap(val string) map[string]interface{} {
	var target map[string]interface{}
	err := json.Unmarshal([]byte(val), &target)
	if err != nil {
		panic(err)
	}
	return target
}

// JSONDumps is a helper to create map[string]interface{} from file content.
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

// JSONDumps is a helper to create map[string]interface{} from file content.
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
