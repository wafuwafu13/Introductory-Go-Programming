package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher2(i) // ... 1 snore ...
		                    // ... 2 snore ...
		                    // ... 4 snore ...
		                    // ... 0 snore ...
		                    // ... 3 snore ...
	}
	time.Sleep(4 * time.Second) // 実際に待つ必要があるのは3秒ちょっとなのに、4秒も待ってしまっている
}

func sleepyGopher2(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id,  "snore ...")
}