package main

import (
	"log"

	"github.com/t-eckert/homecluster/internal/run"
)

func main() {
	if err := run.Cmd("kind", "create", "cluster", "-n", "homecluster"); err != nil {
		log.Fatal(err)
	}
	if err := run.Cmd("kubectl", "apply", "-f", "./manifest"); err != nil {
		log.Fatal(err)
	}
}
