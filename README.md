# utils
一个常用的go工具包，使用了go泛型，需要go1.18以上

# 使用方法

```go
package main

import (
    "fmt"

    "github.com/lewjian/utils/collection"
)

func main() {
    s := collection.NewSet[int]()
    s.Add(1, 2, 3, 1, 2, 3)
    fmt.Printf("%v", s.Values())
}
```

