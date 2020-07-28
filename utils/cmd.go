package utils

import (
	"bufio"
	"io"
	"log"
    "os/exec"
    "bytes"
)

func ExecCommandAndReturn(commandName string, args ...string) string {
    cmd := exec.Command(commandName, args...)
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        return err.Error()
    }

    return out.String()
}

func ExecCommand(commandName string, args ...string) {
	cmd := exec.Command(commandName, args...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
	}
	
    stderr, err := cmd.StderrPipe()
    if err != nil {
		log.Fatal(err)
	}
	
    err = cmd.Start()
    if err != nil {
        log.Fatal(err)
    }

    go printOut(stdout)
	go printOut(stderr)
	
    cmd.Wait()
}

func printOut(r io.Reader) {
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        println(scanner.Text())
    }
}