// Q:  Write a function called killServer, which here, a pidFile and returns an error.
// It should read the process ID from the file, convert it to an integer, and instead of killing it, print just killing process ID.

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.New("can't open pid file (is server still running?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid file - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.New("bad PID")
	}

	fmt.Printf("Killing sever with PID - %d\n", pid)
	return nil
}

func main() {
	if err := killServer("sever.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
