Golang Cache Benchmark
======================

A benchmark of GoLang in-memory caching libraries.

*Disclaimer* Although this may improve over time, this benchmark was thrown together pretty quickly.  I've not made any
attempts to optimize the settings for any of the libraries tested, and I'm sure there are some gotchas that would bias
the results.  My goal was to understand _generally_ which library I should go with for my particular use case.  Reference
these benchmarks with caution.


Currently benching:
 - [FreeCache](https://github.com/coocood/freecache)
 - [MangoCache](https://github.com/goburrow/cache).  Note: these tests with the LoadingCache implementation were taking
   **way** too long, so I went with _only_ the standard cache approach (get and put, not auto-loading)
 
Latest Results:

#### 2017 MacBook Pro, 2.9 GHz Intel Core i7, 16 GB 2133 MHz RAM

(Loading Cache tests skipped as they were heavily increasing execution time)
```
goos: darwin
goarch: amd64
pkg: go_cache_bench
BenchmarkSingleKey_Get/FreeCache/keySize=100-8         	                 5000000	       269 ns/op
BenchmarkSingleKey_Get/MangoStandard/keySize=100-8     	                 2000000	       797 ns/op
BenchmarkSingleKey_Get/FreeCache/keySize=1000-8        	                30000000	        46.5 ns/op
BenchmarkSingleKey_Get/MangoStandard/keySize=1000-8    	                 2000000	       781 ns/op
BenchmarkSingleKey_Get/FreeCache/keySize=10000-8       	                30000000	        48.4 ns/op
BenchmarkSingleKey_Get/MangoStandard/keySize=10000-8   	                 2000000	       807 ns/op
BenchmarkSingleKey_SetAndGet/FreeCache/keySize=100-8   	                 2000000	       935 ns/op
BenchmarkSingleKey_SetAndGet/MangoStandard/keySize=100-8         	  500000	      2525 ns/op
BenchmarkSingleKey_SetAndGet/FreeCache/keySize=1000-8            	 5000000	       277 ns/op
BenchmarkSingleKey_SetAndGet/MangoStandard/keySize=1000-8        	 1000000	      1178 ns/op
BenchmarkSingleKey_SetAndGet/FreeCache/keySize=10000-8           	 5000000	       252 ns/op
BenchmarkSingleKey_SetAndGet/MangoStandard/keySize=10000-8       	 1000000	      1495 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=100/reads=2-8         	 1000000	      1323 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=100/reads=2-8     	  500000	      3953 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=100/reads=4-8         	 1000000	      2683 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=100/reads=4-8     	  200000	      8687 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=100/reads=8-8         	  300000	      4018 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=100/reads=8-8     	  200000	     13172 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=1000/reads=2-8        	 2000000	       520 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=1000/reads=2-8    	  200000	      6401 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=1000/reads=4-8        	 2000000	       743 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=1000/reads=4-8    	  200000	      9685 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=1000/reads=8-8        	 1000000	      1073 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=1000/reads=8-8    	  100000	     12741 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=10000/reads=2-8       	 2000000	       831 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=10000/reads=2-8   	  300000	      5004 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=10000/reads=4-8       	 3000000	       864 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=10000/reads=4-8   	  200000	     11589 ns/op
BenchmarkSingleKey_GetSetMultiGet/FreeCache/keySize=10000/reads=8-8       	 1000000	      1209 ns/op
BenchmarkSingleKey_GetSetMultiGet/MangoStandard/keySize=10000/reads=8-8   	  100000	     14076 ns/op
```
