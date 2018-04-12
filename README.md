## bigm

Benchmark:
```
goos: darwin
goarch: amd64
pkg: github.com/cristaloleg/bigm
BenchmarkMul-4           1000000              1714 ns/op
BenchmarkMulBig-4       20000000               106 ns/op
BenchmarkAdd-4           1000000              1098 ns/op
BenchmarkAddBig-4       20000000                79.4 ns/op
BenchmarkSub-4           2000000               825 ns/op
BenchmarkSubBig-4       20000000                80.8 ns/op
PASS
ok      github.com/cristaloleg/bigm     11.065s
```