package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := make(chan int) // 通信チャネルを作成

	for i := 0; i < 5; i++ {
		go sleepyGopher3(i, c)
	}

	timeout := time.After(2 * time.Second) // 1個のチャネルを返す 8秒とかだったら受信される前にプログラムが終了する

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "はスリープを終えました")
		
		case <-timeout:
			fmt.Println("忍耐の限度に達した！")
			return
		}
	}
}

func sleepyGopher3(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...", id, "snore...")
	c <- id // mainに送る
}


// ... 4 snore...
// gopher 4 はスリープを終えました
// 忍耐の限度に達した！
