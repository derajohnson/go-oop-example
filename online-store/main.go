package main

import (
	"fmt"

	"github.com/derajohnson/accounts"
)

func main() {
	a := accounts.Account{"Vera", "Wang"}
	e := accounts.Employee{a, 54}
	fmt.Println(a.PrintName())
	e.AddCredits(23.4)
	e.RemoveCredits(2.5)
	getCredits := e.CheckCredits()
	e.ChangeAccountName("Kristina", "Burrows")
	printName := e.PrintName()

	fmt.Printf("%f \n %v \n", getCredits, printName)

}
