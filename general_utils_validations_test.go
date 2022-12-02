package general_goutils

import "testing"


func TestIsGreaterThan(t *testing.T) {
	result := IsGreaterThan(1,0)
	println(result)
}


func TestIsGreaterThanOrEqualTo(t *testing.T) {
	result := IsGreaterThanOrEqualTo(0,0)
	println(result)
}

func TestIsLessThan(t *testing.T) {
	result := IsLessThan(-1,0)
	println(result)
}

func TestIsLessThanOrEqualTo(t *testing.T) {
	result := IsLessThanOrEqualTo(0,0)
	println(result)
}


func TestStringContains(t *testing.T) {
	result := StringContains("Daniel","iel")
	println(result)
}
