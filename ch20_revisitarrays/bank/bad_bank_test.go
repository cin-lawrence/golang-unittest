package bank_test

import (
	"golang-unittest/ch20/assert"
	"golang-unittest/ch20/bank"
	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = bank.Account{Name: "Riya", Balance: 100}
		chris = bank.Account{Name: "Chris", Balance: 75}
		adil  = bank.Account{Name: "Adil", Balance: 200}

		transactions = []bank.Transaction{
			bank.NewTransaction(chris, riya, 100),
			bank.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account *bank.Account) float64 {
		return bank.NewBalanceFor(account, transactions).Balance
	}

	assert.AssertEqual(t, newBalanceFor(&riya), 200)
	assert.AssertEqual(t, newBalanceFor(&chris), 0)
	assert.AssertEqual(t, newBalanceFor(&adil), 175)
}
