package core

import (
	"reflect"
)

// Metadata contains info extracted from struct tags
type Metadata struct {
	Rules map[string]string
}

// ExtractMetadata reads 'brdd' tags from a struct
func ExtractMetadata(s interface{}) Metadata {
	meta := Metadata{
		Rules: make(map[string]string),
	}
	
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("brdd")
		if tag != "" {
			meta.Rules[field.Name] = tag
		}
	}
	
	return meta
}

// GetUseCaseID returns the use case ID from a class if it implements a specific method
func GetUseCaseID(u interface{}) string {
	if provider, ok := u.(interface{ UseCaseID() string }); ok {
		return provider.UseCaseID()
	}
	return "UNKNOWN"
}
