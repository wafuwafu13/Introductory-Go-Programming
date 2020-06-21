package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // 通信チャネルを作成

	for i := 0; i < 5; i++ {
		go sleepyGopher3(i, c)
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // 受信 time.Sleepを使わなくても待てるようになった
		fmt.Println("gopher", gopherID, "はスリープを終えました")
	}
}

func sleepyGopher3(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore...")
	c <- id // mainに送る
}

//  ... 4 snore...
//  gopher 4 はスリープを終えました
//  ... 3 snore...
//  gopher 3 はスリープを終えました
//  ... 2 snore...
//  gopher 2 はスリープを終えました
//  ... 1 snore...
//  ... 0 snore...
//  gopher 1 はスリープを終えました
//  gopher 0 はスリープを終えました
