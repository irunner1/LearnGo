package arrays

// Transaction in Bank, passing sum from one accound to another
type Transaction struct {
	From string
	To   string
	Sum  float64
}

// NewTransaction creates transaction
func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

// Account in Bank collects name and balance
type Account struct {
	Name    string
	Balance float64
}

// NewBalanceFor creates transaction
func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

// applyTransaction applies transaction to both participants
func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}
