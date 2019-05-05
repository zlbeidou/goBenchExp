# goBenchExp

linux:
```
goos: linux
goarch: amd64
pkg: git.code.oa.com/monitor_sys/benchExp
BenchmarkInlineFunc-4                   2000000000               0.41 ns/op            0 B/op          0 allocs/op
BenchmarkEmptyFuncVs0-4                 2000000000               0.41 ns/op            0 B/op          0 allocs/op
BenchmarkEmptyFunc-4                    1000000000               2.73 ns/op            0 B/op          0 allocs/op
BenchmarkSimpleDefer-4                  20000000                64.9 ns/op             0 B/op          0 allocs/op
BenchmarkGoEmptyFunc-4                   3000000               389 ns/op               0 B/op          0 allocs/op
BenchmarkStrToByteSlice-4               20000000               100 ns/op             112 B/op          1 allocs/op
BenchmarkByteSliceToStr-4               20000000                79.5 ns/op           112 B/op          1 allocs/op
BenchmarkByteSliceToStrDirect-4         2000000000               0.39 ns/op            0 B/op          0 allocs/op
BenchmarkTypeAssert-4                   1000000000               2.33 ns/op            0 B/op          0 allocs/op
BenchmarkSimpleReflectTypeOf-4          100000000               18.1 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleReflectValueOf-4         20000000                94.0 ns/op            16 B/op          1 allocs/op
BenchmarkGetFromCh-4                    10000000               239 ns/op               0 B/op          0 allocs/op
BenchmarkGetFromChGoroutine2-4          10000000               240 ns/op               0 B/op          0 allocs/op
BenchmarkGetFromBufferCh-4              20000000               119 ns/op               0 B/op          0 allocs/op
BenchmarkGetFromBufferChCap2-4          10000000               136 ns/op               0 B/op          0 allocs/op
BenchmarkJsonStd-4                        300000              5779 ns/op             544 B/op         18 allocs/op
BenchmarkJson3rd-4                       1000000              2019 ns/op             432 B/op         20 allocs/op
PASS
ok      git.code.oa.com/monitor_sys/benchExp    31.855s
```

windows:
```
goos: windows
goarch: amd64
pkg: github.com/zlbeidou/goBenchExp
BenchmarkInlineFunc-8                   2000000000               0.28 ns/op       0 B/op          0 allocs/op
BenchmarkEmptyFuncVs0-8                 2000000000               0.27 ns/op       0 B/op          0 allocs/op
BenchmarkEmptyFunc-8                    2000000000               1.86 ns/op       0 B/op          0 allocs/op
BenchmarkSimpleDefer-8                  30000000                46.8 ns/op        0 B/op          0 allocs/op
BenchmarkGoEmptyFunc-8                   5000000               279 ns/op          0 B/op          0 allocs/op
BenchmarkStrToByteSlice-8               30000000                47.0 ns/op      112 B/op          1 allocs/op
BenchmarkByteSliceToStr-8               30000000                40.4 ns/op      112 B/op          1 allocs/op
BenchmarkByteSliceToStrDirect-8         2000000000               0.28 ns/op       0 B/op          0 allocs/op
BenchmarkTypeAssert-8                   2000000000               1.64 ns/op       0 B/op          0 allocs/op
BenchmarkSimpleReflectTypeOf-8          100000000               13.0 ns/op        0 B/op          0 allocs/op
BenchmarkSimpleReflectValueOf-8         20000000                67.3 ns/op       16 B/op          1 allocs/op
BenchmarkGetFromCh-8                     5000000               247 ns/op          0 B/op          0 allocs/op
BenchmarkGetFromChGoroutine2-8           5000000               294 ns/op          0 B/op          0 allocs/op
BenchmarkGetFromBufferCh-8              20000000                88.4 ns/op        0 B/op          0 allocs/op
BenchmarkGetFromBufferChCap2-8          20000000                91.6 ns/op        0 B/op          0 allocs/op
BenchmarkJsonStd-8                        500000              3942 ns/op        544 B/op         18 allocs/op
BenchmarkJson3rd-8                       1000000              1424 ns/op        432 B/op         20 allocs/op
PASS
ok      github.com/zlbeidou/goBenchExp  28.952s
```
