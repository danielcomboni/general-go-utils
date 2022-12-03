/**
 * Author: Daniel Comboni
 */

package general_goutils

import (
	"reflect"
	"strings"
)

// HasField returns true if the struct [T] has the field specified by `fieldName`. This is a case sensitive test
func HasField[T any](fieldName string) bool {
	metaValue := reflect.ValueOf(new(T)).Elem()
	field := metaValue.FieldByName(fieldName)	
	return field != (reflect.Value{})
}

// GetFieldValue returns the value of the specified field/property of the struct t
func GetFieldValue[T any](t *T,field string) interface{} {
    r := reflect.ValueOf(t)
    f := reflect.Indirect(r).FieldByName(field)
    return f
}

// HasField_CaseInsensitive returns true if the struct [T] has the field specified by `fieldName`. This is a case insensitive test
func HasField_CaseInsensitive[T any](fieldName string) bool {
	var t T
	v := reflect.ValueOf(t)
	typeOfS := v.Type()

    for i := 0; i< v.NumField(); i++ {
		if strings.EqualFold(typeOfS.Field(i).Name,fieldName) {
			return true
		}
    }

	return false
}