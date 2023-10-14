package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-shellwords"
	"os"
	"os/exec"
)

func main() {
	getConfig()
	length := len(os.Args)
	if length != 3 {
		println("invalid args")
		return
	}
	var i = os.Args[1]
	var j = os.Args[2]
	fmt.Printf("param -i : %s\n", i)
	fmt.Printf("param -i : %s\n", j)
	c, err := shellwords.Parse(j)
	if err != nil {
		return
	}
	cmd := runCmdStr(c)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	stdoutCh := make(chan string)
	stderrCh := make(chan string)

	err = cmd.Start()
	if err != nil {
		fmt.Printf("コマンドの実行エラー: %s\n", err)
		return
	}

	pid := cmd.Process.Pid
	fmt.Printf("param -i : %s\n", pid)
	go func() {
		defer close(stdoutCh)
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			stdoutCh <- scanner.Text()
		}
	}()

	go func() {
		defer close(stderrCh)
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			stderrCh <- scanner.Text()
		}
	}()

	go func() {
		for {
			select {
			case stdoutData, ok := <-stdoutCh:
				if !ok {
					return
				}
				fmt.Printf("stdout: %s\n", stdoutData)

			case stderrData, ok := <-stderrCh:
				if !ok {
					return
				}
				fmt.Printf("stderr: %s\n", stderrData)
			}
		}
	}()
	err = cmd.Wait()
	if err != nil {
		return
	}
	fmt.Printf("param -i : %s\n", err)
	println("Hello world!")
}

func runCmdStr(c []string) *exec.Cmd {
	switch len(c) {
	case 1:
		return exec.Command(c[0])
	default:
		return exec.Command(c[0], c[1:]...)
	}
}
