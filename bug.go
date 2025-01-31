```go
package main

import "fmt"

func main() {
    var m map[string]int
    m["key"] = 10 // This will cause a panic if m is nil
    fmt.Println(m["key"])
}
```