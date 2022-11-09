package mapx

import (
	"reflect"
	"time"
)

// Structs2Maps structs to maps
func Structs2Maps(data ...any) []map[string]interface{} {
	var maps []map[string]interface{}
	for _, v := range data {
		maps = append(maps, Struct2Map(v))
	}

	return maps
}

// Struct2Map struct to map
func Struct2Map(data any) map[string]interface{} {

	m := make(map[string]interface{})

	switch data.(type) {
	case map[string]interface{}:
		return data.(map[string]interface{})
	}

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := v.Type().Field(i).Tag.Get("json")

			if fieldName == "" {
				continue
			}

			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			value, isTime := field.Interface().(time.Time)
			if isTime {
				m[fieldName] = value.Format("2006-01-02 15:04:05")

				// 增加一个时间戳字段，方便其它地方用来排序
				switch fieldName {
				case "createdAt":
					m["createdTime"] = value.Unix()
				case "updatedAt":
					m["updatedTime"] = value.Unix()
				}
			} else {
				m[fieldName] = field.Interface()
			}
		}
	}

	return m
}
