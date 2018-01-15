package strtemplate

import (
	"reflect"
	"strconv"
)

func getStr(data interface{}) string {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Int:
		{
			return strconv.Itoa(data.(int))
		}
	case reflect.String:
		{
			return data.(string)
		}
	default:
		{
			return "???"
		}
	}
}
