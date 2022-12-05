/**
 * Author: Daniel Comboni
 */

 package general_goutils
 
 import "github.com/iancoleman/strcase"

 func ToCamelCaseLower(s string) string {
	return strcase.ToLowerCamel(s)
}

func ToSnakeCase(s string) string {
	return strcase.ToSnake(s)
}

func ToCamelCase(s string) string {
	return strcase.ToCamel(s)
}