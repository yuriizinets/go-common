
# GO Common

Common and missing Go things/quirks/features, like operations with slices, ranges, data transformations and so on.  

## Features map

- [Operations](Operations)
- [Transforms](Transforms)
- [Dates](Dates)
- [Other](Other)

## Import

```go
import(
    ...
    common "github.com/yuriizinets/go-common"
)
```

## Operations

Most common operations, like `Sum`, `Sub`, `Div`, `Mult`. Mainly created for using inside of templates, where such basic things doesn't exist.  

Supported types: `int`,`float32`,`float64`,`string(partialy)`

Examples:  

```go
// Sum
var ressum int
var ressumstr string
ressum = Sum(10, 2).(int) // 12
ressum = Sum(10, 2, 3, 3).(int) // 18. You can add multiple items
ressumstr = Sum("Foo", "Bar").(string) // FooBar. Works are usual string concat

// Sub
var ressub int
var ressubstr string
ressub = Sub(10, 2, 3).(int) // 5
ressubstr = Sub("FooBar", "Bar").(string) // Foo. Works as ReplaceAll

// Mult
var resmult int
var resmultstr string
resmult = Mult(10, 2, 3).(int) // 60
resmultstr = Mult("Foo", 2) // FooFoo. Multiplies string, as in Python

// Div
var resdiv int
resdiv = Div(10, 2, 2).(int) // 2. Use floats for accurate results


// Using in templates

// First, attach operations funcs to the funcmap
funcmap := template.FuncMap{}
common.TFMAttachOperations(funcmap)
// And use them in the template like here
var template = `
 {{ sum 10 2 3 3 }}
 {{ sub 10 2 3 }}
 {{ mult 10 2 3 }}
 {{ div 10 2 }}
`
...
```

## Transforms

Data transformations, like JSON to string and back.  

Examples:

```go
// Test struct
type Person struct {
    Name string
    Age string
    AdditionalData []string
}

// Test data
var wdata = Person{
    "Name": "John",
    "Age": 35,
    "AdditionalData": []string{"data1", "data2"}
}
var rawdata = `{"Name":"John","Age":35,"AdditionalData":["data1","data2"]}`
var rdata Person
var rrdata map[string]interface{}

// Stringify
datastr := common.JSONDumps(wdata) // {"Name":"John",...}
// Safe to file
common.JSONDump(wdata, "filename.json")
// Load from string. You can use any struct as target
common.JSONLoads(rawdata, &rdata)
// Load from string into map, without declaration
rrdata = common.JSONLoadsMap(rawdata) // map[string]interface{}
// Load from file. You can use any struct as target
common.JSONLoad("filename.json", &rdata)
// Load from file into map, without declaration
rrdata = common.JSONLoadMap("filename.json") // map[string]interface{}

// Using in templates

// First, attach transforms funcs to the funcmap
funcmap := template.FuncMap{}
common.TFMAttachTransforms(funcmap)
// And use them in the template like here
var template = `
 {{ variable | jsondumps }}
 {{ (jsonloadsmap "filename.json") | jsondumps }}
`
```

## Dates

Most common dates operations, like `CurrentYear`, `CurrentDay`, etc. Mainly created for using inside of templates, where such basic things doesn't exist.  

Examples:

```go
tz := 0 // Timezone shift (in hours)
year := CurrentYear(tz) // f.e. 2021
month := CurrentMonth(tz) // f.e. 3
day := CurrentDay(tz) // f.e. 12
datestr := CurrentDateStr(tz, "") // f.e. 2021-03-12. Use second argument for custom format

// Using in templates

// First, attach date funcs to the funcmap
funcmap := template.FuncMap{}
common.TFMAttachDates(funcmap)
// And use them in the template like here
var template = `
 {{ currentyear 0 }}-{{ currentmonth 0 }}-{{ currentday 0 }}
 {{ currentdate 0 "" }}
`
```

## Other

Uncategorized functions. Check sources or examples for details.  

Examples:

```go
rng := Range(1, 5) // []int{1, 2, 3, 4}

// Using in templates

// First, attach other funcs to the funcmap
funcmap := template.FuncMap{}
common.TFMAttachOther(funcmap)
// And use them in the template like here
var template = `
 {{ range (rng 1 5) }}
   {{ . }}
 {{ end }}
`
```
