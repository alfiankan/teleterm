package main

import (
	"fmt"
	"teleterm/handler"
)

func eerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {
	handler.Begin()
	//result,_ := exec.Command("git").Output()
	//fmt.Println(string(result))
	//d1 := result
	//err := ioutil.WriteFile("/tmp/log.txt", d1, 0644)
	//eerr(err)
}
