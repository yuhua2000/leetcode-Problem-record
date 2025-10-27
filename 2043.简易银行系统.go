package leetcode

type Bank struct {
	balance []int64
	n       int
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
		n:       len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	account1--
	account2--
	if account1 < 0 || account2 < 0 || account1 >= this.n || account2 >= this.n || this.balance[account1] < money {
		return false
	}

	this.balance[account1] -= money
	this.balance[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	account--
	if account < 0 || account >= this.n {
		return false
	}

	this.balance[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	account--
	if account < 0 || account >= this.n || this.balance[account] < money {
		return false
	}

	this.balance[account] -= money
	return true
}
