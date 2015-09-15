package runCmdTimeOut

import (
	"os/exec"
	"testing"
	"time"
)

func TestRunCmdTimeOut(t *testing.T) {
	cmdPass := *exec.Command("echo", "test")
	cmdFail := *exec.Command("failing", "cmd")
	cmdTimeOut := *exec.Command("sleep", "2")

	stdOut, timeOut, stdErr := RunCmdTimeOut(cmdPass, time.Second)
	if timeOut || stdErr != nil || stdOut == "" {
		t.Error("fail")
		//fmt.Println("out: [", out, "] timeOut: [", timeOut, "] error: [", err, "]")
	}

	stdOut, timeOut, stdErr = RunCmdTimeOut(cmdFail, time.Second)
	if timeOut || stdErr == nil {
		t.Error("fail")
		//fmt.Println("out: [", out, "] timeOut: [", timeOut, "] error: [", err, "]")
	}

	stdOut, timeOut, stdErr = RunCmdTimeOut(cmdTimeOut, time.Second)
	if !timeOut || stdErr == nil {
		t.Error("fail")
		//fmt.Println("out: [", out, "] timeOut: [", timeOut, "] error: [", err, "]")
	}
}
