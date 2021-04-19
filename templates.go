package common

import (
	"html/template"
	"os"
)

// TFMAttach is a function to apply all function subsets to the provided funcmap
func TFMAttach(t *template.FuncMap) {
	TFMAttachOperations(t)
	TFMAttachSlices(t)
	TFMAttachTransforms(t)
	TFMAttachDates(t)
	TFMAttachEscapes(t)
	TFMAttachOther(t)
}

// TFMAttachOperations is a function to apply operations functions to the provided funcmap
func TFMAttachOperations(t *template.FuncMap) {
	(*t)["sum"] = Sum
	(*t)["sub"] = Sub
	(*t)["mult"] = Mult
	(*t)["div"] = Div
}

func TFMAttachSlices(t *template.FuncMap) {
	(*t)["min"] = Min
	(*t)["max"] = Max
	(*t)["avg"] = Avg
	(*t)["in"] = In
	(*t)["remove"] = Remove
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

// TFMAttachEscapes is a function to apply escape functions to the provided funcmap
func TFMAttachEscapes(t *template.FuncMap) {
	(*t)["html"] = func(src string) template.HTML { return template.HTML(src) }
	(*t)["htmlattr"] = func(src string) template.HTMLAttr { return template.HTMLAttr(src) }
	(*t)["js"] = func(src string) template.JS { return template.JS(src) }
	(*t)["css"] = func(src string) template.CSS { return template.CSS(src) }
	(*t)["url"] = func(src string) template.URL { return template.URL(src) }
}

// TFMAttachOther is a function to apply minor functions to the provided funcmap
func TFMAttachOther(t *template.FuncMap) {
	(*t)["rng"] = Range
	(*t)["env"] = os.Getenv
}
