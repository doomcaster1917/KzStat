package Tools

import (
	"fmt"
	"reflect"
	"strconv"
)

// Convers struct's fields to type string
func FieldsStrConverter(scheme interface{}) []string {
	v := reflect.ValueOf(scheme)
	length := v.NumField()
	arr := make([]string, 0, length)
	for i := 0; i < length; i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			arr = append(arr, v.Field(i).String())
		case reflect.Int:
			arr = append(arr, strconv.Itoa(int(v.Field(i).Int())))
		case reflect.Slice:
			str := ""
			z := v.Field(i).Interface().([]int)
			for _, t := range z {
				fmt.Printf("Полe массива %T", t)
				str += strconv.Itoa(t) + ","
			}
			arr = append(arr, str)
		default:
			{
			}
		}
	}
	return arr
}

// Convers field's value to type string
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
