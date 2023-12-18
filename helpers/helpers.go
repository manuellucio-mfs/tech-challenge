package helpers

func CountTransactionsByMount(mount int, contByMount map[string]int) map[string]int {
	switch mount {
	case 1:
		value := contByMount["enero"]
		contByMount["enero"] = value + 1
	case 2:
		value := contByMount["febrero"]
		contByMount["febrero"] = value + 1
	case 3:
		value := contByMount["marzo"]
		contByMount["marzo"] = value + 1
	case 4:
		value := contByMount["abril"]
		contByMount["abril"] = value + 1
	case 5:
		value := contByMount["mayo"]
		contByMount["mayo"] = value + 1
	case 6:
		value := contByMount["junio"]
		contByMount["junio"] = value + 1
	case 7:
		value := contByMount["julio"]
		contByMount["julio"] = value + 1
	case 8:
		value := contByMount["agosto"]
		contByMount["agosto"] = value + 1
	case 9:
		value := contByMount["septiembre"]
		contByMount["septiembre"] = value + 1
	case 10:
		value := contByMount["octubre"]
		contByMount["octubre"] = value + 1
	case 11:
		value := contByMount["noviembre"]
		contByMount["noviembre"] = value + 1
	case 12:
		value := contByMount["diciembre"]
		contByMount["diciembre"] = value + 1
	}
	return contByMount
}

func InicializarData() (map[string]int, map[string]SimpleTransaction) {
	contByMount := map[string]int{
		"enero":      0,
		"febrero":    0,
		"marzo":      0,
		"abril":      0,
		"mayo":       0,
		"junio":      0,
		"julio":      0,
		"agosto":     0,
		"septiembre": 0,
		"octubre":    0,
		"noviembre":  0,
		"diciembre":  0,
	}
	transactionByMount := map[string]SimpleTransaction{
		"enero": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"febrero": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"marzo": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"abril": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"mayo": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"junio": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"julio": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"agosto": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"septiembre": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"octubre": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"noviembre": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
		"diciembre": {
			SumaryCredit:   0,
			SumarDebit:     0,
			AverageCredit:  0,
			AverageDebit:   0,
			CantTranCredit: 0,
			CantTranDebit:  0,
		},
	}

	return contByMount, transactionByMount
}

func TransactionsByMount(mount int, transaccionValue float32, transaccionType bool, transaccionByMount map[string]SimpleTransaction) map[string]SimpleTransaction {
	var value SimpleTransaction
	switch mount {
	case 1:
		value = transaccionByMount["enero"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["enero"] = value
	case 2:
		value := transaccionByMount["febrero"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["febrero"] = value
	case 3:
		value := transaccionByMount["marzo"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["marzo"] = value
	case 4:
		value := transaccionByMount["abril"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["abril"] = value
	case 5:
		value := transaccionByMount["mayo"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["mayo"] = value
	case 6:
		value := transaccionByMount["junio"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["junio"] = value
	case 7:
		value := transaccionByMount["julio"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["julio"] = value
	case 8:
		value := transaccionByMount["agosto"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["agosto"] = value
	case 9:
		value := transaccionByMount["septiembre"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["septiembre"] = value
	case 10:
		value := transaccionByMount["octubre"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["octubre"] = value
	case 11:
		value := transaccionByMount["noviembre"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["noviembre"] = value
	case 12:
		value := transaccionByMount["diciembre"]
		if transaccionType {
			value.SumarDebit = value.SumarDebit + transaccionValue
			value.CantTranDebit = value.CantTranDebit + 1
		} else {
			value.SumaryCredit = value.SumaryCredit + transaccionValue
			value.CantTranCredit = value.CantTranCredit + 1
		}
		transaccionByMount["diciembre"] = value
	}
	return transaccionByMount
}
