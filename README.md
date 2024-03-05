# uuidv7-go

This package implements a UUIDv7 generator as specified by the [New UUID Formats](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7).
The first 48 bits contain a big-endian epoch timestamp in milliseconds.
This variant provides 74 bytes of entropy.

## Examples

### Generate UUIDv7 as string

```go
package main

import (
    "fmt"
    "github.com/moroz/uuidv7-go"
)

func main() {
    uuid := uuidv7.Generate()
    fmt.Println(uuid)
}
```

```shell
$ go run .
018e0cd7-986e-7f77-a8b4-8d55de59c992
```

## License

BSD 3-Clause "New" or "Revised" License
