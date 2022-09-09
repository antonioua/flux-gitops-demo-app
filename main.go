package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s := "Hello world"
	fmt.Println(s)
	healthFile, err := os.Create("/healthcheck")
	if err != nil {
		panic(err)
	}

	if err = healthFile.Close(); err != nil {
		panic(err)
	}

	time.Sleep(time.Minute * 5)
}
