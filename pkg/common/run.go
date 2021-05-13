package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

//nolint:funlen // TODO
func run(cmd *exec.Cmd) {
	var wg sync.WaitGroup

	stdout, errOut := cmd.StdoutPipe()
	if errOut != nil {
		panic(errOut)
	}

	stderr, errErr := cmd.StderrPipe()
	if errErr != nil {
		panic(errErr)
	}

	errStart := cmd.Start()
	if errStart != nil {
		panic(errStart)
	}

	chPrint := make(chan string)
	scanOut := bufio.NewScanner(stdout)

	// This wait group covers the single goroutine started immediately below
	wg.Add(1)

	go func() {
		for scanOut.Scan() {
			chPrint <- scanOut.Text()
		}
		wg.Done()
	}()

	scanErr := bufio.NewScanner(stderr)

	// This wait group covers the single goroutine started immediately below
	wg.Add(1)

	go func() {
		for scanErr.Scan() {
			chPrint <- scanErr.Text()
		}
		wg.Done()
	}()

	var exitStatus int

	go func() {
		errWait := cmd.Wait()
		if errWait != nil {
			if exitErr, ok := errWait.(*exec.ExitError); ok {
				if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
					exitStatus = status.ExitStatus()
					log.Printf("Exit status: %d", exitStatus)
				}
			} else {
				panic(errWait)
			}
		}

		wg.Wait()
		close(chPrint)
	}()

	for line := range chPrint {
		fmt.Println(line)
	}

	os.Exit(exitStatus)
}
