package main

import (
	"fmt"
	"os"
	"io"
)

type safeWriter struct {
	w io.Writer
	err error // ここに最初のエラーをストアする
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully")

	return err
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // すでにエラーが発生していたら書き込みをスキップ
	}

	_, sw.err = fmt.Fprintln(sw.w, s) // 1行を書き、もしエラーばあればストアする
}

func proverbs2(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully")

	return sw.err // エラーが発生したら、そのエラーを返す
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	proverbs2("proverbs2.txt")
}
