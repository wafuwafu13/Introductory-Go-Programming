package main

import (
	"fmt"
	"errors"
	"os"
	"strings"
)

const rows, colmuns = 9, 9

type Grid [rows][colmuns]int8

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit = errors.New("invalid digit")
)

type SudokuError []error

func (g *Grid) Set(row, colmn int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, colmn) {
		errs = append(errs, ErrBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][colmn] = digit
	return nil
}

func inBounds(row, colmun int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if colmun < 0 || colmun >= colmuns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	if digit < 1 || digit >= 10 {
		return false
	}
	return true
}

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // エラー群を文字列に変換する
	}

	return strings.Join(s, ", ")
}

func main() {
	var g Grid
	err := g.Set(93, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok { // 型アサーションによりerrはSudokuErrorであるとアサートする
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e) // 2 error(s) occurred:
				                        // - out of bounds
				                        // - invalid digit
			}
		}
		os.Exit(1)
	}
}