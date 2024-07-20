package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type Hi struct {
	Abol string
	Gogo int
}

type Ho struct {
	Abol string
	Gogo int
}

type He struct {
	Abol string
	Gogo int
	Arr  []int
}

func main() {
	g := He{Abol: "hello", Gogo: 3, Arr: []int{2, 3, 4}}
	v := reflect.ValueOf(g)
	length := v.NumField()
	baseURL := "http://example.com"
	resource := "/path"
	params := url.Values{}
	typeOfS := v.Type()
	for i := 0; i < length; i++ {
		name := typeOfS.Field(i).Name
		params.Add(name, FieldStrConverter(v.Field(i)))
	}

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	fmt.Printf("%v\n", u.String()) // "http://example.com/path?param1=value1&param2=value2"

}

func FieldStrConverter(field reflect.Value) string {

	var str string
	switch field.Kind() {
	case reflect.String:
		str = field.String()
	case reflect.Int:
		str = strconv.Itoa(int(field.Int()))
	case reflect.Slice:
		nums := field.Interface().([]int)
		for _, t := range nums {
			str += strconv.Itoa(t) + ","
		}
	default:
		{
		}
	}

	return str
}
