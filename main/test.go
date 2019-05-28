package main

import (
	"fmt"
	"reflect"
)

type InterfaceFunc func() interface{}
type HandlerFunc func() (int, interface{})

type Endpoint struct {
	Entity  interface{}
	Handler HandlerFunc
	Params  interface{}
	URI     interface{}
}

type Params struct {
	ID int
}

type Entity struct {
	Name string
}

var endpoint = Endpoint{
	Entity: Entity{},
	Handler: func() (int, interface{}) {
		return 0, &struct{}{}
	},
	Params: Params{},
}

func main() {
	// fmt.Printf("%T - %+v\n", endpoint, endpoint)

	if endpoint.Handler == nil {
		panic("Handler missing")
	}

	var entity interface{}
	fmt.Println("Entity")

	if endpoint.Entity == nil {
		fmt.Println("- is nil")
	} else {
		entity = reflect.New(reflect.TypeOf(endpoint.Entity)).Interface()
		fmt.Printf("- %T: %+v\n", entity, entity)
	}

	var params interface{}
	fmt.Println("Params")

	if endpoint.Params == nil {
		fmt.Println("- is nil")
	} else {
		params = reflect.New(reflect.TypeOf(endpoint.Params)).Interface()
		fmt.Printf("- %T: %+v\n", params, params)
	}

	var uri interface{}
	fmt.Println("URI")

	if endpoint.URI == nil {
		fmt.Println("- is nil")
	} else {
		uri = reflect.New(reflect.TypeOf(endpoint.Params)).Interface()
		fmt.Printf("- %T: %+v\n", uri, uri)
	}
}

func structToMap(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("Interface is not a struct!")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)

		if tag := f.Tag.Get("key"); tag != "" {
			if f.Type.Kind() == reflect.Struct {
				out[tag] = structToMap(v.Field(i).Interface())
			} else {
				out[tag] = v.Field(i).Interface()
			}
		}
	}

	return out
}
