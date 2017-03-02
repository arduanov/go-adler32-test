package main

import (
	"fmt"
	"time"
	"hash/adler32"
	"crypto/sha1"
)

func main() {
	start := time.Now()

	for i := 0; i < 5000000; i++ {
		adler32.Checksum([]byte ("Какая-то не очень длинная строка, для которой надо посчитать контрольную сумму"))
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()

	for i := 0; i < 5000000; i++ {
		sha1.Sum([]byte ("Какая-то не очень длинная строка, для которой надо посчитать контрольную сумму"))
	}

	elapsed = time.Since(start)
	fmt.Println(elapsed)
}
