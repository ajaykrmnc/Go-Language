package main

import (
	"os"
)

func main() {
	// panic("a problem")

	_, err := os.Create("go");
	if err != nil {
		panic(err);
	}
}