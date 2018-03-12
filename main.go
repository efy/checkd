package main

import (
	"fmt"
	"os"
)

func main() {
	statsdUrl := os.Getenv("STATSD_URL")

	if statsdUrl == "" {
		fmt.Println("STATSD_URL must be provided")
		os.Exit(1)
	}
}
