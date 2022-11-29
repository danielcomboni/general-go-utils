/**
 * Author: Daniel Comboni
 */
 package general_goutils


import "testing"

type Test struct {
	Name string
	Age  int
}

func TestHasField(t *testing.T) {
	println(HasField[Test]("Name")) // prints true
	println(HasField[Test]("Non")) // prints false
}
