# Spinner

## Overview
Spinner is a simple and efficient Go package that provides a rotating/spinning data structure. It is implemented using Go's new generics feature, allowing it to hold elements of any type. Under the hood, it uses Go's `container/ring` package to maintain a circular list of elements.

## Usage

```bash
go get github.com/twiny/poxa
```

```go
package main

import (
	"fmt"
	"github.com/twiny/poxa"
)

func main() {
	sp := poxa.NewSpinner[int](1, 2, 3)

	for i := 0; i < 5; i++ {
		fmt.Println(sp.Next())
	}
}
```