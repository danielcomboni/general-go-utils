/**
 * Author: Daniel Comboni
 */
 package general_goutils

 import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	println(ToCamelCase("the_email_of")) // result is: TheEmailOf
}
