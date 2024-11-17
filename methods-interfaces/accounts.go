package accounts

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(first string, last string) {
	a.FirstName = first
	a.LastName = last
}

type Employee struct {
	Account
	NoOfCredits float64
}

func (a Employee) CheckCredits() float64 {
	return a.NoOfCredits
}

func (a *Employee) AddCredits(credit float64) {
	a.NoOfCredits += credit
}

func (a *Employee) RemoveCredits(credit float64) {
	a.NoOfCredits -= credit
}

func (a *Employee) ChangeAccountName(first string, last string) {
	a.Account.FirstName = first
	a.Account.LastName = last
}

func (a Account) PrintName() string {
	return fmt.Sprintf("%v %v", a.FirstName, a.LastName)
}
