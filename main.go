package main

import (
	"fmt"


	"soor/soor"

)

func main() {

	go soor.LoginServer()
	go soor.Vlan102Redirector()
	go soor.Vlan103Redirector()
	go soor.Vlan104Redirector()

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

