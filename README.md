## bigm

Benchmark:
```
goos: darwin
goarch: amd64
pkg: github.com/cristaloleg/bigm
BenchmarkMul-4           1000000              1897 ns/op
BenchmarkMulBig-4         500000              2330 ns/op
BenchmarkAdd-4           2000000               809 ns/op
BenchmarkAddBig-4       10000000               141 ns/op
BenchmarkSub-4           2000000               899 ns/op
BenchmarkSubBig-4       10000000               141 ns/op
PASS
ok      github.com/cristaloleg/bigm     11.436s
```


// Cmp(y *Int) (r int)
// Div(x, y *Int) *Int
// DivMod(x, y, m *Int) (*Int, *Int)
// Exp(x, y, m *Int) *Int
// Format(s fmt.State, ch rune)
// GCD(x, y, a, b *Int) *Int
// GobDecode(buf []byte) error
// GobEncode() ([]byte, error)
// MarshalJSON() ([]byte, error)
// MarshalText() (text []byte, err error)
// Mod(x, y *Int) *Int
// ModSqrt(x, p *Int) *Int
// ProbablyPrime(n int) bool
// Quo(x, y *Int) *Int
// QuoRem(x, y, r *Int) (*Int, *Int)
// Rand(rnd *rand.Rand, n *Int) *Int
// Rem(x, y *Int) *Int
// Sqrt(x *Int) *Int
// UnmarshalJSON(text []byte) error
// UnmarshalText(text []byte) error
