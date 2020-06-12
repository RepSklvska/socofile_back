package swp

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"math/rand"
)

func random1(a ...string) int {
	s := ""
	for _, v := range a {
		s += v
	}
	
	HexString := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	seedBigInt, ok := new(big.Int).SetString(HexString, 16)
	if !ok {
		return 0
	}
	seedInt := seedBigInt.Int64()
	// seedBigInt mod 2^31
	// md5's sum may needs a int128 type to contain...
	
	rand.Seed(seedInt)
	return rand.Int()
}

func random2(a ...string) int {
	seedBigInt := new(big.Int)
	
	for _, v := range a {
		HexString := fmt.Sprintf("%x", md5.Sum([]byte(v)))
		BigInt, ok := new(big.Int).SetString(HexString, 16)
		if !ok {
			return 0
		}
		seedBigInt.Add(seedBigInt, BigInt)
		// seedBigInt += BigInt
	}
	seedInt := seedBigInt.Int64()
	// seedBigInt1 mod 2^31
	
	rand.Seed(seedInt)
	return rand.Int()
}

func random3(a string, b int) int {
	HexString := fmt.Sprintf("%x", md5.Sum([]byte(a)))
	seedBigInt, ok := new(big.Int).SetString(HexString, 16)
	if !ok {
		return 0
	}
	seedBigInt.Add(seedBigInt, big.NewInt(int64(b)))
	seedInt := seedBigInt.Int64()
	
	rand.Seed(seedInt)
	return rand.Int()
}
