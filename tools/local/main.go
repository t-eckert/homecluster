package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	run("kind", "create", "cluster", "-n", "homecluster")
	run("kubectl", "apply", "-f", "./manifest")
}

func run(bin string, args ...string) {
	cmd := exec.Command(bin, args...)

	// Create a pipe for capturing command output
	reader, writer := io.Pipe()
	cmd.Stdout = writer

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting command:", err)
	}

	// Close the writer once command finishes writing to it
	defer writer.Close()

	// Read from the reader (pipe) in a separate goroutine and print the output
	go func() {
		defer reader.Close()
		// Buffer to store command output
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal("Error reading from pipe:", err)
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Fatal("Error waiting for command to finish:", err)
	}
}
