# PoolShark
Poolshark is a simple library that allows you to reduce the number of allocations you do for byte slices similar to sync.Pool.

![SHARK](poolshark.jpg?raw=true "Pool Shark")

# Why?
When profiling my applications I have found that sync.Pool frequently was causing lots of blocking. 
It was also still creating high volumes of garbage for the GC. Doing a simple change to use a [][]byte with a mutex immediately reduced my cpu usage by 1/2 and improved overall through put by 25%.


# How do I use it?

```go
import "github.com/vrecan/poolShark"

//init our pool shark. Allocate 5 []bytes that are all 30bytes in size.
pool := poolShark.NewBytePool(5, 30)
		bytes := pool.Get()
    //do some work
    //...
    pool.Put(bytes)
```

PoolShark is thread safe, frequently I will use it at a package level to control memory allocations across goroutines.
