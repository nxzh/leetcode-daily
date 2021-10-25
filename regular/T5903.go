package regular

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	bank := Bank{}
	bank.balance = balance
	return bank
}

func (this *Bank) toIndex(account int) int {
	return account - 1
}

func (this *Bank) isValidAccountId(account int) bool {
	return this.toIndex(account) >= 0 && this.toIndex(account) < len(this.balance)
}
func (this *Bank) hasSufficientAmountUnderAccount(account int, money int64) bool {
	if this.balance[this.toIndex(account)] >= money {
		return true
	}
	return false
}
func (this *Bank) minusAmount(account int, money int64) {
	this.balance[this.toIndex(account)] -= money
}
func (this *Bank) addAmount(account int, money int64) {
	this.balance[this.toIndex(account)] += money
}
func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.isValidAccountId(account1) || !this.isValidAccountId(account2) {
		return false
	}

	if !this.hasSufficientAmountUnderAccount(account1, money) {
		return false
	}
	this.minusAmount(account1, money)
	this.addAmount(account2, money)
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.isValidAccountId(account) {
		return false
	}
	this.addAmount(account, money)
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isValidAccountId(account) {
		return false
	}

	if !this.hasSufficientAmountUnderAccount(account, money) {
		return false
	}
	this.minusAmount(account, money)
	return true
}

func T5903() {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	bank.Withdraw(3, 10) // 返回 true ，账户 3 的余额是 $20 ，所以可以取款 $10 。
	// 账户 3 余额为 $20 - $10 = $10 。
	bank.Transfer(5, 1, 20) // 返回 true ，账户 5 的余额是 $30 ，所以可以转账 $20 。
	// 账户 5 的余额为 $30 - $20 = $10 ，账户 1 的余额为 $10 + $20 = $30 。
	bank.Deposit(5, 20) // 返回 true ，可以向账户 5 存款 $20 。
	// 账户 5 的余额为 $10 + $20 = $30 。
	bank.Transfer(3, 4, 15) // 返回 false ，账户 3 的当前余额是 $10 。
	// 所以无法转账 $15 。
	bank.Withdraw(10, 50) // 返回 false ，交易无效，因为账户 10 并不存在。
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
