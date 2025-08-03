package dogo_test

import (
	"bytes"
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

func ExampleDo_multiplevalues() {
	buf := bytes.NewBufferString("abc")
	v, err := dogo.Do(func() rune {
		r, s, err := buf.ReadRune()
		dogo.Check(err)
		fmt.Println(s)
		return r
	})
	fmt.Println(v, err)
	// Output:
	// 1
	// 97 <nil>
}
