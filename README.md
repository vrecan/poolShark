# PoolShark [![Build Status](https://travis-ci.org/vrecan/poolShark.svg?branch=master)](https://travis-ci.org/vrecan/poolShark)
Poolshark is a simple library that allows you to reduce the number of allocations you do for byte slices similar to sync.Pool.

![SHARK](poolshark.jpg?raw=true "Pool Shark")

# Why?
When profiling my applications I have found that sync.Pool frequently was causing lots of blocking. 
It was also still creating high volumes of garbage for the GC. Doing a simple change to use a [][]byte with a mutex immediately reduced my cpu usage by 1/2 and improved overall throughput by 25% but in micro benchmarks was slower then sync.pool when used from multiple threads. Finally I implemented a version just using a channel and found that it was faster then both implementations in all cases.

# How do I use it?

```go
import "github.com/vrecan/poolShark"

//init our pool shark. Allocate 5 []bytes that are all 30bytes in size.
pool := poolShark.NewByteChanPool(5, 30)
		bytes := pool.Get()
    //do some work
    //...
    pool.Put(bytes)
```

#Benchmarks
Interesting note is that the channel based version is faster in all instances.
```
#Mutex around [][]byte
BenchmarkBytePoolSizehundredKBInflight100-16        	  200000	      7069 ns/op	    2560 B/op	       1 allocs/op
BenchmarkBytePoolSizeoneMBInflight100-16            	  200000	      7386 ns/op	    2560 B/op	       1 allocs/op
BenchmarkBytePoolSizeoneMBInflight10000-16          	    2000	    597217 ns/op	  245837 B/op	       1 allocs/op
BenchmarkBytePool2Concurrent-16                     	    1000	   1120469 ns/op	  491765 B/op	       4 allocs/op
BenchmarkBytePool5Concurrent-16                     	     200	   7712267 ns/op	 1229824 B/op	       7 allocs/op
BenchmarkBytePool10Concurrent-16                    	     100	  19092812 ns/op	 2459643 B/op	      14 allocs/op
#sync.Pool for []byte
BenchmarkByteSyncPoolSizehundredKBInflight100-16    	   50000	     30040 ns/op	   14559 B/op	     101 allocs/op
BenchmarkByteSyncPoolSizeoneMBInflight100-16        	   50000	     33558 ns/op	   14572 B/op	     101 allocs/op
BenchmarkByteSyncPoolSizeoneMBInflight10000-16      	     500	   3243073 ns/op	 1456106 B/op	   10017 allocs/op
BenchmarkByteSyncPool2Concurrent-16                 	     500	   3039794 ns/op	 2683047 B/op	   20059 allocs/op
BenchmarkByteSyncPool5Concurrent-16                 	     300	   4785348 ns/op	 5913857 B/op	   50344 allocs/op
BenchmarkByteSyncPool10Concurrent-16                	     200	   7661518 ns/op	11468424 B/op	  100802 allocs/op
#fixed sized channel for []byte
BenchmarkByteChanPoolSizehundredKBInflight100-16    	  300000	      3795 ns/op	    2560 B/op	       1 allocs/op
BenchmarkByteChanPoolSizeoneMBInflight100-16        	  500000	      3184 ns/op	    2560 B/op	       1 allocs/op
BenchmarkByteChanPoolSizeoneMBInflight10000-16      	    5000	    294494 ns/op	  245786 B/op	       1 allocs/op
BenchmarkByteChanPool2Concurrent-16                 	    5000	    312461 ns/op	  491613 B/op	       4 allocs/op
BenchmarkByteChanPool5Concurrent-16                 	    2000	    561528 ns/op	 1228946 B/op	       7 allocs/op
BenchmarkByteChanPool10Concurrent-16                	    2000	    789672 ns/op	 2457745 B/op	      12 allocs/op
```

PoolShark is thread safe, frequently I will use it at a package level to control memory allocations across goroutines.
