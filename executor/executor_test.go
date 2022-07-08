package executor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecStandartCommandFullOutput(t *testing.T) {
	cmd := new(CommandOutputWriter)
	outputOk, outputErr, err := cmd.ExecFullOutput("go")
	fmt.Println("Success :")
	fmt.Print(string(outputOk))
	fmt.Println("Error :")
	fmt.Print(string(outputErr))
	assert.Nil(t, err)
}

func TestExecStandartCommandTailOutput(t *testing.T) {
	cmd := new(CommandOutputWriter)
	outputOk, outputErr, err := cmd.ExecTailOutput("ls /bin")
	fmt.Println("Success :")
	fmt.Print(string(outputOk))
	fmt.Println("Error :")
	fmt.Print(string(outputErr))
	assert.Nil(t, err)
}

func TestExecStandartCommandHeadOutput(t *testing.T) {
	cmd := new(CommandOutputWriter)
	outputOk, outputErr, err := cmd.ExecHeadOutput("ls /bin")
	fmt.Println("Success :")
	fmt.Print(string(outputOk))
	fmt.Println("Error :")
	fmt.Print(string(outputErr))
	assert.Nil(t, err)
}
