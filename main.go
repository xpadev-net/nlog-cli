package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-shellwords"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
)

var config Config

func main() {
	config = getConfig()
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}
	length := len(os.Args)
	if length != 3 {
		println("invalid args")
		return
	}
	itemId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
	var command = os.Args[2]
	fmt.Printf("param -i : %s\n", itemId)
	fmt.Printf("param -i : %s\n", command)
	c, err := shellwords.Parse(command)
	if err != nil {
		return
	}
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
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
	taskId, err := createTask(itemId, currentUser.Username, workDir, command, pid)
	if err != nil {
		log.Fatalln(err)
	}
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
				_, err := appendLog(taskId, "out", stdoutData)
				if err != nil {
					log.Fatalln(err)
				}

			case stderrData, ok := <-stderrCh:
				if !ok {
					return
				}
				fmt.Printf("stderr: %s\n", stderrData)
				_, err := appendLog(taskId, "err", stderrData)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	exitCode := cmd.ProcessState.ExitCode()
	err = endTask(taskId, exitCode)
	if err != nil {
		log.Fatalln(err)
	}
}

func runCmdStr(c []string) *exec.Cmd {
	switch len(c) {
	case 1:
		return exec.Command(c[0])
	default:
		return exec.Command(c[0], c[1:]...)
	}
}
