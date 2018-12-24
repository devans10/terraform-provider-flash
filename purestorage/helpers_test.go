package purestorage

import (
	"testing"
)

func Test_difference(t *testing.T) {
	slice1 := []string{"One", "Two", "Three"}
	slice2 := []string{"One", "Two", "Three", "foo"}

	diff := difference(slice2, slice1)

	if diff[0] != "foo" {
		t.Fatalf("Wrong value returned: %s", diff)
	}
}

func Test_sameStringSlice_true(t *testing.T) {
	slice1 := []string{"One", "Two", "Three"}
	slice2 := []string{"One", "Two", "Three"}

	if !sameStringSlice(slice1, slice2) {
		t.Fatal("Returned false")
	}
}

func Test_sameStringSlice_false(t *testing.T) {
	slice1 := []string{"One", "Two", "Three"}
	slice2 := []string{"One", "Two", "Three", "foo"}

	if sameStringSlice(slice1, slice2) {
		t.Fatal("Returned true")
	}
}

func Test_stringInSlice(t *testing.T) {
	slice1 := []string{"One", "Two", "Three", "foo"}
	if !stringInSlice("foo", slice1) {
		t.Fatal("Returned false")
	}
}
