**Numconvert** is a wrapper for internal number converters.

[![Build Status](https://travis-ci.org/rogozhka/numconvert.svg?branch=master)](https://travis-ci.org/rogozhka/numconvert)


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