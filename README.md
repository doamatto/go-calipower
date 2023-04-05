| :exclamation:  [This repository is now maintained on Sourcehut here.]([https://git.sr.ht/~maatt/nobs-uuid](https://git.sr.ht/~maatt/go-calipower))   |
|------------------------------------------------------------------------------------------------------------------------------------------------------|

---

# go-calipower [![Go Reference](https://pkg.go.dev/badge/git.sr.ht/~maatt/go-calipower.svg)](https://pkg.go.dev/git.sr.ht/~maatt/go-calipower)
Get available capacity, estimates, and reserves from California ISO.

## Usage
```go
package main

import (
	"fmt"
	"log"
	"strconv"

	"git.sr.ht/~maatt/go-calipower"
)

func main() {
	data, err := calipower.GetPowerData()
	fmt.Printf(strconv.Itoa(data.CurrentReserve))
}
```
Better example can be found in `test/test.go`.

## Acknowledgements
Data is collected from FlexAlert.org, a website run by the California ISO. Data is meant to be used within the allowance of the California ISO.
