package util

import (
	"fmt"
	"math/rand"
)

func RandomOwner() string {
	return fmt.Sprintf("owner_%d", rand.Intn(1000000))
}

func RandomMoney() int64 {
	return int64(rand.Intn(100000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "NGN"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
