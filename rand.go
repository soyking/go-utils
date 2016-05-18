package gu

import (
	"crypto/rand"
	"math/big"
)

// math.rand 返回随机数需要设置不同的 seed
func Intn(n int) int {
	max := big.NewInt(int64(n))
	i, _ := rand.Int(rand.Reader, max)
	return int(i.Int64())
}
