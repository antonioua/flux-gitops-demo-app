package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s := "Hello world2"
	healthFile, err := os.Create("/tmp/healthy")
	if err != nil {
		panic(err)
	}

	if err = healthFile.Close(); err != nil {
		panic(err)
	}

	fmt.Println(s)

	time.Sleep(time.Minute * 5)
}
