package executor

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func Execute(commandString string) {
	fmt.Fprintf(os.Stderr, "$ %s\n\n", commandString)
	execCmd := exec.Command("sh", "-c", commandString)
	stdout, err := execCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := execCmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	execCmd.Start()
	go scanAndPrint(stdout, os.Stdout)
	go scanAndPrint(stderr, os.Stderr)
	execCmd.Wait()
	os.Exit(execCmd.ProcessState.ExitCode())
}

func scanAndPrint(r io.ReadCloser, fp *os.File) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Fprintln(fp, m)
	}
}
