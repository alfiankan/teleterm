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
	//result,er := exec.Command("cat","kkk.txt").Output()
	//fmt.Println(string(result))
	//fmt.Println(len(string(result)))
	//fmt.Println(er)
	//fmt.Println("pol")
	//d1 := result
	//err := ioutil.WriteFile("/tmp/log.txt", d1, 0644)
	//eerr(err)
}
