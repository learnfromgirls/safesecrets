package example

import (
	"testing"

	"math/big"
	"fmt"
	"crypto/rand"
)

func TestExample(t *testing.T) {
	ExampleRun()
}


func TestLargePrime(t *testing.T) {

	for j := 0; j < 1; j++ {
		bts := 128 //in tests up to 12000 bits ran in reasonable time (hours)

		p, err := rand.Prime(rand.Reader, bts)
		var one big.Int
		one.SetUint64(1)

		if (err != nil) {
			fmt.Printf("random prime failed err=%v\n", err)
		} else {

			fmt.Printf("random number has bitlen=%v\n", p.BitLen())

			pp := p.ProbablyPrime(20)
			fmt.Printf("probable prime =%v\n", pp)
			fmt.Printf("prime =%v\n", p)

		}
	}
}
