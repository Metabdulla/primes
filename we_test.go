package primes

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestExample0(t *testing.T) {
	var product = 894797020974016837
	N := 1000000
	start := time.Now()
	primes := FindPrimes(N)
	fmt.Println(len(primes), "len primes", time.Now())
	var first, second int
	for _, prime := range primes {
		if product%prime == 0 {
			second = prime
			first = product / prime
			//if !IsPrimeFast(first,primes) {
			//	t.Fatal(first,second ,"not prime")
			//}
			fmt.Println(first, second)
			fmt.Println(time.Now().Sub(start), time.Now())
			return
		}
	}
	for i := N + 1; float64(i) < math.Sqrt(float64(product)); i = i + 2 {
		ok := IsPrimeFast(i, primes)
		if ok {
			if product%i == 0 {
				second = i
				first = product / i
				//if !IsPrimeFast(first,primes) {
				//	t.Fatal(first,second ,"not prime")
				//}
				fmt.Println(first, " ", second)
				fmt.Println(time.Now().Sub(start))
				return
			}
		}
	}
	t.Fatal("problem error")
}

func TestExample1(t *testing.T) {

	var product = 894797020974016837
	var first, second int
	start := time.Now()
	for i := 3; float64(i) < math.Sqrt(float64(product)); i = i + 2 {
		if IsPrime(i) {
			if product%i == 0 {
				second = i
				first = product / i
				//if !IsPrime(first) {
				//	t.Fatal(first, second, "not prime")
				//}
				fmt.Println(first, second)
				fmt.Println(time.Now().Sub(start))
				return
			}
		}
	}
	t.Fatal("failed")
}

func TestExample2(t *testing.T) {

	var product = 894797020974016837
	//707829217
	var first, second int
	start := time.Now()
	for i := 3; float64(i) < math.Sqrt(float64(product)); i = i + 2 {
		if IsPrimeFast(i, nil) {
			if product%i == 0 {
				second = i
				first = product / i
				//if !IsPrimeFast(first,nil) {
				//	t.Fatal(first, second, "not prime")
				//}
				fmt.Println(first, second)
				fmt.Println(time.Now().Sub(start))
				return
			}
		}
	}
	t.Fatal("failed")
}

func TestSecPrime(t *testing.T) {
	var first, second int
	beginN := 2563897585
	endNum := 348998695
	primes := FindPrimes(10000)
	var i int
	for i = beginN; i < beginN*2; i = i + 2*6+34  {
		if IsPrimeFast(i, primes) {
				first = i
				fmt.Println(first)
				break

		}
	}
	fmt.Println(i)
	if first == 0 {
		t.Fatal("fail")
	}
	for i = endNum; i < endNum*2; i = i + 2*5 + 14 {
		if IsPrimeFast(i, primes) {
			second = i
			fmt.Println(first*second, first, second)
			return
		}

	}
	//894797020974016837 2563897723 348998719
	t.Fatal("fail")
}
