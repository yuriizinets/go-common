package common

import (
	"html/template"
)

// TFMApply is a function to apply all function subsets to the provided funcmap
func TFMApply(t *template.FuncMap) {
	TFMApplyOperations(t)
	TFMApplyDates(t)
	TFMApplyOther(t)
}

// TFMApplyOperations is a function to apply operations subset to the provided funcmap
func TFMApplyOperations(t *template.FuncMap) {
	(*t)["sum"] = Sum
	(*t)["sub"] = Sub
	(*t)["mult"] = Mult
	(*t)["div"] = Div
}

func TFMApplyDates(t *template.FuncMap) {
	(*t)["currentyear"] = CurrentYear
	(*t)["currentmonth"] = CurrentMonth
	(*t)["currentday"] = CurrentDay
	(*t)["currentdate"] = CurrentDateStr
}

// TFMApplyOther is a function to apply minor functions subset to the provided funcmap
func TFMApplyOther(t *template.FuncMap) {
	(*t)["range"] = Range
}
