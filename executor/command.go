package executor

import (
	"fmt"
	"os/exec"
)

type CommandOutputWriter struct{}

func (c *CommandOutputWriter) ExecFullOutput(command string) (outOk []byte, outErr []byte, err error) {

	cmd := exec.Command("bash", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		outErr = out
		return
	}
	outOk = out

	return
}

func (c *CommandOutputWriter) ExecHeadOutput(command string) (outOk []byte, outErr []byte, err error) {
	return c.ExecFullOutput(fmt.Sprintf("%s | head", command))
}

func (c *CommandOutputWriter) ExecTailOutput(command string) (outOk []byte, outErr []byte, err error) {
	return c.ExecFullOutput(fmt.Sprintf("%s | tail", command))
}
