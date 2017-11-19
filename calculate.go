package main

import (
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := []byte("t :$=&]0~881:T<*:[G1(*6>^|686@-G")

	for cost := bcrypt.MinCost; cost <= bcrypt.MaxCost; cost++ {

		start := time.Now()

		bcrypt.GenerateFromPassword(password, cost)

		elapsed := time.Since(start)
		fmt.Printf("Cost %d took %s\n", cost, elapsed)

	}

}
