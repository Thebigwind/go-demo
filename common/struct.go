package common

import (
	"fmt"
	"reflect"
)

// GetByFieldName 根据字段名称获取值 效率低谨慎调用
func GetByFieldName(from interface{}, name string, def string) (string, error) {

	refT := reflect.TypeOf(from)
	refV := reflect.ValueOf(from)
	field, ok := refT.FieldByName(name)
	if !ok {
		return def, nil
	}
	if len(field.Index) <= 0 {
		return def, nil
	}

	valueI := refV.Field(field.Index[0])
	valueT, ok := valueI.Interface().(string)
	if ok {
		return valueT, nil
	}
	return def, fmt.Errorf("getByFieldName: not found %s", name)
}
