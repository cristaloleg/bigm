package bigm

import (
	"math/big"
)

var (
	bigZero   = big.NewInt(0)
	bigOne    = big.NewInt(1)
	Base      = []int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	BaseBig   = []*big.Int{}
	Reverse   = [][]int32{}
	BaseCount int
)

func init() {
	Init()
}

// Init ...
func Init() {
	BaseCount = len(Base)

	BaseBig = make([]*big.Int, BaseCount)
	for i := range BaseBig {
		BaseBig[i] = big.NewInt(int64(Base[i]))
	}

	Reverse = make([][]int32, BaseCount)
	for i := range Reverse {
		Reverse[i] = make([]int32, BaseCount)
	}

	for i := 0; i < BaseCount; i++ {
		for j := i + 1; j < BaseCount; j++ {
			Reverse[i][j] = modInverse(Base[i], Base[j])
		}
	}
	println("Init done.")
}

func modInverse(a, m int32) int32 {
	a = a % m
	for x := int64(1); x < int64(m); x++ {
		if (int64(a)*x)%int64(m) == 1 {
			return int32(x)
		}
	}
	return 0
}

// Int ...
type Int struct {
	nums []int32
}

// New ...
func New() *Int {
	ii := &Int{
		nums: make([]int32, BaseCount),
	}
	return ii
}

// FromString ...
func FromString(s string) *Int {
	i, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return nil
	}
	return FromBigInt(i)
}

// FromBigInt ...
func FromBigInt(x *big.Int) *Int {
	res := New()
	for i := range Base {
		tmp := big.NewInt(0).Rem(x, BaseBig[i])
		res.nums[i] = int32(tmp.Int64())
	}
	return res
}

// String ...
func (ii *Int) String() string {
	if ii == nil {
		return ""
	}
	return ii.BigInt().String()
}

// BigInt ...
func (ii *Int) BigInt() *big.Int {
	res := big.NewInt(0)
	mult := big.NewInt(1)

	x := make([]int32, BaseCount)
	for i := 0; i < BaseCount; i++ {
		x[i] = ii.nums[i]

		for j := 0; j < i; j++ {
			tmp := int64(x[i]-x[j]) * int64(Reverse[j][i])
			tmp2 := (tmp%int64(Base[i]) + int64(Base[i]))
			x[i] = int32(tmp2) % Base[i]
		}
		tmp := big.NewInt(0).Mul(mult, big.NewInt(int64(x[i])))
		res.Add(res, tmp)
		mult.Mul(mult, BaseBig[i])
	}
	return res
}

// Add ...
func Add(x, y *Int) *Int {
	res := New()
	for i := range res.nums {
		res.nums[i] = (x.nums[i] + y.nums[i]) % Base[i]
	}
	return res
}

// Sub ...
func Sub(x, y *Int) *Int {
	res := New()
	for i := range res.nums {
		res.nums[i] = (x.nums[i] - y.nums[i] + Base[i]) % Base[i]
	}
	return res
}

// Mul ...
func Mul(x, y *Int) *Int {
	res := New()
	for i := range res.nums {
		tmp := int64(x.nums[i] * y.nums[i])
		res.nums[i] = int32(tmp % int64(Base[i]))
	}
	return res
}

// ModInv ...
func ModInv(x, y *Int) *Int {
	xx := x.BigInt()
	yy := y.BigInt()
	tmp2 := big.NewInt(0).GCD(nil, nil, xx, yy)
	if tmp2.Cmp(bigOne) != 0 {
		return nil
	}
	tmp := big.NewInt(0).ModInverse(xx, yy)
	return FromBigInt(tmp)
}

// IsEqual ...
func IsEqual(x, y *Int) bool {
	for i := range x.nums {
		if x.nums[i] != y.nums[i] {
			return false
		}
	}
	return true
}

// IsCoprime ...
func IsCoprime(x, y *Int) bool {
	xx := x.BigInt()
	yy := y.BigInt()
	tmp2 := big.NewInt(0).GCD(nil, nil, xx, yy)
	return tmp2.Cmp(bigOne) == 0
}
