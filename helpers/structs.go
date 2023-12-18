package helpers

type Account struct {
	IdAccount          string        `json:"idAccount"`
	TotalBalance       float32       `json:"totalBalance"`
	AverageCredit      float32       `json:"averageCredit"`
	AverageDebit       float32       `json:"averageDebit"`
	CreditTransactions []Transaction `json:"creditTransactions"`
	DebitTransactions  []Transaction `json:"debitTransactions"`
}

type Transaction struct {
	IdTransaction string `json:"id"`
	Date          string `json:"date"`
	Value         string `json:"value"`
}

type TransationsMount struct {
	Mount map[string]SimpleTransaction
}

type SimpleTransaction struct {
	SumaryCredit   float32
	SumarDebit     float32
	AverageCredit  float32
	AverageDebit   float32
	CantTranCredit int
	CantTranDebit  int
}
