/**
 * Author: Daniel Comboni
 */

package general_goutils

import "reflect"

// HasField returns true if the struct [T] has the field specified by `fieldName` 
func HasField[T any](fieldName string) bool {
	metaValue := reflect.ValueOf(new(T)).Elem()
	field := metaValue.FieldByName(fieldName)
	return field != (reflect.Value{})
}
