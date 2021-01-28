package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	var prime, ok = new(big.Int).SetString("8683317618811886495518194401279999999", 0)
	if !ok {
		log.Fatal("Fuck")
	}

	i, _ := new(big.Int).SetString("2", 0)
	z := new(big.Int)
	m := new(big.Int)

	zero, _ := new(big.Int).SetString("0", 0)

	milestone, _ := new(big.Int).SetString("100000000", 0)
	milestoneI, _ := new(big.Int).SetString("0", 0)
	milestoneOne, _ := new(big.Int).SetString("1", 0)

	for {
		_, b := z.DivMod(prime, i, m)
		if b.Cmp(zero) == 0 {
			fmt.Println(i)
			break
		}
		milestoneI.Add(milestoneI, milestoneOne)

		if milestoneI.Cmp(milestone) == 0 {
			fmt.Println(milestoneI)
		}
	}
}
