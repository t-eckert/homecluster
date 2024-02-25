package run

import (
	"fmt"
	"io"
	"os/exec"
)

// Cmd runs a command and streams its output to the terminal.
func Cmd(bin string, args ...string) error {
	cmd := exec.Command(bin, args...)

	reader, writer := io.Pipe()
	cmd.Stdout = writer

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting command: %v", err)
	}
	defer writer.Close()

	// Read from the reader (pipe) in a separate goroutine and print the output
	go func() {
		defer reader.Close()
		buf := make([]byte, 1024*8) // 8 KB buffer
		for {
			n, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for command to finish: %v", err)
	}

	return nil
}
