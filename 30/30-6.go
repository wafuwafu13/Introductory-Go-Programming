package main

import (
	"fmt"
	"strings"
)

// ラインの始点
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	// downstream <- "" // おしまいを知らせる
	close(downstream) // チャネルに値を書き込めなくなる
}

// 上流から品の値を読み、badという文字列を含んでいないときに限り、それを下流のチャネルに送る
func filterGopher(upstream, downstream chan string) {
	//for {
		//item, ok := <- upstream

		// if item == "" {
		// 	downstream <- ""
		// 	return
		// }

		// if !ok {
		// 	close(downstream)
		// 	return
		// }
		// if !strings.Contains(item, "bad") {
		// 	downstream <- item
		// }

	// クローズされるまでチャネルから値を読み出す
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
	    }
	}
	close(downstream)
	//}
}

// リターンすると、他の2つのゴルーチンが仕事を終えたことがわかる
func printGopher(upstream chan string) {
	// for {
	// 	v := <-upstream
	// 	if v == "" {
	// 		return
	// 	}
	// 	fmt.Println(v)
	// }
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}