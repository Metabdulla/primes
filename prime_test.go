package primes

import (
	"fmt"
	"math/big"
	"testing"
	math "github.com/ethereum/go-ethereum/common/math"
	"time"
)

func TestPrimes(t *testing.T) {
	//201811161319
	//201811161331937
	//17908251027575790097
	primes := FindPrimes(10000)
	fmt.Println(len(primes), "len primes")
	start := time.Now()
	ok := IsPrimeFast(201811161331937, primes)
	fmt.Println(ok)
	fmt.Println(time.Now().Sub(start))
	start = time.Now()
	ok = IsPrime(201811161331937)
	fmt.Println(ok)
	fmt.Println(time.Now().Sub(start))
	return
	p := math.BigPow(2, 201811161319)
	y := big.NewInt(1)
	p.Sub(p, y)
	ok = SecPrime(p, 10000)
	fmt.Println(ok)
	return
	ns := []int{1, 3, 0, 20}
	for _, n := range ns {
		prime := p.ProbablyPrime(n)
		if prime == false {
			fmt.Println("not prime", prime)
			return
		}
	}
	fmt.Println("prime")
}


