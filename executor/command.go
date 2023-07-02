package executor

import (
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)

type CommandOutputWriter struct {
	TimeoutSecond int
}

type CmdExecutionResult struct {
  Stdout []byte
  Stderr []byte
}


func (c *CommandOutputWriter) ExecFullOutput(command string) (outOk []byte, outErr []byte, err error) {

  chResponse := make(chan CmdExecutionResult)

  executor := viper.GetString("shell_executor")
	if executor == "" {
		executor = "/bin/bash"
	}

	cmd := exec.Command(executor, "-c", command)

	go func(cnl chan CmdExecutionResult, cmd *exec.Cmd) {

    out, err := cmd.CombinedOutput()
    errOut := []byte("")
    if err != nil {
      errOut = []byte(err.Error())
    }
    cnl <- CmdExecutionResult{Stdout: out, Stderr: errOut}

	}(chResponse, cmd)

  if c.TimeoutSecond == 0 {
    result := <-chResponse
    outOk = result.Stdout
    outErr = result.Stderr
    return
  }

  select {
    case result := <-chResponse:
      outOk = result.Stdout
      outErr = result.Stderr
      cmd.Process.Kill()
	  case <-time.After(time.Second * time.Duration(c.TimeoutSecond)):
      cmd.Process.Kill()
      errMsg := "TIMEOUT: EXCEEDED -- Process killed"
		  err = errors.New(errMsg)
      outErr = []byte(errMsg)
	}

	return
}

func (c *CommandOutputWriter) ExecHeadOutput(command string) (outOk []byte, outErr []byte, err error) {
	return c.ExecFullOutput(fmt.Sprintf("%s | head", command))
}

func (c *CommandOutputWriter) ExecTailOutput(command string) (outOk []byte, outErr []byte, err error) {
	return c.ExecFullOutput(fmt.Sprintf("%s | tail", command))
}
