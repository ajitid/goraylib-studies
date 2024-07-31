package ctrl

import (
	"encoding/json"
	"log"
)

type namesStruct struct {
	name  string
	cType string
}

var (
	IsDrawing = false // no need to protect with a mutex as it is only being read/set by SetFloat and alikes + by GUI and both of them work on main thread
	names     []namesStruct
	stringMap = make(map[string]*string)
	floatMap  = make(map[string]*float32)
	intMap    = make(map[string]*int32)
	boolMap   = make(map[string]*bool)
	// TODO: RGBMap
)

///////// common

func findInAllNames(nameToCheck string) (namesStruct, bool) {
	for _, v := range names {
		if nameToCheck == v.name {
			return v, true
		}
	}
	return namesStruct{}, false
}

func informUI(name string, value any, cType string) {
	b, err := json.Marshal(payload{Name: name, Value: value, CType: cType})
	if err != nil {
		log.Fatal(b)
	}
	// because we are doing `go ctrl.RunServer()` and SSE is a part of it, it is better to communicate data via a channel
	informUIChan <- b
}

///////// float

func SetFloat(name string, value *float32) float32 {
	existingValue, existing := findInAllNames(name)
	if existing && existingValue.cType != "float" {
		log.Fatal("can't set", name, "as it is already set with some other type")
	}

	if IsDrawing && existing {
		return *floatMap[name]
	}

	floatMap[name] = value
	names = append(names, namesStruct{name: name, cType: "float"})
	informUI(name, *value, "float")
	return *value
}

func SetFloatV(name string, value float32) float32 {
	return SetFloat(name, &value)
}

// only to to be used when we need to set value from UI to Go
func setFloatInternal(name string, value float32) {
	if floatMap[name] == nil {
		return
	}
	*(floatMap[name]) = value
}

func GetFloat(name string) float32 {
	return *floatMap[name]
}

func GetFloatPtr(name string) *float32 {
	return floatMap[name]
}

///////// int

func SetInt(name string, value *int32) int32 {
	existingValue, existing := findInAllNames(name)
	if existing && existingValue.cType != "int" {
		log.Fatal("can't set", name, "as it is already set with some other type")
	}

	if IsDrawing && existing {
		return *intMap[name]
	}

	intMap[name] = value
	names = append(names, namesStruct{name: name, cType: "int"})
	informUI(name, *value, "int")
	return *value
}

func SetIntV(name string, value int32) int32 {
	return SetInt(name, &value)
}

// only to to be used when we need to set value from UI to Go
func setIntInternal(name string, value int32) {
	if intMap[name] == nil {
		return
	}
	*(intMap[name]) = value
}

func GetInt(name string) int32 {
	return *intMap[name]
}

func GetIntPtr(name string) *int32 {
	return intMap[name]
}

///////// string

func SetString(name string, value *string) string {
	existingValue, existing := findInAllNames(name)
	if existing && existingValue.cType != "string" {
		log.Fatal("can't set", name, "as it is already set with some other type")
	}

	if IsDrawing && existing {
		return *stringMap[name]
	}

	stringMap[name] = value
	names = append(names, namesStruct{name: name, cType: "string"})
	informUI(name, *value, "string")
	return *value
}

func SetStringV(name string, value string) string {
	return SetString(name, &value)
}

// only to to be used when we need to set value from UI to Go
func setStringInternal(name string, value string) {
	if stringMap[name] == nil {
		return
	}
	*(stringMap[name]) = value
}

func GetString(name string) string {
	return *stringMap[name]
}

func GetStringPtr(name string) *string {
	return stringMap[name]
}

///////// bool

func SetBool(name string, value *bool) bool {
	existingValue, existing := findInAllNames(name)
	if existing && existingValue.cType != "bool" {
		log.Fatal("can't set", name, "as it is already set with some other type")
	}

	if IsDrawing && existing {
		return *boolMap[name]
	}

	boolMap[name] = value
	names = append(names, namesStruct{name: name, cType: "bool"})
	informUI(name, *value, "bool")
	return *value
}

func SetBoolV(name string, value bool) bool {
	return SetBool(name, &value)
}

// only to to be used when we need to set value from UI to Go
func setBoolInternal(name string, value bool) {
	if boolMap[name] == nil {
		return
	}
	*(boolMap[name]) = value
}

func GetBool(name string) bool {
	return *boolMap[name]
}

func GetBoolPtr(name string) *bool {
	return boolMap[name]
}
