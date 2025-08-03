package dogo_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/seiyab/dogo"
)

func TestDo_Error(t *testing.T) {
	_, err := dogo.Do(func() int {
		a := dogo.Must(strconv.Atoi("2"))
		b := dogo.Must(strconv.Atoi("abc"))
		return a + b
	})
	if err == nil {
		t.Error("expected error")
	}
}

func ExampleDo() {
	v, err := dogo.Do(func() int {
		a := dogo.Must(strconv.Atoi("1"))
		b := dogo.Must(strconv.Atoi("2"))
		return a + b
	})

	fmt.Println(v, err)
	// Output: 3 <nil>
}
