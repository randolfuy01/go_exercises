package bank

import (
	"fmt"
)

type address struct {
	house_number int16
	street       string
	city         string
	state        string
}

type account struct {
	account_name          string
	account_balance       float32
	account_interest_rate float32
	address               address
}

func create_account(account_name string, deposit_amount float32, account_address address) (result account) {
	result.account_balance = deposit_amount
	result.account_name = account_name
	result.account_interest_rate = 1.1
	result.address = account_address
	return
}

func generate_interest(user_account account, period int) account {
	for range period {
		user_account.account_balance = user_account.account_balance * (user_account.account_interest_rate)
	}
	return user_account
}

func withdraw(user_account account, withdraw_amount float32) (account, float32) {
	if withdraw_amount > user_account.account_balance {
		fmt.Print("Not enough money to withdraw")
		return user_account, 0
	}

	user_account.account_balance = user_account.account_balance - withdraw_amount
	return user_account, withdraw_amount
}
