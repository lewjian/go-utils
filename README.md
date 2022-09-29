# utils
一个常用的go工具包，使用了go泛型，需要go1.18以上。（a package contains set of useful functions, such as set、bloom、goroutine group ...）

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
# 已实现功能清单
- array 常用的切片数组操作函数，如：InArray, Diff, Intersect等
- bloom 基于Redis的布隆过滤器
- codec aes、rsa等加密函数
- collection: 一些数据结构集合，如：set、queue、timing wheel等
- number: 数字操作相关，如四舍五入
- rescue: panic recover相关
- retry: 重试操作
- workergroup: 协程池


