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

	_, timeOut, stdErr := RunCmdTimeOut(cmdPass, time.Second)
	if timeOut {
		t.Error("cmdPass has timed out")
	}
	if stdErr != nil {
		t.Error("cmdPass returned with a stdErr")
	}

	_, timeOut, stdErr = RunCmdTimeOut(cmdFail, time.Second)
	if timeOut {
		t.Error("cmdFail timed out")
	}
	if stdErr == nil {
		t.Error("cmdFail did not return a stdErr")
	}

	_, timeOut, stdErr = RunCmdTimeOut(cmdTimeOut, time.Millisecond)
	if !timeOut {
		t.Error("cmdTimeOut did not return with timeOut == true")
	}
}
