# dogo

[![Go Report Card](https://goreportcard.com/badge/github.com/seiyab/dogo)](https://goreportcard.com/report/github.com/seiyab/dogo)
[![Go Reference](https://pkg.go.dev/badge/github.com/seiyab/dogo.svg)](https://pkg.go.dev/github.com/seiyab/dogo)

`dogo` provides `try-catch`-like error handling for Go to reduce boilerplate `if err != nil` checks.

## Installation

```sh
go get github.com/seiyab/dogo
```

## Usage

The `dogo.Do` function executes a block of code and recovers from any panics triggered by `dogo.Check` or `dogo.Must`, returning them as a standard error.

### Basic Example

This example shows a simple case where no error occurs.

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/seiyab/dogo"
)

func main() {
	v, err := dogo.Do(func() int {
		a := dogo.Must(strconv.Atoi("1"))
		b := dogo.Must(strconv.Atoi("2"))
		return a + b
	})

	fmt.Println(v, err)
	// Output: 3 <nil>
}
```

### Error Handling

When an error is encountered inside the `Do` block, it's returned as the second value.

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/seiyab/dogo"
)

func main() {
	v, err := dogo.Do(func() int {
		a := dogo.Must(strconv.Atoi("2"))
		b := dogo.Must(strconv.Atoi("abc")) // This will fail
		return a + b
	})

	fmt.Println(v, err)
	// Output: 0 strconv.Atoi: parsing "abc": invalid syntax
}
```

### Using `Check`

You can use `dogo.Check` for functions that only return an error.

```go
package main

import (
	"fmt"

	"github.com/seiyab/dogo"
)

func main() {
	_, err := dogo.Do(func() any {
		dogo.Check(fmt.Errorf("something went wrong"))
		return nil
	})

	fmt.Println(err)
	// Output: something went wrong
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
