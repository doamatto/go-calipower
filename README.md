# go-calipower
Get available capacity, estimates, and reserves from California ISO.

## Usage
```go
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/doamatto/go-calipower"
)

func main() {
	data, err := calipower.GetPowerData()
	fmt.Printf(strconv.Itoa(data.CurrentReserve))
}
```
Better example can be found in `test/test.go`.

## Acknowledgements
Data is collected from FlexAlert.org, a website run by the California ISO. Data is meant to be used within the allowance of the California ISO.
