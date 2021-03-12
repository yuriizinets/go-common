package common

import (
	"html/template"
)

// TFMAttach is a function to apply all function subsets to the provided funcmap
func TFMAttach(t *template.FuncMap) {
	TFMAttachOperations(t)
	TFMAttachDates(t)
	TFMAttachOther(t)
}

// TFMAttachOperations is a function to apply operations functions to the provided funcmap
func TFMAttachOperations(t *template.FuncMap) {
	(*t)["sum"] = Sum
	(*t)["sub"] = Sub
	(*t)["mult"] = Mult
	(*t)["div"] = Div
}

// TFMAttachTransforms is a function to apply transform functions to the provided funcmap
func TFMAttachTransforms(t *template.FuncMap) {
	(*t)["jsonloadsmap"] = JSONLoadsMap
	(*t)["jsondumps"] = JSONDumps
}

// TFMAttachDates is a function to apply dates functions to the provided funcmap
func TFMAttachDates(t *template.FuncMap) {
	(*t)["currentyear"] = CurrentYear
	(*t)["currentmonth"] = CurrentMonth
	(*t)["currentday"] = CurrentDay
	(*t)["currentdate"] = CurrentDateStr
}

// TFMAttachOther is a function to apply minor functions subset to the provided funcmap
func TFMAttachOther(t *template.FuncMap) {
	(*t)["rng"] = Range
}
