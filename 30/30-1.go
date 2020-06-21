package main

import (
	"fmt"
	"time"
)

func main() {
	// ゴルーチンじゃなかったら、3秒経過 => 出力 => 4秒経過 => 終了
	go sleepyGopher()

	time.Sleep(4 * time.Second) // 4秒待たなかったら、出力される前にプログラムが終了する

} // ここに到達すると、すべてのゴルーチンが停止する

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}