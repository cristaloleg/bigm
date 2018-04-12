package bigm

import (
	"math/big"
	"testing"
)

var sink interface{}

const (
	xs = "321321321321321321321321321321321321321321321321321321321"
	ys = "123123123123123123123123123123123123123123123123123123123"
)

func BenchmarkMul(b *testing.B) {
	x := FromString(xs)
	y := FromString(ys)
	for i := 0; i < b.N; i++ {
		sink = Mul(x, y)
	}
}

func BenchmarkMulBig(b *testing.B) {
	x, _ := big.NewInt(0).SetString(xs, 10)
	y, _ := big.NewInt(0).SetString(ys, 10)
	for i := 0; i < b.N; i++ {
		sink = big.NewInt(0).Mul(x, y)
	}
}

func BenchmarkAdd(b *testing.B) {
	x := FromString(xs)
	y := FromString(ys)
	for i := 0; i < b.N; i++ {
		sink = Add(x, y)
	}
}

func BenchmarkAddBig(b *testing.B) {
	x, _ := big.NewInt(0).SetString(xs, 10)
	y, _ := big.NewInt(0).SetString(ys, 10)
	for i := 0; i < b.N; i++ {
		sink = big.NewInt(0).Add(x, y)
	}
}

func BenchmarkSub(b *testing.B) {
	x := FromString(xs)
	y := FromString(ys)
	for i := 0; i < b.N; i++ {
		sink = Sub(x, y)
	}
}

func BenchmarkSubBig(b *testing.B) {
	x, _ := big.NewInt(0).SetString(xs, 10)
	y, _ := big.NewInt(0).SetString(ys, 10)
	for i := 0; i < b.N; i++ {
		sink = big.NewInt(0).Sub(x, y)
	}
}
