## bigm

Benchmark:
```
goos: darwin
goarch: amd64
pkg: github.com/cristaloleg/bigm
BenchmarkMul-4          10000000               154 ns/op
BenchmarkMulBig-4        3000000               563 ns/op
BenchmarkAdd-4          20000000                97.0 ns/op
BenchmarkAddBig-4       20000000               101 ns/op
BenchmarkSub-4          20000000                93.2 ns/op
BenchmarkSubBig-4       20000000               104 ns/op
PASS
ok      github.com/cristaloleg/bigm     12.317s
```