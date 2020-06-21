package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64 // 先頭は必ず大文字
	}

	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	existOnError(err)

	fmt.Println(string(bytes)) // {"Lat":-4.5895,"Long":137.4417}
}

func existOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}