**Numconvert** is a wrapper of internal number converters.



## Install

```bash
go get -u github.com/rogozhka/numconvert
```



## Example

```go
package main

import (
	"github.com/rogozhka/numconvert"
)

func main() {

    str := "12345"
    log.Printf( numconvert.NumberInString(str).AsUint() )
}
```