package banking

import (
	"errors"
	"fmt"
)

// Account comment, account 로 설정해두면 private 이므로 다른 코드에서 사용할 수 없음
type Account struct {
	owner   string
	balance int
}

var errorNoMoney = errors.New("can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Deposit method
//      receiver
// * 하지 않으면, 복사된 Account 를 사용하는 것이기 때문에 복사된 객체에 amount 를 더함
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

func (account Account) Balance() int {
	return account.balance
}

// Withdraw method
// go 에는 exception 이 없음, error 를 return 해줘야 한다.
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errorNoMoney
	}
	account.balance -= amount
	return nil
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

func (account Account) Owner() string {
	return account.owner
}

func (account Account) String() string {
	return fmt.Sprint(account.Owner(), "'s account has: ", account.balance)
}
