/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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
