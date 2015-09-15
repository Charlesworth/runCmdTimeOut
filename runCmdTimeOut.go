package runCmdTimeOut

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func test() {
	fmt.Println("passing cmd")
	out, timeOut, err := RunCmdTimeOut(*exec.Command("echo", "test!"), time.Second)
	fmt.Println("out: [", out, "] timeOut: [", timeOut, "] error: [", err, "]")

	fmt.Println("failing cmd")
	out, timeOut, err = RunCmdTimeOut(*exec.Command("shit", "test!"), time.Second)
	fmt.Println("out: [", out, "] timeOut: [", timeOut, "] error: [", err, "]")

	fmt.Println("timeout")
	out, timeOut, err = RunCmdTimeOut(*exec.Command("sleep", "5"), time.Second)
	fmt.Println("out: [", out, "] timeOut: [", timeOut, "] error: [", err, "]")
}

//RunCmdTimeOut takes a exec.Cmd argument and a timeOut, runs that command and returns stdOut, stdErr and a bool to indicate a time out
func RunCmdTimeOut(cmd exec.Cmd, timeOut time.Duration) (stdOut string, timedOut bool, stdErr error) {
	done := make(chan bool, 1)
	go func() {
		var cmdStdOut []byte
		cmdStdOut, stdErr = cmd.CombinedOutput()
		stdOut = string(cmdStdOut)
		done <- true
	}()

	select {
	//case timeout
	case <-time.After(timeOut):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatalln("Fatal Error: Cmd", cmd.Args, "timed out, but unable to kill process: Fatal Error")
		}
		<-done // allow goroutine to exit
		return "", true, errors.New(fmt.Sprint("Cmd ", cmd.Args, " timed out"))
	//case finished
	case <-done:
		return stdOut, false, stdErr
	}
}
