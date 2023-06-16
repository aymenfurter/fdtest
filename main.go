package main

import (
	"log"
	"os"
)

func main() {
	for i := 0; ; i++ {
		_, err := os.OpenFile("dummyFile.txt", os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
	}
}

