package controller

import (
	"io/ioutil"
	"os/exec"
)

func eerr(e error) {
	if e != nil {
		panic(e)
	}
}

func CmdExec(cmd string) {
	result, err := exec.Command(cmd).Output()
	eerr(err)
	//fmt.Println(string(result))
	d1 := result
	err = ioutil.WriteFile("/tmp/log.txt", d1, 0644)
	eerr(err)
}
