/**
 * Author: Daniel Comboni
 */
package general_goutils

import (
	"fmt"
	"testing"
)

type Test struct {
	Name string
	Age  int
}

func TestHasField(t *testing.T) {
	println(HasField[Test]("Name")) // prints true
	println(HasField[Test]("name")) // prints false
}

func TestGetFieldValue(t *testing.T) {
	v := Test{
		Name: "Daniel",
		Age: 99,
	}

	u := GetFieldValue(&v,"Name")
	println(fmt.Sprintf("place: %v", u))

	println()

	println(HasField_CaseInsensitive[Test]("name"))

}
