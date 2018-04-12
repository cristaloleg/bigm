## bigm

Benchmark:
```
goos: darwin
goarch: amd64
pkg: github.com/cristaloleg/bigm
BenchmarkMul-4           1000000              1721 ns/op
BenchmarkMulBig-4       20000000               103 ns/op
BenchmarkAdd-4           2000000               818 ns/op
BenchmarkAddBig-4       20000000                75.1 ns/op
BenchmarkSub-4           2000000               788 ns/op
BenchmarkSubBig-4       20000000                78.8 ns/op
PASS
ok      github.com/cristaloleg/bigm     12.080s
```